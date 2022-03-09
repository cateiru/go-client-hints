package ch_test

import (
	"testing"

	"github.com/cateiru/go-client-hints/ch"
	"github.com/stretchr/testify/require"
)

func TestParseSecChUa(t *testing.T) {
	samples := map[string]map[string]string{
		`"(Not(A:Brand";v="8", "Chromium";v="98"`: {
			"(Not(A:Brand": "8",
			"Chromium":     "98",
		},

		// Google Chrome
		`" Not A;Brand";v="99", "Chromium";v="96", "Google Chrome";v="96"`: {
			" Not A;Brand":  "99",
			"Chromium":      "96",
			"Google Chrome": "96",
		},
		`" Not A;Brand";v="99", "Chromium";v="99", "Google Chrome";v="99"`: {
			" Not A;Brand":  "99",
			"Chromium":      "99",
			"Google Chrome": "99",
		},
		`"(Not(A:Brand";v="8", "Chromium";v="100", "Google Chrome";v="100"`: {
			"(Not(A:Brand":  "8",
			"Chromium":      "100",
			"Google Chrome": "100",
		},

		// Edge
		`" Not A;Brand";v="99", "Chromium";v="96", "Microsoft Edge";v="96"`: {
			" Not A;Brand":   "99",
			"Chromium":       "96",
			"Microsoft Edge": "96",
		},
		`" Not A;Brand";v="99", "Chromium";v="99", "Google Chrome";v="99.0.1150.36"`: {
			" Not A;Brand":  "99",
			"Chromium":      "99",
			"Google Chrome": "99.0.1150.36",
		},

		// Opera
		`"Opera";v="81", " Not;A Brand";v="99", "Chromium";v="95"`: {
			"Opera":        "81",
			" Not;A Brand": "99",
			"Chromium":     "95",
		},
	}

	for headText, parsed := range samples {
		result := ch.ParseSecChUa(headText)

		require.Equal(t, result, parsed)
	}
}

func TestMobile(t *testing.T) {
	samples := map[string]ch.MobileType{
		"!0":      ch.NoMobile,
		"!1":      ch.Mobile,
		"":        ch.Unknown,
		"nyancat": ch.Unknown,
	}

	for text, parsed := range samples {
		result := ch.ParseChUaMobile(text)

		require.Equal(t, result, parsed)
	}
}
