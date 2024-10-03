package subtitles

import (
	"capsynth/constants"
	"encoding/json"
	"net/http"
)

// Message represents a standard response message structure in JSON format.
type Message struct {
	Message string `json:"message"`
}

// Error represents an error message structure in JSON format.
type Error struct {
	Error string `json:"error"`
}

// SubtitleController handles HTTP requests for fetching subtitles from YouTube.
func SubtitleController(w http.ResponseWriter, r *http.Request) {
	videoID := r.URL.Query().Get("videoID")
	lang := r.URL.Query().Get("lang")
	if videoID == "" || lang == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Error: constants.MissingParametersError})
		return
	}
	url := constants.BaseYouTubeURL + videoID
	getSubtitlesYtdlp(url, lang)
	response := Message{Message: constants.SubtitlesFetchedSuccessfully}
	json.NewEncoder(w).Encode(response)
}
