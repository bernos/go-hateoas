package hateoas

import (
	"encoding/json"
	"net/http"
)

var (
	apis       []Api = make([]Api, 0) // Keep a map of created apis to use when creating links
	defaultApi Api
)

func AddResourceHandler(path string, handler ResourceHandler) Api {
	return DefaultApi().AddResourceHandler(path, handler)
}

func DefaultApi() Api {
	if defaultApi == nil {
		defaultApi = NewApi("/")
	}
	return defaultApi
}

func Handler() http.Handler {
	return DefaultApi()
}

func LinkTo(h ResourceHandler) *Link {
	for _, api := range apis {
		l := api.LinkTo(h)

		if l != nil {
			return l
		}
	}
	return &Link{Rel: "self", Href: "/"}
}

func Send(w http.ResponseWriter, resource Resource, status int) {
	w.WriteHeader(status)
	e := json.NewEncoder(w)
	e.Encode(resource)
}

func SendBadRequest(w http.ResponseWriter, msg string) {
	resource, status := BadRequest(msg)
	Send(w, resource, status)
}

func SendMethodNotAllowed(w http.ResponseWriter) {
	resource, status := MethodNotAllowed()
	Send(w, resource, status)
}

func SendNotFound(w http.ResponseWriter) {
	resource, status := NotFound()
	Send(w, resource, status)
}

func UnmarshalRequestBody(r *http.Request, resource Resource) error {
	d := json.NewDecoder(r.Body)
	return d.Decode(resource)
}
