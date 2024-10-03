// Package main is the entry point of the application that starts the HTTP server
// and handles routes to serve the subtitle fetching functionality.
package main

import (
	"capsynth/middleware"
	"capsynth/subtitles"
	"fmt"
	"net/http"
)

func main() {
	// Register the "/subtitles" route and apply the ContentTypeMiddleware.
	// The route is handled by the SubtitleController from the subtitles package.
	http.HandleFunc("/subtitles", middleware.ContentTypeMiddleware((subtitles.SubtitleController)))
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
