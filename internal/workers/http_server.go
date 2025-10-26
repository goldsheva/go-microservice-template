package workers

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/goldsheva/go-microservice-template/internal/configs"
	"github.com/goldsheva/go-microservice-template/internal/http_handlers"
	"github.com/goldsheva/go-microservice-template/internal/middlewares"
	"github.com/sirupsen/logrus"
)

func GoHTTPServer(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	config := configs.GetEnvConfig()

	switch config.LogLevel {
	case logrus.DebugLevel:
		gin.SetMode(gin.DebugMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(
		gin.Logger(),
		gin.Recovery(),
		middlewares.CorsMiddleware(),
	)

	router.GET("/api/healthcheck", http_handlers.HealthCheckHttpHandler)

	httpServer := &http.Server{
		Addr:              fmt.Sprintf(":%d", config.HTTP_PORT),
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
	}

	go func() {
		<-ctx.Done()

		logrus.WithFields(logrus.Fields{"gopher": "http_server"}).Warn("HTTP server stopped ...")

		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			logrus.WithFields(logrus.Fields{"gopher": "http_server"}).Errorf("Shutdown error: %v", err)
		}
	}()

	logrus.WithFields(logrus.Fields{"gopher": "http_server"}).Infof("Starting HTTP server on :%d", config.HTTP_PORT)

	if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logrus.WithFields(logrus.Fields{"gopher": "http_server"}).Errorf("HTTP server error: %v", err)
	}
}
