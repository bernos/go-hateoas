package hateoas

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"path"
	"strings"
)

type Api struct {
	handlers map[string]ResourceHandler
	router   *httprouter.Router
	basePath string
}

func NewApi() *Api {
	return &Api{
		router:   httprouter.New(),
		handlers: make(map[string]ResourceHandler),
	}
}

func (api *Api) buildGetHandler(handler ResourceHandler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		id := p.ByName("id")
		resource, status := handler.Get(id)

		if resource != nil {
			resource.Links().Set("self", r.URL.String())
		}

		api.send(w, resource, status)
	}
}

func (api *Api) buildIndexHandler(handler ResourceHandler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		resources, status := handler.Index()
		if resources != nil {
			for i, _ := range resources {
				resources[i].Links().Set("self", "asdf")
			}
		}
		api.send(w, NewResourceCollection(resources), status)
	}
}

func (api *Api) Handler(basePath string) func(w http.ResponseWriter, r *http.Request) {
	// Set up router
	api.basePath = basePath

	for p, h := range api.handlers {
		hp := path.Join(basePath, p)
		api.router.GET(path.Join(hp, ":id"), api.buildGetHandler(h))
		api.router.GET(hp, api.buildIndexHandler(h))
	}

	return func(w http.ResponseWriter, r *http.Request) {
		if api.isMethodSupported(r.Method) {
			h, p, _ := api.router.Lookup(r.Method, r.URL.Path)

			if h == nil {
				api.sendNotFound(w)
			} else {
				h(w, r, p)
			}
		} else {
			api.sendMethodNotAllowed(w)
		}
	}
}

func (api *Api) AddResourceHandler(path string, handler ResourceHandler) *Api {
	api.handlers[path] = handler
	return api
}

func (api *Api) resolveResourceHandler(path string) ResourceHandler {
	return nil
}

func (api *Api) send(w http.ResponseWriter, resource Resource, status int) {
	w.WriteHeader(status)
	e := json.NewEncoder(w)
	e.Encode(resource)
}

func (api *Api) isMethodSupported(method string) bool {
	m := strings.ToLower(method)
	return m == "get" ||
		m == "post" ||
		m == "put" ||
		m == "patch" ||
		m == "delete"
}

func (api *Api) sendNotFound(w http.ResponseWriter) {
	resource, status := NotFound()
	api.send(w, resource, status)
}

func (api *Api) sendMethodNotAllowed(w http.ResponseWriter) {
	resource, status := MethodNotAllowed()
	api.send(w, resource, status)
}
