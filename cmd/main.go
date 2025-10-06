package main

import (
	"fmt"

	"github.com/Hot-One/kizen-go-service/api"
	"github.com/Hot-One/kizen-go-service/config"
	"github.com/Hot-One/kizen-go-service/models"
	"github.com/Hot-One/kizen-go-service/pkg/logger"
	postgresConn "github.com/Hot-One/kizen-go-service/pkg/postgres"
	"github.com/Hot-One/kizen-go-service/storage"
	"github.com/gin-gonic/gin"
	gormLog "gorm.io/gorm/logger"
)

// @title Kizen API
// @version 1.0
// @description API for Kizen application
// @BasePath /v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	var (
		cfg         = config.Load()
		loggerLevel = new(string)

		gormConfig = &postgresConn.GormConfig{
			SkipDefaultTransaction: true,
			Logger:                 gormLog.Default.LogMode(gormLog.Info),
		}
	)

	*loggerLevel = logger.LevelDebug

	switch cfg.Environment {
	case config.DebugMode:
		*loggerLevel = logger.LevelDebug
		gin.SetMode(gin.DebugMode)
	case config.TestMode:
		*loggerLevel = logger.LevelDebug
		gin.SetMode(gin.TestMode)
	default:
		*loggerLevel = logger.LevelInfo
		gin.SetMode(gin.ReleaseMode)
	}

	log := logger.New(*loggerLevel, cfg.ServiceName)
	defer func() {
		err := logger.Cleanup(log)
		if err != nil {
			log.Error("Failed to cleanup logger", logger.Error(err))
			return
		}
	}()

	postgres, err := postgresConn.ConnectPostgres(gormConfig, &cfg)
	if err != nil {
		log.Error("Failed to connect to PostgreSQL", logger.Error(err))
		return
	}

	if err := models.Migrate(postgres); err != nil {
		log.Error("Failed to migrate PostgreSQL database", logger.Error(err))
		return
	}

	var storage = storage.NewStorage(postgres, log)
	var server = api.SetUpRouter(cfg, log, storage)

	if err := server.Run(fmt.Sprintf(":%d", cfg.Port)); err != nil {
		log.Error("Failed to run server", logger.Error(err))
		return
	}
}
