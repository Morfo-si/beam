package server

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/morfo-si/beam/internal/config"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

// Server interface
type Server interface {
	Start() error
	Query(c fiber.Ctx) error
	App() *fiber.App
}

type ACEServer struct {
	app *fiber.App
}

func NewACEServer() Server {

	app := fiber.New(fiber.Config{
		AppName:       config.APP_NAME,
		BodyLimit:     fiber.DefaultBodyLimit,
		ServerHeader:  config.APP_NAME,
		StrictRouting: false,
		ReadTimeout:   1 * time.Second,
		WriteTimeout:  1 * time.Second,
		IdleTimeout:   10 * time.Second,
	})

	// Middleware for logging
	app.Use(
		logger.New(logger.Config{
			Format:        "${time} [${ip}]:${port} ${status} - ${method} ${path}\n",
			TimeZone:      "UTC",
			Output:        os.Stdout,
			DisableColors: false,
		}),
		cors.New(cors.Config{
			AllowOrigins: []string{"http://localhost:3000"},
			AllowMethods: []string{"POST"},
		}),
	)

	server := &ACEServer{
		app: app,
	}

	server.app.Post("/api/v1/chat", server.Query)
	return server
}

// Start the server
func (s *ACEServer) Start() error {
	if err := s.app.Listen(":" + config.LoadConfig().Port); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server shutdown occurred: %s", err)
		return err
	}
	return nil
}

func (s *ACEServer) App() *fiber.App {
	return s.app
}
