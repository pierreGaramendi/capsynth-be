package subtitles

import (
	"encoding/json"
	"net/http"
)

type Message struct {
	Message string `json:"message"`
}

func WelcomeController(w http.ResponseWriter, r *http.Request) {
	response := Message{Message: "Hello, Cruel World!"}
	getSubtitles("https://www.youtube.com/watch?v=jaiMvRLyGRM", "es")
	json.NewEncoder(w).Encode(response)
}
