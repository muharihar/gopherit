package adapter

import (
	"00-newapp-template/internal/pkg"
	"00-newapp-template/pkg/acme"
	"strings"
)

type Filter struct {
	Config *pkg.Config
}

type GopherFilter struct {
	Gophers []acme.Gopher
}
type ThingFilter struct {
	Things []acme.Thing
}

func NewFilter(config *pkg.Config) (filter *Filter) {
	filter = new(Filter)
	filter.Config = config
	return
}

func (f *Filter) Gophers(in []acme.Gopher) (out []acme.Gopher) {
	gopherID := f.Config.Client.GopherID
	name := f.Config.Client.GopherName

	// IDs can be comma separated and don't guaranted just ONE match.
	if gopherID != "" {
		gg := strings.Split(gopherID, ",")
		for i := range gg {
			filter := &GopherFilter{Gophers: in}
			out = append(out, filter.ID(gg[i])...)
		}
		if len(out) == 0 {
			return
		}
	} else {
		out = in
	}

	if name != "" {
		filter := &GopherFilter{Gophers: out}
		out = filter.Name(name)
		if len(out) == 0 {
			return
		}
	}

	return
}

func (f *Filter) Things(in []acme.Thing) (out []acme.Thing) {
	thingID := f.Config.Client.ThingID
	name := f.Config.Client.ThingName
	gopherID := f.Config.Client.GopherID

	// IDs can be comma separated and don't guaranted just ONE match.
	if thingID != "" {
		gg := strings.Split(thingID, ",")
		for i := range gg {
			filter := &ThingFilter{Things: in}
			out = append(out, filter.ID(gg[i])...)
		}
		if len(out) == 0 {
			return
		}
	} else {
		out = in
	}

	if gopherID != "" {
		gg := strings.Split(thingID, ",")
		for i := range gg {
			filter := &ThingFilter{Things: out}
			out = append(out, filter.GopherID(gg[i])...)
		}
		if len(out) == 0 {
			return
		}
	}

	if name != "" {
		filter := &ThingFilter{Things: out}
		out = filter.Name(name)
	}

	return
}

func (gf *GopherFilter) ID(id string) (out []acme.Gopher) {
	for i := range gf.Gophers {
		if string(gf.Gophers[i].ID) == id {
			out = append(out, gf.Gophers[i])
		}
	}
	return
}
func (gf *GopherFilter) Name(name string) (out []acme.Gopher) {
	for i := range gf.Gophers {
		if strings.Contains(strings.ToLower(gf.Gophers[i].Name), strings.ToLower(name)) {
			out = append(out, gf.Gophers[i])
		}
	}
	return
}
func (gf *ThingFilter) ID(id string) (out []acme.Thing) {
	for _, t := range gf.Things {
		if string(t.ID) == id {
			out = append(out, t)
		}
	}
	return
}
func (gf *ThingFilter) GopherID(id string) (out []acme.Thing) {
	for _, t := range gf.Things {
		if string(t.GopherID) == id {
			out = append(out, t)
		}
	}
	return
}
func (gf *ThingFilter) Name(name string) (out []acme.Thing) {
	for i := range gf.Things {
		if strings.Contains(strings.ToLower(gf.Things[i].Name), strings.ToLower(name)) {
			out = append(out, gf.Things[i])
		}
	}
	return
}
