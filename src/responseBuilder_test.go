package src

import "testing"

func TestBuildURLResponse(t *testing.T) {
	got := BuildURLResponse("abc", "test.com")
	want := &URLResponse{
		ShortURL:    "abc",
		OriginalURL: "test.com",
	}

	if got[ShortURL] != want[ShortURL] {
		t.Errorf("got %q, wanted %q", got[ShortURL], want[ShortURL])
	}
	if got[OriginalURL] != want[OriginalURL] {
		t.Errorf("got %q, wanted %q", got[OriginalURL], want[OriginalURL])
	}
}
