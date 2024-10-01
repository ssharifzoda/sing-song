package consts

const (
	GetSongsByGroupSQL = "select * from songs where created_at between ? and ? group by ? limit ? offset ?"
	GetSongsSQL        = "select * from songs where created_at between ? and ? limit ? offset ?"
)
