package src

import (
	"testing"
)

func TestGenerateShortURL(t *testing.T) {
	url := GenerateShortURL()
	var length = len([]rune(url))
	if length != 8 {
		t.Errorf("URL Got Nil But expected Non Nil")
	}
}
