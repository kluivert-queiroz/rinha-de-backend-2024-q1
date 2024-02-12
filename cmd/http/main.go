package main

import (
	"github.com/kluivert-queiroz/rinha-de-backend-2024-q1/internal/infra/http"
)

func main() {
	server := http.NewServer()
	server.Start()
}
