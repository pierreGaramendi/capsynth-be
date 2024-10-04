package subtitles

import (
	"capsynth/constants"
	"context"
	"log"
	"os"

	cohere "github.com/cohere-ai/cohere-go/v2"
	client "github.com/cohere-ai/cohere-go/v2/client"
	"github.com/joho/godotenv"
)

func AskCohereWithOwnPackage(subtitles string, lang string) string {
	myprompt := constants.ComposePrompt(subtitles, lang)
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error cargando el archivo .env: %v", err)
	}
	apiKey := os.Getenv("CO_API_KEY")
	if apiKey == "" {
		log.Println("The CO_API_KEY environment variable is not defined.")
	}
	co := client.NewClient(client.WithToken(apiKey))
	resp, err := co.Chat(
		context.TODO(),
		&cohere.ChatRequest{
			Message: myprompt,
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	//log.Printf("%+v", resp)
	return resp.Text
}
