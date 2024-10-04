package constants

import (
	"fmt"
)

const (
	// promtp
	englishPrompt = `The text below should be summarized in a paragraph of no more than X words.\n
	Then, generate a list with the main ideas, each accompanied by a brief summary of its meaning in 1-2 sentences.\n
	In addition, the result must be in English.\n

	The output should follow the format:\n
	- Summary:\n
	- Main ideas:\n
	1. [Idea 1]: [Brief summary].\n
	2. [Idea 2]: [Brief summary].\n

	Text:\n
	`
)

func ComposePrompt(subtitles string, langCode string) string {
	finalLang := getLangNameByLangCode(langCode)
	multiline := fmt.Sprintf(`The text at the end must be summarized in a paragraph with a size of 30%% of the original text.
	Then, generate a list with the main ideas, each accompanied by a brief summary of its meaning in 1-2 sentences.
	In addition, the result must be in %s.

	The output should follow the format:
	- Summary:
	- Main ideas:
	1. [Idea 1]: [Brief summary].
	2. [Idea 2]: [Brief summary].

	Text: %s`, finalLang, subtitles)

	return multiline
}

func getLangNameByLangCode(params string) string {
	responseMap := map[string]string{
		"en": "english",
		"es": "spanish",
	}
	if val, exists := responseMap[params]; exists {
		return val
	}
	return "english"
}
