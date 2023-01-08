package main

import (
	"github.com/AbdullohAbdullayev/todo-app-L.git"
	"github.com/AbdullohAbdullayev/todo-app-L.git/pkg/handler"
	"github.com/AbdullohAbdullayev/todo-app-L.git/pkg/repository"
	"github.com/AbdullohAbdullayev/todo-app-L.git/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func main() {
	if err := initConfig(); err != nil {
		logrus.Fatalf("error intialion conifs : %s", err)
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env veribles %s", err)
	}

	db, err := repository.NewDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("error initialating db %s", err)
	}

	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	handlers := handler.NewHandler(service)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRouters()); err != nil {
		logrus.Fatalf("error occur while running http sever %s", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
