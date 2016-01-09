package hateoas

type Resource interface {
	Links() *Links
}

type resourceBase struct {
	Links *Links `json:"_links,omitempty"`
}

type ResourceBase struct {
	resourceBase
}

func (r *ResourceBase) Links() *Links {
	if r.resourceBase.Links == nil {
		r.resourceBase.Links = &Links{}
	}
	return r.resourceBase.Links
}
