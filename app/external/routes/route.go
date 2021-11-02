package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/tuanna7593/gosample/app/interface/restapi/handler"
)

func Handler() http.Handler {
	r := chi.NewRouter()

	// base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// init handler
	itemHandler := handler.NewItemHandler()

	r.Route("/items", func(r chi.Router) {
		r.Post("/", itemHandler.Create)
		r.Post("/{item_id}", itemHandler.BuyItem)
		r.Get("/", itemHandler.List)
	})

	return r
}
