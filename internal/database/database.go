package database

import (
	"gorm.io/gorm"
	"sing-song/internal/models"
)

type Data interface {
	AddNewSong(newSong *models.Songs) error
	EditInfo(song *models.Songs) error
	RemoveSong(id int) error
	GetSong(id int) (*models.Songs, error)
	GetSongs(limit, offset int, group, dateFrom, dateTo string) (songs []*models.Songs, err error)
}

type Database struct {
	Data
}

func NewDatabase(conn *gorm.DB) *Database {
	return &Database{
		Data: NewDataPostgres(conn),
	}
}
