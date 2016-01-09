package hateoas

import (
	"net/http"
)

type ResourceHandler interface {
	Get(id string) (Resource, int)
	Post() (Resource, int)
	Put(id string) (Resource, int)
	Delete(id string) int
	Patch(id string) (Resource, int)
}

type GetNotSupported struct{}

func (r *GetNotSupported) Get(id string) (Resource, int) {
	return nil, http.StatusMethodNotAllowed
}

type PostNotSupported struct{}

func (r *PostNotSupported) Post() (Resource, int) {
	return nil, http.StatusMethodNotAllowed
}

type PutNotSupported struct{}

func (r *PutNotSupported) Put(id string) (Resource, int) {
	return nil, http.StatusMethodNotAllowed
}

type DeleteNotSupported struct{}

func (r *DeleteNotSupported) Delete(id string) int {
	return http.StatusMethodNotAllowed
}

type PatchNotSupported struct{}

func (r *PatchNotSupported) Patch(id string) (Resource, int) {
	return nil, http.StatusMethodNotAllowed
}
