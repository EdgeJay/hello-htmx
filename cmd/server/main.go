package main

import "github.com/EdgeJay/hello-htmx/web"

func main() {
	router := web.NewRouter()
	router.SetupRoutes()
	router.Start()
}
