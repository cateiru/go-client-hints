package ch

import (
	"fmt"
	"regexp"
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

	r := regexp.MustCompile(`"([^"]+)";v="([0-9.]+)"(,\s?)?`)

	brandsBuf := r.FindAllStringSubmatch(chUa, -1)
	brands := make(map[string]string)

	fmt.Println(brandsBuf)

	for _, brand := range brandsBuf {
		brands[brand[1]] = brand[2]
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
