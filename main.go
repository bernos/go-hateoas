package main

import (
	"github.com/bernos/go-hateoas/hateoas"
	"net/http"
)

type Person struct {
	hateoas.ResourceBase
	Name string
	Age  int
}

type PersonHandler struct {
	hateoas.PostNotSupported
	hateoas.PutNotSupported
	hateoas.DeleteNotSupported
	hateoas.PatchNotSupported
}

func (p *PersonHandler) Get(id string, req *http.Request) (hateoas.Resource, int) {
	person := &Person{
		Name: id,
		Age:  27,
	}

	return person, 200
}

func (p *PersonHandler) Index(req *http.Request) ([]hateoas.Resource, int) {
	people := make([]hateoas.Resource, 1)

	people[0] = &Person{
		Name: "Brendan",
		Age:  123,
	}
	return people, http.StatusOK
}

type FooHandler struct {
	hateoas.GetNotSupported
	hateoas.PostNotSupported
	hateoas.PutNotSupported
	hateoas.PatchNotSupported
	hateoas.DeleteNotSupported
}

func main() {
	// test(&PersonHandler{})
	//	test(&FooHandler{})
	api := hateoas.NewApi()
	api.AddResourceHandler("/person", &PersonHandler{})
	http.HandleFunc("/", api.Handler("/"))
	panic(http.ListenAndServe(":8080", nil))
}
