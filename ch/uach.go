package ch

import (
	"strings"
)

type MobileType int

const (
	Mobile MobileType = iota
	NoMobile
	Unknown
)

// Parse the Sec-CH-UA header.
// Returns a map of "brand name - version".
//
// Example:
// 	chUa := r.Header.Get(ch.SecChUa)
// 	brands := ch.ParseSecChUa(chUa)
func ParseSecChUa(chUa string) map[string]string {
	if chUa == "" {
		return map[string]string{}
	}

	brandsBuf := strings.Split(chUa, ", ")
	brands := make(map[string]string)

	for _, brand := range brandsBuf {
		c := strings.Split(brand, ";v=")

		if len(c) == 2 {
			brands[c[0][1:len(c[0])-1]] = c[1][1 : len(c[1])-1]
		}
	}

	return brands
}

// Determine if mobile or not.
//
// Example:
// 	chUaMobile := r.Header.Get(ch.SecChUaMobile)
// 	mobileType := ch.ParseChUaMobile(chUaMobile)
func ParseChUaMobile(chUaMobile string) MobileType {

	if chUaMobile == "!0" {
		return NoMobile
	}
	if chUaMobile == "!1" {
		return Mobile
	}

	return Unknown
}
