package hateoas

import (
	"encoding/json"
	"path"
)

type Link struct {
	Rel  string `json:"rel"`
	Href string `json:"href"`
}

func (l *Link) Slash(s string) *Link {
	l.Href = path.Join(l.Href, s)
	return l
}

func (l *Link) WithRel(s string) *Link {
	l.Rel = s
	return l
}

type Links struct {
	links []*Link
}

func (l *Links) Add(link *Link) *Links {
	if link != nil {
		l.links = append(l.links, link)
	}
	return l
}

func (l *Links) MarshalJSON() ([]byte, error) {
	return json.Marshal(l.links)
}
