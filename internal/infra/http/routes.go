package http

import (
	"github.com/kluivert-queiroz/rinha-de-backend-2024-q1/internal/application/services"
	"github.com/kluivert-queiroz/rinha-de-backend-2024-q1/internal/infra/db/repositories"
	"github.com/kluivert-queiroz/rinha-de-backend-2024-q1/internal/infra/http/api/rest"
	"github.com/kluivert-queiroz/rinha-de-backend-2024-q1/pkg/mongo"
)

func (s *Server) initRoutes() {
	mongo := mongo.NewMongoClient()
	r := repositories.NewCustomerMongoRepository(*mongo)
	t := services.NewTransactionService(r)
	customerController := rest.NewCustomerController(t)
	s.fiber.Post("/clientes/:id/transacoes", customerController.Transaction)
	s.fiber.Get("/clientes/:id/extrato", customerController.Statements)
}
