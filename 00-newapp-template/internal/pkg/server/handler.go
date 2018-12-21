package server

import (
	"00-newapp-template/pkg/acme"
	"context"
	"encoding/json"
	"net/http"
	"time"
)

func (server *Server) Shutdown(w http.ResponseWriter, r *http.Request) {
	server.Log.Debugf("/shutdown called - beginning server shutdown")

	w.Write([]byte("bye felcia"))
	timeout, cancel := context.WithTimeout(server.Context, 5*time.Second)
	err := server.HTTP.Shutdown(timeout)
	if err != nil {
		server.Log.Errorf("server error during shutdown: %+v", err)
	}
	server.Finished()
	cancel()
}

func (server *Server) Gophers(w http.ResponseWriter, r *http.Request) {

	gophers := server.DB.Gophers()
	b, err := json.Marshal(gophers)
	if err != nil {
		server.Log.Errorf("error marshing gophers: %+v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Write(b)
}
func (server *Server) Gopher(w http.ResponseWriter, r *http.Request) {
	gopherID := GopherID(r)

	for _, gopher := range server.DB.Gophers() {
		if string(gopher.ID) == gopherID {
			b, err := json.Marshal(gopher)
			if err != nil {
				server.Log.Errorf("error marshing gopher: %+v", err)
				return
			}
			w.Write(b)
			return
		}
	}
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

func (server *Server) Things(w http.ResponseWriter, r *http.Request) {
	gopherID := GopherID(r)

	things := server.DB.GopherThings(gopherID)
	b, err := json.Marshal(things)
	if err != nil {
		server.Log.Errorf("error marshing things: %+v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Write(b)
}
func (server *Server) Thing(w http.ResponseWriter, r *http.Request) {
	thingID := ThingID(r)
	gopherID := GopherID(r)

	things := server.DB.GopherThings(gopherID)
	for _, thing := range things {
		if string(thing.ID) == thingID {
			b, err := json.Marshal(thing)
			if err != nil {
				server.Log.Errorf("error marshing thing: %+v", err)
				return
			}
			w.Write(b)
			return
		}
	}
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

func (server *Server) UpdateGopher(w http.ResponseWriter, r *http.Request) {
	gopher := acme.Gopher{
		ID:          json.Number(GopherID(r)),
		Name:        GopherName(r),
		Description: GopherDescription(r),
	}

	server.DB.UpdateGopher(gopher)
	server.Gopher(w, r)

	return
}
func (server *Server) DeleteGopher(w http.ResponseWriter, r *http.Request) {
	gopherID := GopherID(r)

	server.DB.DeleteGopher(gopherID)
	server.Gophers(w, r)
	return

}

func (server *Server) UpdateThing(w http.ResponseWriter, r *http.Request) {}
func (server *Server) DeleteThing(w http.ResponseWriter, r *http.Request) {
	gopherID := GopherID(r)
	thingID := ThingID(r)

	server.DB.DeleteThing(gopherID, thingID)
	server.DB.GopherThings(gopherID)

	server.Things(w, r)
	return
}
