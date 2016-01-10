package hateoas

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strings"
)

type Api struct {
	handlers map[string]ResourceHandler
	router   *httprouter.Router
}

func NewApi() *Api {
	return &Api{
		router: httprouter.New(),
	}
}

func (api *Api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if api.isMethodSupported(r.Method) {
		h, p, _ := api.router.Lookup("get", "foo")

		if h == nil {
			api.sendNotFound(w)
		} else {
			h(w, r, p)
		}
	} else {
		api.sendMethodNotAllowed(w)
	}
}

func (api *Api) AddResourceHandler(path string, handler ResourceHandler) *Api {
	api.router.GET(path+"/:id", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		id := p.ByName("id")
		resource, status := handler.Get(id)
		api.send(w, resource, status)
	})
	return api
}

func (api *Api) resolveResourceHandler(path string) ResourceHandler {
	return nil
}

func (api *Api) send(w http.ResponseWriter, resource Resource, status int) {
	e := json.NewEncoder(w)
	err := e.Encode(resource)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(status)
	}
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
