package url_test

import (
	"testing"

	"github.com/brkcnr/url-shortener/internal/url"
)

func TestShorten_ReturnsExpectedCharacterLimit(t *testing.T) {
	originalUrl := "https://www.example.com"
	shortUrl := url.Shorten(originalUrl)
	if len(shortUrl) != 8 {
		t.Errorf("Expected short URL length of 8, but got %d", len(shortUrl))
	}
}

func TestShorten_EmptyString(t *testing.T) {
    result := url.Shorten("")
    expected := "e3b0c442"
    if result != expected {
        t.Errorf("Expected %s, but got %s", expected, result)
    }
}

func TestShorten_DifferentInputs(t *testing.T) {
    url1 := "https://example.com/first"
    url2 := "https://example.com/second"

    shortURL1 := url.Shorten(url1)
    shortURL2 := url.Shorten(url2)

    if shortURL1 == shortURL2 {
        t.Errorf("Expected different short URLs for different inputs, got %s and %s", shortURL1, shortURL2)
    }
}

func TestShorten_ConsistentOutputForSameInput(t *testing.T) {
    input := "https://www.example.com"
    shortURL1 := url.Shorten(input)
    shortURL2 := url.Shorten(input)

    if shortURL1 != shortURL2 {
        t.Errorf("Expected the same short URL for the same input, but got %s and %s", shortURL1, shortURL2)
    }
}
func TestShorten_LongURL(t *testing.T) {
    longURL := "https://www.example.com/" + string(make([]byte, 1000))
    shortURL := url.Shorten(longURL)

    if len(shortURL) != 8 {
        t.Errorf("Expected short URL length of 8, but got %d", len(shortURL))
    }
}

func TestShorten_URLWithSpecialCharacters(t *testing.T) {
    specialCharURL := "https://www.example.com/path?query=param&another=param#fragment"
    shortURL := url.Shorten(specialCharURL)

    if len(shortURL) != 8 {
        t.Errorf("Expected short URL length of 8, but got %d", len(shortURL))
    }
}


