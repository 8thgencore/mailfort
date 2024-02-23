package http

import (
	"log/slog"
	"strings"

	"github.com/8thgencore/mailfort/internal/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Router is a wrapper for HTTP router
type Router struct {
	*gin.Engine
}

func NewRouter(
	log *slog.Logger,
	cfg *config.Config,
	mailHandler MailHandler,
) (*Router, error) {
	// Disable debug mode in production
	if cfg.Env == config.Prod {
		gin.SetMode(gin.ReleaseMode)
	}

	// CORS
	ginConfig := cors.DefaultConfig()
	allowOrigins := cfg.HTTP.AllowOrigins
	originsList := strings.Split(allowOrigins, ",")
	ginConfig.AllowOrigins = originsList

	// Init router
	router := gin.New()
	router.Use(sloggin.New(log), gin.Recovery(), cors.New(ginConfig))

	// Swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Endpoints
	v1 := router.Group("/v1")
	{
		v1.POST("/email-confirmation", mailHandler.SendConfirmationEmail)
		v1.POST("/password-reset", mailHandler.SendPasswordResetEmail)
	}

	return &Router{
		router,
	}, nil
}

// Serve starts the HTTP server
func (r *Router) Serve(listenAddr string) error {
	return r.Run(listenAddr)
}
