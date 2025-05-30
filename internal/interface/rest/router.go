package rest

import "net/http"

func registerRoutes(
	r *http.ServeMux,
	handler *HandlerPackageManager,
) {
	r.HandleFunc("POST /packages}", handler.calculatePackages)
}
