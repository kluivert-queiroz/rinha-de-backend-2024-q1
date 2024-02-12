package http

import "github.com/gofiber/fiber/v2"

type Server struct {
	fiber *fiber.App
}

func NewServer() *Server {
	app := fiber.New()
	server := Server{
		fiber: app,
	}
	server.initRoutes()
	return &server
}

func (s *Server) Start() error {
	return s.fiber.Listen(":3000")
}
