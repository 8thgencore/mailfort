package http

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"log/slog"

	"github.com/8thgencore/mailfort/internal/service"
	mailpb "github.com/8thgencore/mailfort/protos/gen/go/mail/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type App struct {
	log      *slog.Logger
	port     int
	grpcPort int
	svc      service.MailService
}

// New creates a new HTTP server app.
func New(
	log *slog.Logger,
	svc service.MailService,
	port int,
	grpcPort int,
) *App {
	return &App{
		log:      log,
		svc:      svc,
		grpcPort: grpcPort,
		port:     port,
	}
}

// MustRun runs the HTTP server and panics if any error occurs.
func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	const op = "delivery.http.app.Run"

	log := a.log.With(
		slog.String("op", op),
		slog.Int("port", a.port),
	)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	// opts := []grpc.DialOption{}
	grpcEndpoint := fmt.Sprintf("localhost:%d", a.grpcPort)
	err := mailpb.RegisterMailServiceHandlerFromEndpoint(ctx, mux, grpcEndpoint, opts)
	if err != nil {
		return fmt.Errorf("%s: failed to register handler: %w", op, err)
	}

	log.Info("Starting HTTP server is running", slog.String("port", strconv.Itoa(a.port)))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", a.port), mux); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *App) Stop() {
	const op = "httpapp.Stop"

	a.log.With(slog.String("op", op)).Info("stopping HTTP server", slog.Int("port", a.port))
	// Implement any necessary cleanup here if needed.
}
