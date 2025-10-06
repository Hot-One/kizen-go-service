package api

import (
	"github.com/Hot-One/kizen-go-service/api/docs"
	"github.com/Hot-One/kizen-go-service/api/handler/sms"
	"github.com/Hot-One/kizen-go-service/config"
	"github.com/Hot-One/kizen-go-service/pkg/logger"
	"github.com/Hot-One/kizen-go-service/storage"
	"github.com/gin-gonic/gin"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

type router struct {
	cfg  *config.Config
	log  logger.Logger
	strg storage.StorageInterface
}

func SetUpRouter(cfg config.Config, log logger.Logger, strg storage.StorageInterface) *gin.Engine {
	var (
		r      = gin.Default()
		option = router{
			cfg:  &cfg,
			log:  log,
			strg: strg,
		}
	)

	docs.SwaggerInfo.Title = cfg.ServiceName
	docs.SwaggerInfo.Schemes = []string{"http"}

	r.Use(gin.Recovery(), gin.Logger(), customCORSMiddleware())

	v1 := r.Group("/v1")
	{
		sms.NewHandler(v1, option.cfg, option.log, option.strg.Sms())
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}

func customCORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Max-Age", "3600")
		c.Header("Access-Control-Allow-Methods", "*")
		c.Header("Access-Control-Allow-Headers", "*")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
