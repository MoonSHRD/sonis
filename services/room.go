package services

import (
	"github.com/MoonSHRD/sonis/app"
	"github.com/MoonSHRD/sonis/models"
	"github.com/MoonSHRD/sonis/repositories"
)

type RoomService struct {
	app                    *app.App
	roomRepository         *repositories.RoomRepository
	chatCategoryRepository *repositories.ChatCategoryRepository
}

func NewRoomService(a *app.App, rr *repositories.RoomRepository, ccr *repositories.ChatCategoryRepository) *RoomService {
	return &RoomService{
		app:                    a,
		roomRepository:         rr,
		chatCategoryRepository: ccr,
	}
}

func (rs *RoomService) PutRoom(room *models.Room) (*models.Room, error) {
	return rs.roomRepository.PutRoom(room)
}

func (rs *RoomService) GetRoomsByCoords(lat float64, lon float64, radius int) (*[]models.Room, error) {
	return rs.roomRepository.GetRoomsByCoords(lat, lon, radius)
}

func (rs *RoomService) GetRoomByID(id int) (*models.Room, error) {
	return rs.roomRepository.GetRoomByID(id)
}

func (rs *RoomService) GetAllRooms() ([]models.Room, error) {
	return rs.roomRepository.GetAllRooms()
}

func (rs *RoomService) GetRoomsByCategoryID(categoryID int) ([]models.Room, error) {
	return rs.roomRepository.GetRoomsByCategoryID(categoryID)
}

func (rs *RoomService) GetRoomsByParentGroupID(parentGroupID string) ([]models.Room, error) {
	return rs.roomRepository.GetRoomsByParentGroupID(parentGroupID)
}

func (rs *RoomService) GetAllCategories() ([]models.ChatCategory, error) {
	return rs.chatCategoryRepository.GetAllCategories()
}

func (rs *RoomService) UpdateRoom(room *models.Room) (*models.Room, error) {
	return rs.roomRepository.UpdateRoom(room)
}

func (rs *RoomService) DeleteRoom(id int) error {
	return rs.roomRepository.DeleteRoom(id)
}
