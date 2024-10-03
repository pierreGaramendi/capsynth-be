package subtitles

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
)

type Subtitle struct {
	Text  string `xml:",chardata"`
	Start string `xml:"start,attr"`
	Dur   string `xml:"dur,attr"`
}

type TextTag struct {
	Text string `xml:",chardata"`
}
type Transcript struct {
	XMLName xml.Name  `xml:"transcript"`
	Text    []TextTag `xml:"text"`
}

func getSubtitleURLByLang(videoID string, language string) (string, error) {
	// Obtener el HTML del video
	resp, err := http.Get("https://www.youtube.com/watch?v=" + videoID)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Leer todo el HTML en memoria usando io.ReadAll
	htmlBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	html := string(htmlBytes)

	// Usar regex para obtener la lista de subtítulos
	re := regexp.MustCompile(`"captions":\{"playerCaptionsTracklistRenderer":\{"captionTracks":(\[.*?\])`)
	matches := re.FindStringSubmatch(html)

	if len(matches) < 2 {
		return "", fmt.Errorf("no se encontraron subtítulos para este video")
	}

	// Extraer la lista de subtítulos como un string en formato JSON
	subtitlesList := matches[1]

	// Usar regex para obtener las URL de subtítulos y los códigos de idioma
	reSub := regexp.MustCompile(`\{"baseUrl":"(.*?)","name":\{"simpleText":".*?"\},"vssId":".*?","languageCode":"(.*?)"`)
	subtitleMatches := reSub.FindAllStringSubmatch(subtitlesList, -1)

	// Buscar la URL del subtítulo en el idioma especificado
	for _, match := range subtitleMatches {
		if len(match) < 3 {
			continue
		}
		subtitleURL := strings.ReplaceAll(match[1], `\u0026`, "&")
		fmt.Println("subtitleURL:		", subtitleURL)
		subtitleLang := match[2]
		fmt.Println("subtitleLang:		", subtitleLang)
		fmt.Println("subtitleLang:		", language)
		if subtitleLang == language {
			return subtitleURL, nil
		}
	}

	return "", fmt.Errorf("no se encontró un subtítulo en el idioma especificado: %s", language)
}

// Descargar y procesar los subtítulos en memoria
func downloadAndParseSubtitles(subtitleURL string) (*Transcript, error) {
	// Descargar subtítulos
	resp, err := http.Get(subtitleURL)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Leer todo el contenido en memoria usando `io.ReadAll`
	subtitlesBytes, err := io.ReadAll(resp.Body)
	//fmt.Println("subtitlesBytes:		", subtitlesBytes)
	//fmt.Println("XML Content:", string(subtitlesBytes))
	if err != nil {
		return nil, err
	}

	// Analizar los subtítulos en XML
	var transcript Transcript
	if err := xml.Unmarshal(subtitlesBytes, &transcript); err != nil {
		return nil, err
	}
	//fmt.Println("subtitles:		", subtitles)
	return &transcript, nil
}
