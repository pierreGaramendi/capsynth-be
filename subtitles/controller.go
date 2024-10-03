package subtitles

import (
	"encoding/json"
	"net/http"
)

type Message struct {
	Message string `json:"message"`
}

// Error estructura para los mensajes de error en JSON
type Error struct {
	Error string `json:"error"`
}

func WelcomeController(w http.ResponseWriter, r *http.Request) {
	videoID := r.URL.Query().Get("videoID")
	lang := r.URL.Query().Get("lang")
	if videoID == "" || lang == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Error: "Missing required parameters: videoID and lang"})
		return
	}
	getSubtitles("https://www.youtube.com/watch?v="+videoID, lang)
	response := Message{Message: "Subtitles fetched successfully!"}
	json.NewEncoder(w).Encode(response)
}
