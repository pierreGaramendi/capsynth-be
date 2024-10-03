package subtitles

import (
	"fmt"
	"log"
	"os/exec"
	"time"
)

func getSubtitles(videoUrl string, lang string) {
	start := time.Now() // Capturar el tiempo de inicio
	//cmd := exec.Command("yt-dlp", "--list-subs", "https://www.youtube.com/watch?v=jaiMvRLyGRM")
	cmd := exec.Command("yt-dlp", "--write-auto-subs", "--sub-lang", lang,
		"--sub-format", "ttml",
		"--skip-download", "-o video", videoUrl)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error al ejecutar yt-dlp: %v", err)
	}

	fmt.Printf("Subtítulos disponibles:\n%s", string(out))
	duration := time.Since(start) // Calcular el tiempo transcurrido
	fmt.Printf("El programa tardó %v en completarse.\n", duration)
}
