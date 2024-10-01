package database

import (
	"gorm.io/gorm"
	"sing-song/internal/models"
	"sing-song/pkg/consts"
)

type DataPostgres struct {
	conn *gorm.DB
}

func NewDataPostgres(conn *gorm.DB) *DataPostgres {
	return &DataPostgres{conn: conn}
}

func (d *DataPostgres) AddNewSong(newSong *models.Songs) error {
	db := d.conn.Create(&newSong)
	return db.Error
}
func (d *DataPostgres) EditInfo(song *models.Songs) error {
	db := d.conn.Updates(&song)
	return db.Error
}

func (d *DataPostgres) RemoveSong(id int) error {
	db := d.conn.Table("songs").Where("id", id).Update("active", false)
	return db.Error
}

func (d *DataPostgres) GetSong(id int) (song *models.Songs, err error) {
	db := d.conn.Where("id", id).First(&song)
	return song, db.Error
}

func (d *DataPostgres) GetSongs(limit, offset int, group, dateFrom, dateTo string) (songs []*models.Songs, err error) {
	if group != "" {
		err = d.conn.Raw(consts.GetSongsByGroupSQL, dateFrom, dateTo, group, limit, offset).Scan(&songs).Error
	} else {
		err = d.conn.Raw(consts.GetSongsSQL, dateFrom, dateTo, limit, offset).Scan(&songs).Error
	}
	return songs, err
}
