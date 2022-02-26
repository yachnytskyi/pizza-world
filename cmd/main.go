package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"github.com/yachnytskyi/pizza-world/app/handler"
	"github.com/yachnytskyi/pizza-world/app/repository"
	"github.com/yachnytskyi/pizza-world/app/service"
	"github.com/yachnytskyi/pizza-world/cmd/server"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error initializing: %s", err.Error())
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
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	if err := srv.Run(viper.GetString("8000"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

// db, err := sqlx.Open("postgres", fmt.Sprintf("db_host=%s db_port=%s db_name=%s db_user=%s db_password=%s sslmode=%s",
// cfg.DB_HOST, cfg.DB_PORT, cfg.DB_NAME, cfg.DB_USER, cfg.DB_PASS, cfg.SSL_MODE))
