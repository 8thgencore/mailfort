package grpc

import (
	"fmt"
	"net"

	"log/slog"

	"github.com/8thgencore/mailfort/internal/delivery"
	"github.com/8thgencore/mailfort/internal/service"
	mailpb "github.com/8thgencore/mailfort/protos/gen/go/mail/v1"
	"google.golang.org/grpc"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

// New creates new gRPC server app.
func New(
	log *slog.Logger,
	mailService service.MailService,
	port int,
) *App {
	gRPCServer := grpc.NewServer()

	// Create a new MailServiceServer instance
	mailHandler := handler.NewMailServiceServer(mailService)
	mailpb.RegisterMailServiceServer(gRPCServer, mailHandler)

	return &App{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
	}
}

// MustRun runs gRPC server and panics if any error occurs.
func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	const op = "delivery.grpc.app.Run"

	log := a.log.With(
		slog.String("op", op),
		slog.Int("port", a.port),
	)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	log.Info("grpc server is running", slog.String("addr", l.Addr().String()))

	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *App) Stop() {
	const op = "grpcapp.Stop"

	a.log.With(slog.String("op", op)).Info("stopping gRPC server", slog.Int("port", a.port))

	a.gRPCServer.GracefulStop()
}
