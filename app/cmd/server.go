package main

import (
	"fmt"
)

func main() {
	port := "8080"
	r := NewChiRouterImpl(port)
	if err := r.Serve(); err != nil {
		panic(fmt.Sprintf("error: failed to InitRouter: %s", err.Error()))
	}
}
