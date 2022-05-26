package main

import (
	"net/http"
	"routing/services"
	"strings"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("/home"))
	w.WriteHeader(http.StatusOK)
}

func HomeById(w http.ResponseWriter, r *http.Request) {

	routes := r.Context().Value(services.ContextParamsKey)
	re, ok := routes.([]string)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Write([]byte("/home/:id, " + strings.Join(re, ",")))
	w.WriteHeader(http.StatusOK)
}

func main() {
	routes := []services.Route{
		{
			Method:  http.MethodGet,
			Pattern: "/home/?$",
			Handler: Home,
		},
		{
			Method:  http.MethodGet,
			Pattern: "home/([^/]+)/?$",
			Handler: HomeById,
		},
	}
	router := services.NewRouter(routes)
	router.Route()

	// utworz obiekt o nazwie router
	// utworz slica structur z metodami, routem i handlerem
	// uruchom na nim metode route i przekaz slica

	// Route(http.MethodGet, "/api/home", tester)
}
