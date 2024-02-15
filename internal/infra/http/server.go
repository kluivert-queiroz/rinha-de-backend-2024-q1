package http

import (
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/kluivert-queiroz/rinha-de-backend-2024-q1/pkg/validator"
)

type Server struct {
	fiber *fiber.App
}

func NewServer() *Server {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(422).JSON(validator.GlobalErrorHandlerResp{
				Success: false,
				Message: err.Error(),
			})
		},
	})
	server := Server{
		fiber: app,
	}
	server.initRoutes()
	return &server
}

func (s *Server) Start() error {
	return s.fiber.Listen(":3000")
}
