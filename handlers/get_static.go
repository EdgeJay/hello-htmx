package handlers

import (
	"log"
	"net/http"
)

func GetStatic(w http.ResponseWriter, r *http.Request) {
	filePath := r.PathValue("path")
	log.Printf("Static asset: Received request for /static/%s\n", filePath)
	http.ServeFile(w, r, "./static/"+filePath)
}
