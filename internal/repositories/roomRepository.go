package repositories

import (
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/MoonSHRD/sonis/internal/database"
	"github.com/MoonSHRD/sonis/internal/models"
	"github.com/prprprus/scheduler"
)

type RoomRepository struct {
	db                    *database.Database
	deletingRoomScheduler *scheduler.Scheduler
	logger                *logrus.Logger
}

func NewRoomRepository(db *database.Database) (*RoomRepository, error) {
	if db != nil {
		deletingRoomScheduler, err := scheduler.NewScheduler(10000)
		if err != nil {
			return nil, err
		}
		return &RoomRepository{
			db:                    db,
			deletingRoomScheduler: deletingRoomScheduler,
			logger:                logrus.New(),
		}, nil
	}
	return nil, fmt.Errorf("database connection is null")
}

func (rr *RoomRepository) PutRoom(room *models.Room) error {
	stmt, err := rr.db.GetDatabaseConnection().Preparex("INSERT INTO rooms (latitude, longitude, ttl, room_id, category) VALUES (?, ?, ?, ?, ?);")
	if err != nil {
		return err
	}
	res, err := stmt.Exec(room.Latitude, room.Longitude, room.TTL, room.RoomID, room.Category)
	if err != nil {
		return err
	}
	roomID, err := res.LastInsertId()
	if err != nil {
		return err
	}
	rr.deletingRoomScheduler.Delay().Second(room.TTL).Do(func() {
		stmt, err := rr.db.GetDatabaseConnection().Preparex("DELETE FROM rooms WHERE id = ?")
		_, err = stmt.Exec(roomID)
		if err != nil {
			rr.logger.Error("Cannot delete room " + room.RoomID + ". Reason: " + err.Error())
		}
	})
	return nil
}

func (rr *RoomRepository) GetRoomsByCoords(userLat float64, userLon float64, radius int) (*[]models.Room, error) {
	var rooms []models.Room
	stmt, err := rr.db.GetDatabaseConnection().Preparex("SELECT * FROM rooms WHERE SQRT(POW(latitude-?, 2) + POW(longitude-?, 2)) < ?")
	if err != nil {
		return nil, err
	}
	err = stmt.Select(&rooms, userLat, userLon, radius)
	if err != nil {
		return nil, err
	}
	return &rooms, nil
}

func (rr *RoomRepository) GetRoomByRoomID(roomID string) (*models.Room, error) {
	stmt, err := rr.db.GetDatabaseConnection().Preparex("SELECT * FROM rooms WHERE room_id = ?")
	if err != nil {
		return nil, err
	}
	var room models.Room
	stmt.Get(&room, roomID)
	return &room, nil
}
