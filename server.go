package main

import (
	"capsynth/middleware"
	"capsynth/subtitles"
	"net/http"
)

func main() {
	http.HandleFunc("/subtitles", middleware.ContentTypeMiddleware((subtitles.WelcomeController)))
	http.ListenAndServe(":8080", nil)
}
