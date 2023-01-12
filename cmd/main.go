package main

import (
	"context"
	"github.com/AbdullohAbdullayev/todo-app-L.git"
	"github.com/AbdullohAbdullayev/todo-app-L.git/docs"
	"github.com/AbdullohAbdullayev/todo-app-L.git/pkg/handler"
	"github.com/AbdullohAbdullayev/todo-app-L.git/pkg/repository"
	"github.com/AbdullohAbdullayev/todo-app-L.git/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"os"
	"os/signal"
	"syscall"
)

//	@title			Todo App API
//	@version		1.0
//	@description	API Server for TodoList Application

//	@host		localhost:81
//	@BasePath	/

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization

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
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRouters()); err != nil {
			logrus.Fatalf("error occur while running http sever %s", err)
		}
	}()
	logrus.Println("Todo app started")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	logrus.Println("Todo app Shutting down")
	err = srv.Shutdown(context.Background())
	if err != nil {
		logrus.Printf("Error shutdowning server : %s", err.Error())
	}
	err = db.Close()
	if err != nil {
		logrus.Printf("Error closing db : %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func swagConfig() {
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "petstore.swagger.io"
	docs.SwaggerInfo.BasePath = "/v2"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}
