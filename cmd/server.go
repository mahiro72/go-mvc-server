package main

import (
	"fmt"
)

func main() {
	port := "8080"
	r := NewChiRouter(port)
	if err := r.InitRouter(); err != nil {
		panic(fmt.Sprintf("error: failed to InitRouter: %s",err.Error()))
	}
}
