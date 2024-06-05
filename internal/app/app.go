package app

import (
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/8thgencore/mailfort/internal/config"
	"github.com/8thgencore/mailfort/internal/delivery/grpc"
	"github.com/8thgencore/mailfort/internal/delivery/http"
	mailService "github.com/8thgencore/mailfort/internal/service/mail"
	"github.com/8thgencore/mailfort/pkg/logger/slogpretty"
)

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
	mailService := mailService.NewMailService(log, &cfg.Mail)

	// // Start gRPC API Server
	grpcApp := grpc.New(log, mailService, cfg.GRPC.Port)
	go grpcApp.MustRun()

	// Start the HTTP server with gRPC-Gateway
	httpApp := http.New(log, mailService, cfg.HTTP.Port, cfg.GRPC.Port)
	go httpApp.MustRun()

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
