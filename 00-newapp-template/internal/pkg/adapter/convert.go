package adapter

import (
	"00-newapp-template/pkg/acme"
)

type Convert struct{}

func (convert *Convert) Gophers(rawGophers []acme.Gopher) (gophers map[string]Gopher) {
	gophers = make(map[string]Gopher)
	for _, g := range rawGophers {
		id := string(g.ID)
		gophers[id] = Gopher{
			ID:          id,
			Name:        g.Name,
			Description: g.Description,
		}
	}
	return
}

func (convert *Convert) Things(rawThings []acme.Thing) (things map[string]Thing) {
	things = make(map[string]Thing)
	for _, t := range rawThings {
		id := string(t.ID)
		things[id] = Thing{
			ID:          id,
			Name:        t.Name,
			Description: t.Description,
			Gopher:      Gopher{ID: string(t.GopherID)},
		}
	}
	return
}

func NewConvert() (convert Convert) {
	return
}
