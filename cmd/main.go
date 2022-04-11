package main

import (
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/yachnytskyi/pizza-world/app/handler"
	"github.com/yachnytskyi/pizza-world/app/repository"
	"github.com/yachnytskyi/pizza-world/app/service"
	"github.com/yachnytskyi/pizza-world/cmd/server"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error initializing: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		DB_HOST:  viper.GetString("db.db_host"),
		DB_PORT:  viper.GetString("db.db_port"),
		DB_NAME:  viper.GetString("db.db_name"),
		DB_USER:  viper.GetString("db.db_user"),
		DB_PASS:  os.Getenv("db_pass"),
		SSL_MODE: viper.GetString("db.ssl_mode"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("Error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
