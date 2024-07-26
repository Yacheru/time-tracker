package server

import (
	"EffectiveMobile/internal/server/http/router"
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"EffectiveMobile/init/config"
	"EffectiveMobile/init/logger"
	"EffectiveMobile/internal/repository/postgres"
	"EffectiveMobile/pkg/constants"
)

type HTTPServer struct {
	server *http.Server
}

func NewServer() (*HTTPServer, error) {
	ctx := context.Background()

	db, err := postgres.NewPostgresConnection(ctx, config.ServerConfig.POSTGRESDsn)
	if err != nil {
		return nil, err
	}

	routes := setupRouter()
	api := routes.Group("/api")

	router.NewRoutes(api, db).Routers()

	server := &http.Server{
		Addr:         ":" + config.ServerConfig.APIPort,
		Handler:      routes,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	return &HTTPServer{server: server}, nil
}

func (s *HTTPServer) Run() error {
	go func() {
		if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Fatal("", logrus.Fields{constants.LoggerCategory: constants.Server})
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	logger.Info("Shutdown server...", logrus.Fields{constants.LoggerCategory: constants.Server})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		logger.Fatal("Error shuts down http-server", logrus.Fields{constants.LoggerCategory: constants.Server})

		return err
	}

	<-ctx.Done()

	return nil
}

func setupRouter() *gin.Engine {
	var mode = gin.ReleaseMode
	if config.ServerConfig.APIDebug {
		mode = gin.DebugMode
	}
	gin.SetMode(mode)

	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(gin.LoggerWithFormatter(logger.HTTPLogger))

	return router
}
