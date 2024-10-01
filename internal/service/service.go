package service

import (
	"sing-song/internal/database"
	"sing-song/internal/models"
)

type Data interface {
	AddNewSong(newSong *models.Songs) error
	EditInfo(song *models.Songs) error
	RemoveSong(id int) error
	GetSong(id, page, couplet int) (*models.Songs, error)
	GetSongs(page, count int, group, dateFrom, dateTo string) (songs []*models.Songs, err error)
}
type Service struct {
	Data
}

func NewService(db *database.Database) *Service {
	return &Service{
		Data: NewDataService(db),
	}
}
