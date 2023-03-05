package handler

import (
	"fmt"
	"log"
	"net/http"

	"pos/internal/api/v1"

	"github.com/go-chi/chi"
)

func HandleHTTP(port string) {
	r := chi.NewRouter()
	r.Route("/api/v1", apiV1.HandleAPI)
	log.Printf("Starting server on port %s...", port)
	panic(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}
