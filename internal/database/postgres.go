package database

import (
	"fmt"
	"github.com/spf13/viper"
	postgresDriver "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"sing-song/internal/models"
)

func NewPostgresGorm() (*gorm.DB, error) {
	Host := viper.GetString("db.host")
	Port := viper.GetUint16("db.port")
	Username := viper.GetString("db.username")
	Password := os.Getenv("DB_PASSWORD")
	DBName := viper.GetString("db.dbname")
	connString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Dushanbe",
		Host, Username, Password, DBName, Port)
	conn, err := gorm.Open(postgresDriver.Open(connString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		log.Printf("%s GetPostgresConnection -> Open error: ", err.Error())
		return nil, err
	}
	err = conn.AutoMigrate(&models.Songs{})
	if err != nil {
		log.Println(err)
	}
	log.Println("Postgres Connection success: ", Host)
	return conn, nil
}
