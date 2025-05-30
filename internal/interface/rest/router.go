package rest

import "net/http"

func registerRoutes(
	r *http.ServeMux,
	handler *HandlerPacksManager,
) {
	r.HandleFunc("POST /packages}", handler.calculatePacks)
}
