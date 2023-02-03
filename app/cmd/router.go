package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mahiro72/go-mvc-server/pkg/controller"
)

type Router struct {
	mux *chi.Mux
	port string
}

type IRouter interface {
	InitRouter() error
}

func NewChiRouter(port string) IRouter {
	return &Router{
		mux: chi.NewRouter(),
		port: fmt.Sprintf(":%s",port),
	}
}

func (r *Router) InitRouter() error {
	r.registerRouters(r.userRouter,r.healthRouter)
	return http.ListenAndServe(r.port,r.mux)
}

func (r *Router) registerRouters(routers ...func()){
	for _,initRouter := range routers {
		initRouter()
	}
}

func (r *Router) userRouter() {
	r.mux.Route("/users",func(r chi.Router) {
		r.Get("/{id}",controller.GetUser)
	})
}

func (r *Router) healthRouter(){
	r.mux.Get("/health",func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})
}
