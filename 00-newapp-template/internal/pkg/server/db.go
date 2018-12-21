package server

import (
	"00-newapp-template/pkg/acme"
)

type SimpleDB struct {
	gophers []acme.Gopher
	things  []acme.Thing
}

func NewDB() (s SimpleDB) {
	s.gophers = []acme.Gopher{
		{ID: "1", Name: "Gopher1", Description: "The first Gopher (#1st)"},
		{ID: "2", Name: "Gopher2", Description: "The second Gopher (#2nd)"},
		{ID: "4", Name: "Gopher4", Description: "The fourth Gopher (#4th)"},
		{ID: "8", Name: "Gopher8", Description: "The eighth Gopher (#8th)"},
	}

	s.things = []acme.Thing{
		{ID: "1", GopherID: "1", Name: "Head", Description: "Hat"},
		{ID: "2", GopherID: "2", Name: "Head", Description: "Hat"},
		{ID: "3", GopherID: "4", Name: "Head", Description: "Hat"},
		{ID: "4", GopherID: "8", Name: "Head", Description: "Hat"},

		{ID: "5", GopherID: "1", Name: "Feet", Description: "Shoes"},
		{ID: "6", GopherID: "2", Name: "Feet", Description: "Shoes"},
		{ID: "7", GopherID: "4", Name: "Feet", Description: "Shoes"},
		{ID: "8", GopherID: "8", Name: "Feet", Description: "Shoes"},

		{ID: "9", GopherID: "1", Name: "Waist", Description: "Belt"},
		{ID: "10", GopherID: "2", Name: "Waist", Description: "Belt"},
		{ID: "11", GopherID: "4", Name: "Waist", Description: "Belt"},
		{ID: "12", GopherID: "8", Name: "Waist", Description: "Belt"},
	}
	return
}

func (s *SimpleDB) Gophers() []acme.Gopher {
	return s.gophers
}
func (s *SimpleDB) GopherThings(gopherID string) (things []acme.Thing) {
	for _, v := range s.things {
		if string(v.GopherID) == gopherID {
			things = append(things, v)
		}
	}
	return
}

// DeleteGopher 'cascade deleted' from gophers and things.
func (s *SimpleDB) DeleteGopher(gopherID string) {
	var gophers []acme.Gopher
	var things []acme.Thing

	for _, g := range s.gophers {
		if string(g.ID) == gopherID {
			continue
		}
		gophers = append(gophers, g)
	}
	s.gophers = gophers

	for _, t := range s.things {
		if string(t.GopherID) == gopherID {
			continue
		}
		things = append(things, t)
	}
	s.things = things
}

func (s *SimpleDB) DeleteThing(gopherID string, thingID string) {
	var things []acme.Thing

	for _, t := range s.things {
		if string(t.GopherID) == gopherID && string(t.ID) == thingID {
			continue
		}
		things = append(things, t)
	}

	s.things = things
}

func (s *SimpleDB) UpdateGopher(newGopher acme.Gopher) {
	var gophers []acme.Gopher
	for _, g := range s.gophers {
		if string(newGopher.ID) == string(g.ID) {
			gophers = append(gophers, newGopher)
			continue
		}
		gophers = append(gophers, g)
	}
	s.gophers = gophers
}
