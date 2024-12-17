package main

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	handler "jwt/internal/contoller"
	"jwt/internal/repository"
	"jwt/internal/usecase"
	"jwt/pkg/database"
	"jwt/pkg/logger"
	"net/http"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("error loading .env file")
	}

	port := os.Getenv("HTTP_PORT")
	db := database.InitializeDBPostgres(3, 10)
	logger.InitLogger()

	repository := repository.NewRepository(db.GetDB())
	usecase := usecase.NewUseCase(&repository)
	handlers := handler.NewHandler(&usecase)
	router := handlers.Handle()

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatalf("connection failed: %s\n", err.Error())
	}

	log.Infof("server is running on port %s\n", port)
}
