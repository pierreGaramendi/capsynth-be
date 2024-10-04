package subtitles

import (
	"capsynth/constants"
	"capsynth/helpers"
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
	subtitleURL, err := getSubtitleUrlByLang(videoID, lang)
	if err != nil {
		helpers.ErrorResponse(w, err.Error(), http.StatusNotFound)
		return
	}
	subtitles, err := downloadAndParseSubtitles(subtitleURL)
	if err != nil {
		helpers.ErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	synthesis := AskCohereWithOwnPackage(subtitles, lang)
	response := Message{Message: synthesis}
	json.NewEncoder(w).Encode(response)
}
