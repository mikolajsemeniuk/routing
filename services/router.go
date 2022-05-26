package services

import (
	"context"
	"net/http"
	"regexp"
)

type ParamsKey string

const ContextParamsKey ParamsKey = "params"

type Router interface {
	Route() error
}

type Route struct {
	Method  string
	Pattern string
	Handler http.HandlerFunc
}

type router struct {
	routes []Route
}

func (r router) Route() error {
	http.HandleFunc("/", r.Handler)
	return http.ListenAndServe(":8080", nil)
}

func (r router) Handler(writer http.ResponseWriter, request *http.Request) {
	url := request.URL.Path

	for _, route := range r.routes {

		re := regexp.MustCompile(route.Pattern)
		matches := re.FindAllStringSubmatch(url, -1)

		if len(matches) > 0 && request.Method != route.Method {
			writer.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		if len(matches) > 0 && request.Method == route.Method {
			params := matches[0][1:]
			newContext := context.WithValue(request.Context(), ContextParamsKey, params)
			route.Handler(writer, request.WithContext(newContext))
			return
		}
	}

	writer.WriteHeader(http.StatusNotFound)
}

func NewRouter(routes []Route) Router {
	return router{
		routes: routes,
	}
}
