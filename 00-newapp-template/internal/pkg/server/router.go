package server

import (
	"00-newapp-template/internal/pkg/server/middleware"
	"github.com/go-chi/chi"
	chimiddleware "github.com/go-chi/chi/middleware"
)

func (server *Server) NewRouter() {
	server.Router.Use(chimiddleware.RequestID)
	server.Router.Use(middleware.NewStructuredLogger(server.Log))
	server.Router.Use(chimiddleware.Recoverer)

	server.Router.Route("/", func(r chi.Router) {
		r.Use(InitCtx)

		r.Get("/shutdown", server.Shutdown)
		r.Get("/gophers", server.Gophers)

		r.Route("/gopher", func(r chi.Router) {
			r.Route("/{GopherID}", func(r chi.Router) {
				r.Use(GopherCtx)
				r.Get("/", server.Gopher)
				r.Put("/", server.UpdateGopher)
				r.Delete("/", server.DeleteGopher)

				// Things doesn't a ThingID and therefore doesn't have a ThingCtx
				r.Get("/things", server.Things)

				r.Route("/thing/{ThingID}", func(r chi.Router) {
					r.Use(ThingCtx)
					r.Get("/", server.Thing)
					r.Put("/", server.UpdateThing)
					r.Delete("/", server.DeleteThing)
				})
			})
		})
	})
}
