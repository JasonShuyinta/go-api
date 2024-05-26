package handlers

import (
	"github.com/JasonShuyinta/go-api/internal/middleware"
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux) {
  //Global middleware
  r.Use(chimiddle.StripSlashes)

  r.Route("/account", func(router chi.Router) {

    //middleware for /account 
    router.Use(middleware.Authorization)

    router.Get("/coins", GetCoinBalance)
  })
}


