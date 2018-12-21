package adapter

import (
	"00-newapp-template/internal/pkg"
	"00-newapp-template/pkg/acme"
)

type Unmarshal struct {
	Config *pkg.Config
}

func NewUnmarshal(config *pkg.Config) (u Unmarshal) {
	u.Config = config
	return
}

func (u *Unmarshal) Service() (s acme.Service) {
	s = acme.NewService(u.Config.Client.BaseURL, u.Config.Client.SecretKey, u.Config.Client.AccessKey)
	return
}

func (u *Unmarshal) Gophers() (gophers []acme.Gopher) {
	service := u.Service()
	gophers = service.GetGophers()
	return
}

func (u *Unmarshal) Things(gopherID string) (things []acme.Thing) {
	service := u.Service()
	things = service.GetThings(gopherID)
	return
}

// DeleteGopher returns all gophers are aren't deleted.
func (u *Unmarshal) DeleteGopher(gopherID string) (gophers []acme.Gopher) {
	service := u.Service()
	gophers = service.DeleteGopher(gopherID)
	return
}

// DeleteThings returns all things for a gopher that aren't deleted.
func (u *Unmarshal) DeleteThing(gopherID string, thingID string) (things []acme.Thing) {
	service := u.Service()
	things = service.DeleteThing(gopherID, thingID)
	return
}
