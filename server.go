package main

import (
	"capsynth/middleware"
	"capsynth/subtitles"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/subtitles", middleware.ContentTypeMiddleware((subtitles.WelcomeController)))
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
