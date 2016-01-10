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

type ResourceCollection struct {
	ResourceBase
	Items []Resource `json:"items"`
}

func NewResourceCollection(items []Resource) *ResourceCollection {
	return &ResourceCollection{
		Items: items,
	}
}
