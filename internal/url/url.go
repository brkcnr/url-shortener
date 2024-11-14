package url

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// Shorten generates a short URL for a given original URL using a SHA-256 hash.
func Shorten(originalUrl string) string {
	h := sha256.New()
	h.Write([]byte(originalUrl))
	fmt.Println(h.Sum(nil))
	hash := hex.EncodeToString(h.Sum(nil))
	ShortURL := hash[:8]

	return ShortURL
}