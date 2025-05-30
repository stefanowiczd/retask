package rest

import "net/http"

func registerRoutes(r *http.ServeMux) {
	r.HandleFunc("POST /packages}", calculatePackages)
}
