package main

import (
	"fmt"
)

func main() {
	port := "8080"
	r := NewChiRouterImpl(port)
	r.SetMiddlewares()
	if err := r.Serve(); err != nil {
		panic(fmt.Sprintf("error: failed to InitRouter: %s", err.Error()))
	}
}
