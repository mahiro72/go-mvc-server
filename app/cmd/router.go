package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/mahiro72/go-mvc-server/pkg/controller"
)

type IRouter interface {
	Serve() error
	SetMiddlewares()
}

type ChiRouter struct {
	mux  *chi.Mux
	port string
}

func NewChiRouterImpl(port string) IRouter {
	return &ChiRouter{
		mux:  chi.NewRouter(),
		port: fmt.Sprintf(":%s", port),
	}
}

func (r *ChiRouter) SetMiddlewares() {
	r.setMiddlewares(middleware.Logger, middleware.Recoverer)
}

func (r *ChiRouter) setMiddlewares(middlewares ...func(next http.Handler) http.Handler) {
	for _, middleware := range middlewares {
		r.mux.Use(middleware)
	}
}

func (r *ChiRouter) Serve() error {
	r.registerRouters(r.userRouter, r.healthRouter)
	return http.ListenAndServe(r.port, r.mux)
}

func (r *ChiRouter) registerRouters(routers ...func()) {
	for _, initRouter := range routers {
		initRouter()
	}
}

func (r *ChiRouter) userRouter() {
	r.mux.Route("/users", func(r chi.Router) {
		r.Get("/{id}", controller.GetUser)
		r.Post("/", controller.CreateUser)
	})
}

func (r *ChiRouter) healthRouter() {
	r.mux.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})
}
