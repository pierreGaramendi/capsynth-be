package subtitles

import (
	"capsynth/constants"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
)

type TextTag struct {
	Text string `xml:",chardata"`
}
type Transcript struct {
	XMLName xml.Name  `xml:"transcript"`
	Text    []TextTag `xml:"text"`
}

func getSubtitleUrlByLang(videoID string, language string) (string, error) {
	url := constants.BaseYouTubeURL + videoID

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	htmlBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	html := string(htmlBytes)
	re := regexp.MustCompile(`"captions":\{"playerCaptionsTracklistRenderer":\{"captionTracks":(\[.*?\])`)
	matches := re.FindStringSubmatch(html)
	if len(matches) < 2 {
		return "", fmt.Errorf("no subtitles were found for this video")
	}
	subtitlesList := matches[1]
	reSub := regexp.MustCompile(`\{"baseUrl":"(.*?)","name":\{"simpleText":".*?"\},"vssId":".*?","languageCode":"(.*?)"`)
	subtitleMatches := reSub.FindAllStringSubmatch(subtitlesList, -1)

	for _, match := range subtitleMatches {
		if len(match) < 3 {
			continue
		}
		subtitleURL := strings.ReplaceAll(match[1], `\u0026`, "&")
		subtitleLang := match[2]
		if strings.HasPrefix(subtitleLang, language) {
			return subtitleURL, nil
		}
	}
	return "", fmt.Errorf("no subtitle was found in the specified language: %s", language)
}

func downloadAndParseSubtitles(subtitleURL string) (string, error) {
	resp, err := http.Get(subtitleURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	subtitlesBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var transcript Transcript
	if err := xml.Unmarshal(subtitlesBytes, &transcript); err != nil {
		return "", err
	}
	var plainText strings.Builder
	for _, text := range transcript.Text {
		plainText.WriteString(text.Text + " ")
	}
	return plainText.String(), nil
}
