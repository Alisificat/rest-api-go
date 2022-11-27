package main

import (
	"context"
	"github.com/Serminaz/GoRun/todo2"
	"github.com/Serminaz/GoRun/todo2/pkg/handler"
	"github.com/Serminaz/GoRun/todo2/pkg/repository"
	"github.com/Serminaz/GoRun/todo2/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		//если функция возвр ошибку, то прерываем приложение
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}
	// иницилизация базы
	if err := godotenv.Load(); err != nil {
		logrus.Fatal("error loading env variables: %s", err.Error())
	}
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatal("failed to initialize1 db %s: ", err.Error())
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)    //сервис зависит от репозитория
	handlers := handler.NewHandler(services) // хандлер от сервиса

	srv := new(todo.Server) // экземпляр сервера
	//в качестве значение порта можно передать значение из вайпера по ключу
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatal("error occured while running http server: %s", err.Error())
		}
	}()
	logrus.Println("TodoApp Started")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	logrus.Println("TodoApp Shutting Down")
	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Error("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Error("error occured on db  connection close: %s", err.Error())
	}
}

func initConfig() error { //для иницил  конфиг файлов
	viper.AddConfigPath("configs") // имч директории
	viper.SetConfigName("config")  // имя файла
	return viper.ReadInConfig()
	//считывает значение конфигов и записывает их во внутр объект вайпера
	// возвр ошибку
}
