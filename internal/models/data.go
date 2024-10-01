package models

import "time"

type Songs struct {
	Id          int    `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	SongGroup   string `json:"song_group"`
	Lyric       string
	ReleaseDate time.Time  `json:"release_date"`
	Link        string     `json:"link"`
	Active      bool       `json:"active"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}
