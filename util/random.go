package util

import (
	rc "crypto/rand"
	"encoding/base64"
	"math/rand"
)

func RandomString(length int) string {
	b := make([]byte, length)
	rc.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}

func RandomNumber(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func RandomCurrency() string {
	currencies := []string{"USD", "VND", "CAD"}
	return currencies[RandomNumber(0, len(currencies)-1)]
}
