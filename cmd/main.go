package main

import (
	"context"
	_ "embed"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"syscall"

	"github.com/goldsheva/go-microservice-template/internal/configs"
	"github.com/goldsheva/go-microservice-template/internal/database"
	"github.com/goldsheva/go-microservice-template/internal/models"
	"github.com/goldsheva/go-microservice-template/internal/workers"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}

	// Load environment variables from config.env file
	cwd, err := os.Getwd()
	if err != nil {
		logrus.Fatal("Cannot determine working directory: ", err)
	}
	envPath := filepath.Join(cwd, "config.env")

	if err := godotenv.Overload(envPath); err != nil {
		logrus.Fatal("Failed to load config.env: ", err)
	}

	config := configs.GetEnvConfig()

	logrus.SetLevel(config.LogLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	// Init DB Connections
	db := database.InitDB()
	if err := db.AutoMigrate(&models.ExampleModel{}); err != nil {
		logrus.Fatal("Failed to automigrate: ", err)
	}

	wg.Add(1)
	go workers.GoHTTPServer(ctx, wg)

	// Handle sigterm and await termChan signal
	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)

	<-termChan
	logrus.WithFields(logrus.Fields{"gopher": "main"}).Warn("Initiating shutdown...")
	cancelFunc()

	wg.Wait()
	logrus.WithFields(logrus.Fields{"gopher": "main"}).Warn("Shutdown complete. All processes stopped!")
}
