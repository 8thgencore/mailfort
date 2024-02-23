package app

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/8thgencore/mailfort/docs"
	"github.com/8thgencore/mailfort/internal/app/grpc"
	"github.com/8thgencore/mailfort/internal/config"
	"github.com/8thgencore/mailfort/internal/delivery/http"
	mailService "github.com/8thgencore/mailfort/internal/service/mail"
	"github.com/8thgencore/mailfort/pkg/logger/slogpretty"
)

// @title			MailFort API
// @version		1.0
// @description	MailFort API is a service for handling email-related operations.
//
// @contact.name	Tom Jerry
// @contact.url	https://github.com/8thgencore/mailfort
// @contact.email	test@gmail.com
//
// @license.name	MIT
// @license.url	https://opensource.org/licenses/MIT
//
// @host			api.example.com
// @BasePath		/v1
// @schemes		http https
func Run(configPath string) {
	// Load configuration
	cfg, err := config.NewConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}

	// Set up logger based on configuration
	log := newSlogLogger(cfg.Log.Slog)

	// Log information about the start of the application
	log.Info("starting passfort", slog.String("env", string(cfg.Env)))
	log.Debug("debug messages are enabled")

	// Dependency injection
	mailService := mailService.NewMailService(log)
	mailHandler := http.NewMailHandler(mailService)

	// Init router
	router, err := http.NewRouter(log, cfg, *mailHandler)
	if err != nil {
		log.Error("Error initializing router", "error", err.Error())
		os.Exit(1)
	}

	// Start REST API server
	listenAddr := fmt.Sprintf("%s:%s", cfg.HTTP.Host, cfg.HTTP.Port)

	log.Info("Starting the HTTP server", "listen_address", listenAddr)

	go func() {
		err = router.Serve(listenAddr)
		if err != nil {
			log.Error("Error starting the HTTP server", "error", err.Error())
			os.Exit(1)
		}
	}()

	// Start gRPC API Server
	grpcApp := grpc.New(log, mailService, cfg.GRPC.Port)

	go grpcApp.MustRun()

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	sign := <-stop

	log.Info("stopping application", slog.String("signal", sign.String()))

	grpcApp.Stop()

	log.Info("application stopped")
}

func newSlogLogger(c config.Slog) *slog.Logger {
	o := &slog.HandlerOptions{Level: c.Level, AddSource: c.AddSource}
	w := os.Stdout
	var h slog.Handler

	switch c.Format {
	case "pretty":
		h = slogpretty.NewHandler().
			WithAddSource(c.AddSource).
			WithLevel(c.Level).
			WithLevelEmoji(c.Pretty.Emoji).
			WithTimeLayout(c.Pretty.TimeLayout).
			WithFieldsFormat(c.Pretty.FieldsFormat)
	case "json":
		h = slog.NewJSONHandler(w, o)
	case "text":
		h = slog.NewTextHandler(w, o)
	}
	return slog.New(h)
}
