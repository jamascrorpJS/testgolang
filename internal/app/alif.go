package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jamascrorpJS/eBank/internal/delivery"
	"github.com/jamascrorpJS/eBank/internal/repository"
	"github.com/jamascrorpJS/eBank/internal/server"
	"github.com/jamascrorpJS/eBank/internal/service"
	"github.com/jamascrorpJS/eBank/pkg/database"
	"github.com/joho/godotenv"
)

func StartServer() {
	config := setEnv()
	db, err := database.NewClientDB(config)
	if err != nil {
		log.Default().Println(err)
	}

	repository := repository.NewRepository(db)

	service := service.NewService(repository)

	handler := delivery.NewHandler(service)

	engine := handler.NewEngine()

	server := server.NewServer(":8080", engine)

	go func() {
		if err := server.Start(); err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)

	defer shutdown()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}

func setEnv() database.ConfigDataBase {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("env file is not exist")
	}
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	dbname := os.Getenv("POSTGRES_DB")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	return database.ConfigDataBase{
		Host:     host,
		Port:     port,
		DBName:   dbname,
		User:     user,
		Password: password,
	}
}
