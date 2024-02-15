package domain

import (
	"github.com/kluivert-queiroz/rinha-de-backend-2024-q1/internal/domain/entities"
	"github.com/valyala/fasthttp"
)

type CustomerRepository interface {
	FindById(ctx *fasthttp.RequestCtx, id string) (entities.Customer, error)
	Save(ctx *fasthttp.RequestCtx, c entities.Customer) error
}
