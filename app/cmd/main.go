package main

import (
	_ "github.com/lib/pq"
	"github.com/linqcod/course-work-5/app/internal/handler"
	"github.com/linqcod/course-work-5/app/internal/repository"
	"github.com/linqcod/course-work-5/app/internal/server"
	"github.com/linqcod/course-work-5/app/internal/service"
	"github.com/linqcod/course-work-5/app/pkg/postgresql"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig("./app/.env"); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	db, err := postgresql.NewPostgresDB(postgresql.Config{
		Host:     viper.GetString("POSTGRES_HOST"),
		Port:     viper.GetString("POSTGRES_PORT"),
		Username: viper.GetString("POSTGRES_USER"),
		Password: viper.GetString("POSTGRES_PASSWORD"),
		DBName:   viper.GetString("POSTGRES_DB"),
		SSLMode:  "disable",
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db :%s", err.Error())
	}

	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	if err := srv.Run(viper.GetString("SERVER_PORT"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig(file string) error {
	viper.SetConfigFile(file)
	viper.AutomaticEnv()

	return viper.ReadInConfig()
}
