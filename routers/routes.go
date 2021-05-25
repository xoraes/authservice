package routers

import (
	"expvar"
	"github.com/go-chi/chi"
	"github.com/xoraes/dappauth/middleware"
	"net/http"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()
	// Define route and call expvar http handler
	router.Get("/debug/vars", func(w http.ResponseWriter, r *http.Request) {
		expvar.Handler().ServeHTTP(w, r)
	})
	router.Post("/signup", middleware.CreateUser)
	router.Post("/login", middleware.Authenticate)
	router.Get("/users", middleware.GetAllUser)
	router.Put("/users", middleware.UpdateUser)
	return router
}
