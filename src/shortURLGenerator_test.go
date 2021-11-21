package src

import "testing"

func TestGenerateShortURL(t *testing.T) {
	url := GenerateShortURL()
	if url == nil {
		t.Errorf("URL Got Nil But expected Non Nil")
	}
}
