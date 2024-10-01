package service

import (
	"sing-song/internal/database"
	"sing-song/internal/models"
	"sing-song/pkg/utils"
)

type DataService struct {
	db database.Data
}

func NewDataService(db database.Data) *DataService {
	return &DataService{db: db}
}

func (d *DataService) AddNewSong(newSong *models.Songs) error {
	return d.db.AddNewSong(newSong)
}
func (d *DataService) EditInfo(song *models.Songs) error {
	return d.db.EditInfo(song)
}

func (d *DataService) RemoveSong(id int) error {
	return d.db.RemoveSong(id)
}

func (d *DataService) GetSong(id, page, couplet int) (*models.Songs, error) {
	song, err := d.db.GetSong(id)
	if err != nil {
		return nil, err
	}
	song.Lyric = utils.LyricHelper(song.Lyric, page, couplet)
	return song, nil
}

func (d *DataService) GetSongs(page, count int, group, dateFrom, dateTo string) (songs []*models.Songs, err error) {
	offset := (page * count) - count
	return d.db.GetSongs(count, offset, group, dateFrom, dateTo)
}
