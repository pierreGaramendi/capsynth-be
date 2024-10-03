package subtitles

import (
	"encoding/json"
	"net/http"
)

type Message struct {
	Message string `json:"message"`
}

func WelcomeController(w http.ResponseWriter, r *http.Request) {
	videoID := r.URL.Query().Get("videoID")
	lang := r.URL.Query().Get("lang")
	if videoID == "" || lang == "" {
		http.Error(w, `{"error": "Missing required parameters: videoID and lang"}`, http.StatusBadRequest)
		return
	}
	getSubtitles("https://www.youtube.com/watch?v="+videoID, lang)
	response := Message{Message: "Subtitles fetched successfully!"}
	json.NewEncoder(w).Encode(response)
}
