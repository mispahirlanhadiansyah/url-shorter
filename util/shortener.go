package util

import (
	"math/rand"
	"regexp"
	"time"
)

var regex = "^[0-9a-zA-Z_]{4,}$"
var charset = "0123456789abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Shortener struct {
}

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func (s Shortener) GetRandomString() string {
	b := make([]byte, 6) // make untuk membuat object [] untuk type object ke 2 jumlah datanya
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func (s Shortener) ValidateUrl(url string) bool {
	match, _ := regexp.MatchString(regex, url)

	return match
}
