package src

import "testing"

func TestBuildURLResponse(t *testing.T) {
	got := BuildURLResponse("localhost:8080", "abc", "http://test.com")
	want := &URLResponse{
		ShortURL:    "http://localhost:8080/abc",
		OriginalURL: "http://test.com",
	}

	if got.ShortURL != want.ShortURL {
		t.Errorf("got %q, wanted %q", got.ShortURL, want.ShortURL)
	}
	if got.OriginalURL != want.OriginalURL {
		t.Errorf("got %q, wanted %q", got.OriginalURL, want.OriginalURL)
	}
}
