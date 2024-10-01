package main

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"sing-song/internal/database"
	api "sing-song/internal/handlers"
	"sing-song/internal/server"
	"sing-song/internal/service"
	"sing-song/pkg/logging"
	"sing-song/pkg/utils"
)

// @title My track library
// @version 0.0.1
// @description Api for playlist project

// @host localhost:5005
// @BasePath /api/v1

func main() {
	log := logging.GetLogger()
	if err := utils.InitConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error initializing env value: %s", err.Error())
	}
	conn, err := database.NewPostgresGorm()
	if err != nil {
		log.Fatalf("failed to initializing db: %s", err.Error())
	}
	repository := database.NewDatabase(conn)
	services := service.NewService(repository)
	handlers := api.NewHandler(services, log)
	srv := new(server.Server)
	if err = srv.Run(viper.GetString("server.port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
