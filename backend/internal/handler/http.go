package handler

import (
	"fmt"
	"net/http"

	"pos/internal/api/v1"
	"pos/utils"

	"github.com/go-chi/chi"
)

func HandleHTTP(port string) {
	r := chi.NewRouter()
	r.Route("/api/v1", apiV1.HandleAPI)
	utils.Logger.Info(fmt.Sprintf("Starting server on port %s...", port))
	panic(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}
