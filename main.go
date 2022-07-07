package main

import (
	"context"
	"net/http"
	"os"

	"github.com/DmitriyZhevnov/rest-api/internal/config"
	"github.com/DmitriyZhevnov/rest-api/internal/handler"
	"github.com/DmitriyZhevnov/rest-api/internal/repository"
	"github.com/DmitriyZhevnov/rest-api/internal/service"

	fs "github.com/DmitriyZhevnov/rest-api/pkg/client/firestore"
	"github.com/DmitriyZhevnov/rest-api/pkg/client/postgresql"
	"github.com/DmitriyZhevnov/rest-api/pkg/hash"
	"github.com/DmitriyZhevnov/rest-api/pkg/logging"
	"github.com/julienschmidt/httprouter"
)

func main() {
	logger := logging.GetLogger()

	logger.Info("create router")
	router := httprouter.New()

	cfg := config.GetConfig()

	hasher := hash.NewSHA1Hasher(cfg.Auth.PasswordSalt)

	logger.Info("register user handler")

	firestoreClient, err := fs.NewClient(context.Background(), cfg.Storage.Firestore.ProjectID)
	if err != nil {
		panic(err)
	}

	postgresClient, err := postgresql.NewClient(cfg.Storage.Postgresql)
	if err != nil {
		panic(err)
	}

	storage := repository.NewRepository(postgresClient, *firestoreClient, logger)

	service := service.NewService(hasher, storage, logger)

	handler := handler.NewHandler(service, logger)
	handler.Register(router)

	startServer(router, cfg)
}

func startServer(router *httprouter.Router, cfg *config.Config) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	logger := logging.GetLogger()
	logger.Info("start application")

	if err := http.ListenAndServe(":"+port, router); err != nil {
		logger.Fatalf("Error launching REST API server: %v", err)
	}
}
