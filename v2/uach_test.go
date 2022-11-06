package goclienthints_test

import (
	"fmt"
	"net/http"
	"testing"

	goclienthints "github.com/cateiru/go-client-hints"
	"github.com/stretchr/testify/require"
)

type TestParseCase struct {
	Header    http.Header
	IsSuccess bool
	Response  goclienthints.ClientHints
}

type TestChUaParam struct {
	Value        string
	IsSuccess    bool
	BrandName    string
	BrandVersion string
}

type TestChUaPlatform struct {
	Value     string
	IsSuccess bool
	Platform  goclienthints.Platform
}

type TestItem struct {
	Value        string
	IsSuccess    bool
	ReturnsValue string
}

type TestBool struct {
	Value         string
	IsSuccess     bool
	ReturnsStatus bool
}

func TestParse(t *testing.T) {
	cases := []TestParseCase{
		{
			Header: http.Header{
				"Sec-Ch-Ua":                  {`"Chromium";v="84", "Google Chrome";v="84"`},
				"Sec-Ch-Ua-Mobile":           {"?0"},
				"Sec-Ch-Ua-Platform":         {`"macOS"`},
				"Sec-Ch-Ua-Platform-Version": {`"11"`},
				"Sec-Ch-Ua-Arch":             {`"ARM"`},
				"Sec-Ch-Ua-Bitness":          {`"64"`},
				"Sec-Ch-Ua-Model":            {`"Pixel 6"`},
				"Sec-Ch-Ua-Full-Version":     {`"123456"`},
			},
			IsSuccess: true,
			Response: goclienthints.ClientHints{
				Brand: goclienthints.Brand{
					Brand:        "Google Chrome",
					BrandVersion: "84",
					Brands: map[string]string{
						"Google Chrome": "84",
						"Chromium":      "84",
					}},
				Platform:        goclienthints.MacOS,
				PlatformVersion: "11",
				IsMobile:        false,
				Architecture:    "ARM",
				Bitness:         64,
				Model:           "Pixel 6",
				FullVersion:     "123456",
			},
		},
		{
			Header: http.Header{
				"Sec-Ch-Ua":                   {`"Chromium";v="84", "Google Chrome";v="84"`},
				"Sec-Ch-Ua-Full-Version-List": {`"Microsoft Edge"; v="92.0.902.73", "Chromium"; v="92.0.4515.131", "?Not:Your Browser"; v="3.1.2.0"`},
				"Sec-Ch-Ua-Mobile":            {"?0"},
				"Sec-Ch-Ua-Platform":          {`"macOS"`},
				"Sec-Ch-Ua-Platform-Version":  {`"11"`},
				"Sec-Ch-Ua-Arch":              {`"ARM"`},
				"Sec-Ch-Ua-Bitness":           {`"64"`},
				"Sec-Ch-Ua-Model":             {`"Pixel 6"`},
				"Sec-Ch-Ua-Full-Version":      {`"123456"`},
			},
			IsSuccess: true,
			Response: goclienthints.ClientHints{
				Brand: goclienthints.Brand{
					Brand:        "Microsoft Edge",
					BrandVersion: "92.0.902.73",
					Brands: map[string]string{
						"Microsoft Edge":    "92.0.902.73",
						"Chromium":          "92.0.4515.131",
						"?Not:Your Browser": "3.1.2.0",
					}},
				Platform:        goclienthints.MacOS,
				PlatformVersion: "11",
				IsMobile:        false,
				Architecture:    "ARM",
				Bitness:         64,
				Model:           "Pixel 6",
				FullVersion:     "123456",
			},
		},
		{
			Header:    http.Header{},
			IsSuccess: true,
			Response: goclienthints.ClientHints{
				Brand: goclienthints.Brand{
					Brand:        "",
					BrandVersion: "",
					Brands:       nil,
				},
				Platform:        goclienthints.Unknown,
				PlatformVersion: "",
				IsMobile:        false,
				Architecture:    "",
				Bitness:         0,
				Model:           "",
				FullVersion:     "",
			},
		},
		{
			Header: http.Header{
				"Sec-Ch-Ua":        {`"Chromium";v="84", "Google Chrome";v="84"`},
				"Sec-Ch-Ua-Mobile": {"?0"},

				// Invalid type
				"Sec-Ch-Ua-Platform": {`macOS`},
			},
			IsSuccess: false,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			clientHints, err := goclienthints.Parse(&c.Header)

			if c.IsSuccess {
				require.NoError(t, err)
				require.Equal(t, *clientHints, c.Response)
			} else {
				require.Error(t, err)
			}
		})
	}
}

func TestParseSecChUa(t *testing.T) {
	cases := []TestChUaParam{
		{
			Value:        `"Chrome"; v="74", ";Not)Your=Browser"; v="13"`,
			IsSuccess:    true,
			BrandName:    "Chrome",
			BrandVersion: "74",
		},
		{
			Value:        `"Chrome"; v="74.0.3729.0", "Chromium"; v="74.0.3729.0", "?Not:Your Browser"; v="13.0.1.0"`,
			IsSuccess:    true,
			BrandName:    "Chrome",
			BrandVersion: "74.0.3729.0",
		},
		{
			Value:        `"Chromium";v="84", "Google Chrome";v="84"`,
			IsSuccess:    true,
			BrandName:    "Google Chrome",
			BrandVersion: "84",
		},
		{
			Value:        `"Microsoft Edge"; v="92.0.902.73", "Chromium"; v="92.0.4515.131", "?Not:Your Browser"; v="3.1.2.0"`,
			IsSuccess:    true,
			BrandName:    "Microsoft Edge",
			BrandVersion: "92.0.902.73",
		},
		{
			Value:     `"Microsoft Edge"; v="92.0.`,
			IsSuccess: false,
		},
		{
			Value:        `" Not A;Brand";v="99", "Chromium";v="96"`,
			IsSuccess:    true,
			BrandName:    "Chromium",
			BrandVersion: "96",
		},
		{
			Value:        `" Not A;Brand";v="99"`,
			IsSuccess:    true,
			BrandName:    "",
			BrandVersion: "",
		},
	}

	for _, c := range cases {
		t.Run(c.Value, func(t *testing.T) {
			b, err := goclienthints.ParseSecChUa(c.Value)

			if c.IsSuccess {
				require.NoError(t, err)
				require.Equal(t, b.Brand, c.BrandName)
				require.Equal(t, b.BrandVersion, c.BrandVersion)
			} else {
				require.Error(t, err)
			}
		})
	}
}

func TestParsePlatform(t *testing.T) {
	cases := []TestChUaPlatform{
		{
			Value:     `"Android"`,
			IsSuccess: true,
			Platform:  goclienthints.Android,
		},
		{
			Value:     `"Chrome OS"`,
			IsSuccess: true,
			Platform:  goclienthints.ChromeOS,
		},
		{
			Value:     `"Chromium OS"`,
			IsSuccess: true,
			Platform:  goclienthints.ChromiumOS,
		},
		{
			Value:     `"iOS"`,
			IsSuccess: true,
			Platform:  goclienthints.IOS,
		},
		{
			Value:     `"Linux"`,
			IsSuccess: true,
			Platform:  goclienthints.Linux,
		},
		{
			Value:     `"macOS"`,
			IsSuccess: true,
			Platform:  goclienthints.MacOS,
		},
		{
			Value:     `"Windows"`,
			IsSuccess: true,
			Platform:  goclienthints.Windows,
		},
		{
			Value:     `"Unknown"`,
			IsSuccess: true,
			Platform:  goclienthints.Unknown,
		},
		{
			Value:     `"Cat"`,
			IsSuccess: true,
			Platform:  goclienthints.Unknown,
		},
		{
			Value:     `"Unk`,
			IsSuccess: false,
		},
		{
			Value:     "",
			IsSuccess: true,
			Platform:  goclienthints.Unknown,
		},
		{
			Value:     "?0",
			IsSuccess: false,
		},
	}

	for _, c := range cases {
		t.Run(c.Value, func(t *testing.T) {
			p, err := goclienthints.ParsePlatform(c.Value)

			if c.IsSuccess {
				require.NoError(t, err)
				require.Equal(t, p, c.Platform)
			} else {
				require.Error(t, err)
			}
		})
	}
}

func TestIsClientHints(t *testing.T) {
	supportClientHintHeader := http.Header{
		"Sec-Ch-Ua": {`"Chrome"; v="74", ";Not)Your=Browser"; v="13"`},
	}
	noSupportClientHintHeader := http.Header{}

	isSupport := goclienthints.IsSupportClientHints(&supportClientHintHeader)
	require.True(t, isSupport)

	isSupport = goclienthints.IsSupportClientHints(&noSupportClientHintHeader)
	require.False(t, isSupport)
}

func TestParseItem(t *testing.T) {
	cases := []TestItem{
		{
			Value:        `"aaa"`,
			IsSuccess:    true,
			ReturnsValue: "aaa",
		},
		{
			Value:        `""`,
			IsSuccess:    true,
			ReturnsValue: "",
		},
		{
			Value:        "",
			IsSuccess:    true,
			ReturnsValue: "",
		},
		{
			Value:     "aaaa",
			IsSuccess: false,
		},
	}

	for _, c := range cases {
		t.Run(c.Value, func(t *testing.T) {
			item, err := goclienthints.ParseItem(c.Value)

			if c.IsSuccess {
				require.NoError(t, err)
				require.Equal(t, item, c.ReturnsValue)
			} else {
				require.Error(t, err)
			}
		})
	}
}

func TestParseBool(t *testing.T) {
	cases := []TestBool{
		{
			Value:         "?0",
			IsSuccess:     true,
			ReturnsStatus: false,
		},
		{
			Value:         "?1",
			IsSuccess:     true,
			ReturnsStatus: true,
		},
		{
			Value:     "???1",
			IsSuccess: false,
		},
		{
			Value:         "",
			IsSuccess:     true,
			ReturnsStatus: false,
		},
	}

	for _, c := range cases {
		t.Run(c.Value, func(t *testing.T) {
			s, err := goclienthints.ParseBool(c.Value)

			if c.IsSuccess {
				require.NoError(t, err)
				require.Equal(t, s, c.ReturnsStatus)
			} else {
				require.Error(t, err)
			}
		})
	}
}
