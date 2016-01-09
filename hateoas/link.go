package hateoas

import (
	"encoding/json"
	"path"
)

type link struct {
	Rel  string `json:"rel"`
	Href string `json:"href"`
}

type Links struct {
	links   map[string]string
	baseUrl string
}

func (l *Links) Set(rel string, href string) *Links {
	if l.links == nil {
		l.links = make(map[string]string)
	}
	l.links[rel] = href
	return l
}

func (l *Links) SetBaseUrl(url string) *Links {
	l.baseUrl = url
	return l
}

func (l *Links) BaseUrl() string {
	return l.baseUrl
}

func (l *Links) MarshalJSON() ([]byte, error) {
	links := make([]link, len(l.links))
	i := 0
	for rel, href := range l.links {
		links[i] = link{rel, path.Join(l.baseUrl, href)}
		i++
	}
	return json.Marshal(links)
}
