package main

import (
	"fmt"
	"net/http"

	"github.com/cateiru/go-client-hints/ch"
	"github.com/cateiru/go-client-hints/ch/headers"
)

func AnnounceHandler(w http.ResponseWriter, r *http.Request) {
	announceHeaders := []string{
		headers.SecChUa,
		headers.SecChUaArch,
		headers.SecChUaPlatform,
		headers.SecChUaBitness,
		headers.SecChUaMobile,
	}

	ch.Announce(w, announceHeaders)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	chUa := r.Header.Get(headers.SecChUa)
	userData := ch.ParseSecChUa(chUa)

	for brandName, version := range userData {
		fmt.Printf("BrandName: %s, version: %s", brandName, version)
	}

	isMobile := ch.ParseChUaMobile(r.Header.Get(headers.SecChUaMobile))

	switch isMobile {
	case ch.Mobile:
		fmt.Println("Mobile!")
	case ch.NoMobile:
		fmt.Println("No Mobile!")
	case ch.Unknown:
		fmt.Println("Unknown!")
	}
}
