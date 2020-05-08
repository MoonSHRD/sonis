package migrations

import (
	migrate "github.com/rubenv/sql-migrate"
)

var (
	MigrationsList = &migrate.MemoryMigrationSource{
		Migrations: []*migrate.Migration{
			&migrate.Migration{
				Id:   "1",
				Up:   []string{_1Up},
				Down: []string{_1Down},
			},
			&migrate.Migration{
				Id:   "2",
				Up:   []string{_2Up},
				Down: []string{_2Down},
			},
			&migrate.Migration{
				Id:   "3",
				Up:   []string{_3Up},
				Down: []string{_3Down},
			},
			&migrate.Migration{
				Id:   "4",
				Up:   []string{_4Up},
				Down: []string{_4Down},
			},
			&migrate.Migration{
				Id:   "5",
				Up:   []string{_5Up},
				Down: []string{_5Down},
			},
			&migrate.Migration{
				Id:   "6",
				Up:   []string{_6Up},
				Down: []string{_6Down},
			},
			&migrate.Migration{
				Id:   "7",
				Up:   []string{_7Up},
				Down: []string{_7Down},
			},
			&migrate.Migration{
				Id:   "8",
				Up:   []string{_8Up},
				Down: []string{_8Down},
			},
			&migrate.Migration{
				Id:   "9",
				Up:   []string{_9Up},
				Down: []string{_9Down},
			},
		},
	}
)