package main

import (
	atest "github.com/Elhemist/avito-test"
	"github.com/Elhemist/avito-test/internal/handler"
	repository "github.com/Elhemist/avito-test/internal/repositiry"
	"github.com/Elhemist/avito-test/internal/service"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := InitConfig(); err != nil {
		logrus.Fatalf("Config init error: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Usename:  viper.GetString("db.usename"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLmode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		logrus.Fatalf("DB init fail: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handlers := handler.NewHandler(service)

	srv := new(atest.Server)
	if err := srv.Start(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("Running error: %s", err.Error())
	}
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
