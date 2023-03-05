package apiV1

import (
	"net/http"

	"github.com/go-chi/chi"
)

func HandleAPI(router chi.Router) {
	router.Route("/health", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})
	})
}
