// Package token is used for creating tokens.
package token

import (
	"crypto/rand"
	"encoding/base64"
)

// Creates a new token.
func New(prefix string) string {
	b := make([]byte, 32)
	rand.Read(b)
	return (base64.URLEncoding.EncodeToString([]byte(prefix)) + base64.URLEncoding.EncodeToString(b))
}
