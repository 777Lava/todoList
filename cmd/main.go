package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/777Lava/todo-app"
	"github.com/777Lava/todo-app/pkg/handler"
	"github.com/777Lava/todo-app/pkg/repository"
	"github.com/777Lava/todo-app/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error init configs %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading.env file %s", err.Error())

	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("failed to init database %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	go func(){
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("TodoApp Started")
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGTERM,syscall.SIGINT)
	<- quit	

	logrus.Print("TodoApp Shutdown")

	if err := srv.Shutdown(context.Background()); err!= nil {
		logrus.Errorf("error occured while shutting down http server %s", err.Error())
	}
	if err := db.Close(); err!= nil {
		logrus.Errorf("error occured while closing database %s", err.Error())
	
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
