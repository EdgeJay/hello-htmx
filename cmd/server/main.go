package main

import "github.com/EdgeJay/hello-htmx/routers"

func main() {
	router := routers.NewRouter()
	router.SetupRoutes()
	router.Start()
}
