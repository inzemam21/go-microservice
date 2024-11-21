package server

import (
	"go-microservice/internal/database"
	"go-microservice/internal/models"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Server interface {
	Start() error
	Readiness(ctx echo.Context) error
	Liveness(ctx echo.Context) error
}

type EchoServer struct {
	echo *echo.Echo
	DB   database.DatabaseClient
}

// Liveness implements Server.
func (e *EchoServer) Liveness(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.Health{Status: "OK"})
}

// Readiness implements Server.
func (e *EchoServer) Readiness(ctx echo.Context) error {
	ready := e.DB.Ready()
	if ready {
		return ctx.JSON(http.StatusOK, models.Health{Status: "OK"})
	}

	return ctx.JSON(http.StatusInternalServerError, models.Health{Status: "Failure"})

}

// Start implements Server.
func (e *EchoServer) Start() error {
	if err := e.echo.Start(":8080"); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server shutdown occurred: %s", err)
		return err
	}

	return nil
}

func NewEchoServer(db database.DatabaseClient) Server {
	server := &EchoServer{
		echo: echo.New(),
		DB:   db,
	}
	server.registerRoutes()
	return server
}

func (e *EchoServer) registerRoutes() {
	e.echo.GET("/readiness", e.Readiness)
	e.echo.GET("/liveness", e.Liveness)
}
