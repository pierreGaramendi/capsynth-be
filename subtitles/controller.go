package subtitles

import (
	"capsynth/constants"
	"encoding/json"
	"net/http"
)

type Message struct {
	Message string `json:"message"`
}

// Error structure for error messages in JSON
type Error struct {
	Error string `json:"error"`
}

func WelcomeController(w http.ResponseWriter, r *http.Request) {
	videoID := r.URL.Query().Get("videoID")
	lang := r.URL.Query().Get("lang")
	if videoID == "" || lang == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Error: constants.MissingParametersError})
		return
	}
	url := constants.BaseYouTubeURL + videoID
	getSubtitles(url, lang)
	response := Message{Message: constants.SubtitlesFetchedSuccessfully}
	json.NewEncoder(w).Encode(response)
}
