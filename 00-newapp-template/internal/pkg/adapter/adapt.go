package adapter

import (
	"00-newapp-template/internal/pkg"
	"sync"
)

type Adapter struct {
	Config    *pkg.Config
	Convert   Convert
	Unmarshal Unmarshal
	Filter    *Filter
	Worker    *sync.WaitGroup
}

// NewAdapater manages calls the remote services, converts the results and manages a memory/disk cache.
func NewAdapter(config *pkg.Config) (a *Adapter) {
	a = new(Adapter)
	a.Config = config
	a.Worker = new(sync.WaitGroup)
	a.Unmarshal = NewUnmarshal(config)
	a.Filter = NewFilter(config)
	a.Convert = NewConvert()

	return
}

func (a *Adapter) Gopher(gopherID string) (gopher Gopher) { return }

func (a *Adapter) UpdateGopher(newGopher Gopher) (gopher Gopher) { return }
func (a *Adapter) UpdateThing(newThing Thing) (thing Thing)      { return }

func (a *Adapter) Things(gopherID string) (things map[string]Thing) {
	rawThings := a.Unmarshal.Things(gopherID)
	filtered := a.Filter.Things(rawThings)
	things = a.Convert.Things(filtered)
	return
}

// Gopher returns all gophers with 'things' == nil
func (a *Adapter) Gophers() (gophers map[string]Gopher) {
	rawGophers := a.Unmarshal.Gophers()
	filtered := a.Filter.Gophers(rawGophers)
	gophers = a.Convert.Gophers(filtered)
	return
}

// GopherThings populates each gopher with their things
func (a *Adapter) GopherThings() (gophers map[string]Gopher) {
	gophers = make(map[string]Gopher)

	var matchOnThings = false

	if a.Config.Client.ThingID != "" || a.Config.Client.ThingName != "" || a.Config.Client.ThingDescription != "" {
		matchOnThings = true
	}

	gg := a.Gophers()
	for _, g := range gg {
		things := a.Things(g.ID)

		// If there are no 'things' for this gopher and we are filtering for a thing
		// don't add this 'gopher' to the results
		if len(things) == 0 && matchOnThings {
			continue
		}
		gophers[g.ID] = Gopher{
			ID:          g.ID,
			Name:        g.Name,
			Description: g.Description,
			Things:      things,
		}
	}
	return
}

func (a *Adapter) DeleteGopher(gopherID string) (gophers map[string]Gopher) {
	rawGophers := a.Unmarshal.DeleteGopher(gopherID)
	gophers = a.Convert.Gophers(rawGophers)
	return
}

func (a *Adapter) DeleteThing(gopherID string, thingID string) (things map[string]Thing) {
	rawThings := a.Unmarshal.DeleteThing(gopherID, thingID)
	things = a.Convert.Things(rawThings)
	return
}

func (a *Adapter) FindGopherByThing(thingID string) (gopherID string) {
	gophers := a.GopherThings()
	for g := range gophers {
		for t := range gophers[g].Things {
			if string(gophers[g].Things[t].ID) == thingID {
				gopherID = gophers[g].ID
			}
		}
	}
	return
}
