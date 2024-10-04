package constants

import "testing"

func TestEnglishLanguage(t *testing.T) {
	resultado := getLangNameByLangCode("en")
	esperado := "english"
	if resultado != esperado {
		t.Errorf("Resultado incorrecto: got %s, want %s", resultado, esperado)
	}
}

func TestSpanishLanguage(t *testing.T) {
	resultado := getLangNameByLangCode("es")
	esperado := "spanish"
	if resultado != esperado {
		t.Errorf("Resultado incorrecto: got %s, want %s", resultado, esperado)
	}
}
