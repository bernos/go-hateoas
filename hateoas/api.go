package hateoas

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"path"
	"reflect"
	"strings"
)

type Api interface {
	http.Handler
	BasePath() string
	AddResourceHandler(string, ResourceHandler) Api
	LinkTo(ResourceHandler) *Link
}

type httprouterApi struct {
	handlers map[string]ResourceHandler
	router   *httprouter.Router
	basePath string
}

func NewApi(basePath string) *httprouterApi {
	api := &httprouterApi{
		router:   httprouter.New(),
		handlers: make(map[string]ResourceHandler),
		basePath: basePath,
	}

	// Add api to package api collection so that we
	// can look up handlers in hateoas.LinkTo()
	apis = append(apis, api)

	return api
}

func (api *httprouterApi) BasePath() string {
	return api.basePath
}

func (api *httprouterApi) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if api.isMethodSupported(r.Method) {
		h, p, _ := api.router.Lookup(r.Method, r.URL.Path)

		if h == nil {
			SendNotFound(w)
		} else {
			h(w, r, p)
		}
	} else {
		SendMethodNotAllowed(w)
	}
}

func (api *httprouterApi) AddResourceHandler(p string, handler ResourceHandler) Api {
	api.handlers[p] = handler

	collectionPath := path.Join(api.basePath, p)
	resourcePath := path.Join(collectionPath, ":id")

	api.router.GET(resourcePath, api.buildGetHandler(handler))
	api.router.GET(collectionPath, api.buildIndexHandler(handler))
	api.router.POST(collectionPath, api.buildPostHandler(handler))
	api.router.PUT(resourcePath, api.buildPutHandler(handler))
	api.router.PATCH(resourcePath, api.buildPatchHandler(handler))
	api.router.DELETE(resourcePath, api.buildDeleteHandler(handler))

	return api
}

func (api *httprouterApi) LinkTo(h ResourceHandler) *Link {
	t := reflect.TypeOf(h)
	for p, r := range api.handlers {
		if reflect.TypeOf(r) == t {
			return &Link{
				Rel:  "self",
				Href: path.Join(api.basePath, p),
			}
		}
	}
	return nil
}

func (api *httprouterApi) buildGetHandler(handler ResourceHandler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		id := p.ByName("id")

		if len(id) == 0 {
			SendBadRequest(w, "id parameter not supplied")
		} else {
			resource, status := handler.Get(id, r)
			Send(w, resource, status)
		}
	}
}

func (api *httprouterApi) buildIndexHandler(handler ResourceHandler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		resources, status := handler.Index(r)
		Send(w, NewResourceCollection(resources), status)
	}
}

func (api *httprouterApi) buildPostHandler(handler ResourceHandler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		resource, status := handler.Post(r)

		Send(w, resource, status)

		// TODO: Set location header
	}
}

func (api *httprouterApi) buildPutHandler(handler ResourceHandler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		id := p.ByName("id")

		if len(id) == 0 {
			SendBadRequest(w, "id parameter not supplied")
		} else {
			resource, status := handler.Put(id, r)
			Send(w, resource, status)
		}
	}
}

func (api *httprouterApi) buildPatchHandler(handler ResourceHandler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		id := p.ByName("id")

		if len(id) == 0 {
			SendBadRequest(w, "id parameter not supplied")
		} else {
			resource, status := handler.Patch(id, r)
			Send(w, resource, status)
		}
	}
}

func (api *httprouterApi) buildDeleteHandler(handler ResourceHandler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		id := p.ByName("id")

		if len(id) == 0 {
			SendBadRequest(w, "id parameter not supplied")
		} else {
			resource, status := handler.Delete(id, r)
			Send(w, resource, status)
		}
	}
}

func (api *httprouterApi) isMethodSupported(method string) bool {
	m := strings.ToLower(method)
	return m == "get" ||
		m == "post" ||
		m == "put" ||
		m == "patch" ||
		m == "delete"
}
