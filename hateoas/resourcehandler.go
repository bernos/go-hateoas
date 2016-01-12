package hateoas

import (
	"errors"
	"net/http"
)

type ResourceHandler interface {
	Index(req *http.Request) ([]Resource, int)
	Get(id string, req *http.Request) (Resource, int)
	Post(req *http.Request) (Resource, int)
	Put(id string, req *http.Request) (Resource, int)
	Delete(id string, req *http.Request) (Resource, int)
	Patch(id string, req *http.Request) (Resource, int)
}

type GetNotSupported struct{}

func (r *GetNotSupported) Get(id string, req *http.Request) (Resource, int) {
	return MethodNotAllowed()
}

type PostNotSupported struct{}

func (r *PostNotSupported) Post(req *http.Request) (Resource, int) {
	return MethodNotAllowed()
}

type PutNotSupported struct{}

func (r *PutNotSupported) Put(id string, req *http.Request) (Resource, int) {
	return MethodNotAllowed()
}

type DeleteNotSupported struct{}

func (r *DeleteNotSupported) Delete(id string, req *http.Request) (Resource, int) {
	return MethodNotAllowed()
}

type PatchNotSupported struct{}

func (r *PatchNotSupported) Patch(id string, req *http.Request) (Resource, int) {
	return MethodNotAllowed()
}

func NotFound() (Resource, int) {
	return NewErrorResource(errors.New("The requested resource was not found")), http.StatusNotFound
}

func BadRequest(msg string) (Resource, int) {
	return NewErrorResource(errors.New(msg)), http.StatusBadRequest
}

func Error(err error) (Resource, int) {
	return NewErrorResource(err), http.StatusInternalServerError
}

func MethodNotAllowed() (Resource, int) {
	return NewErrorResource(errors.New("Method not allowed")), http.StatusMethodNotAllowed
}
