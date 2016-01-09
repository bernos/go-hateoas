package hateoas

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Api struct {
	handlers map[string]ResourceHandler
}

func NewApi() *Api {
	return &Api{}
}

func (api *Api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler := api.ResolveResourceHandler(r.URL.Path)

	fmt.Printf("Request for %s, %s", r.RequestURI, r.URL.Path)

	if handler == nil {
		http.NotFound(w, r)
	} else {
		r, s := handler.Get("asdf")

		w.WriteHeader(s)

		e := json.NewEncoder(w)
		e.Encode(r)

	}
}

func (api *Api) AddResourceHandler(path string, handler ResourceHandler) *Api {
	return api
}

func (api *Api) ResolveResourceHandler(path string) ResourceHandler {
	return nil
}
