package goclienthints

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/dunglas/httpsfv"
)

// Sec-CH-UA field
type Brand struct {
	Brand        string            `json:"brand,omitempty"`
	BrandVersion string            `json:"brand_version,omitempty"`
	Brands       map[string]string `json:"brands,omitempty"`
}

type ClientHints struct {
	Brand

	// Sec-Ch-Ua-Platform filed
	Platform Platform `json:"platform"`

	// Sec-CH-UA-Platform-Version filed
	PlatformVersion string `json:"platform_version,omitempty"`

	// Sec-Ch-Ua-Mobile filed
	IsMobile bool `json:"is_mobile"`

	// Sec-CH-UA-Arch filed
	Architecture string `json:"architecture,omitempty"`

	// Sec-CH-UA-Bitness filed
	Bitness int `json:"bitness"`

	// Sec-CH-UA-Model filed
	Model string `json:"model,omitempty"`

	// Sec-Ch-Ua-Full-Version filed
	FullVersion string `json:"full_version,omitempty"`
}

type Platform string

// Platforms
// You get `Sec-Ch-Ua-Platform` field
const (
	Android    Platform = "Android"
	ChromeOS   Platform = "Chrome OS"
	ChromiumOS Platform = "Chromium OS"
	IOS        Platform = "iOS"
	Linux      Platform = "Linux"
	MacOS      Platform = "macOS"
	Windows    Platform = "Windows"
	Unknown    Platform = "Unknown"
)

// User-Agent-Client-Hints Headers
const (
	HeaderSecChUa                = "Sec-Ch-Ua"
	HeaderSecChUaArch            = "Sec-Ch-Ua-Arch"
	HeaderSecChUaBitness         = "Sec-Ch-Ua-Bitness"
	HeaderSecChUaFullVersion     = "Sec-Ch-Ua-Full-Version"
	HeaderSecChUaFullVersionList = "Sec-Ch-Ua-Full-Version-List"
	HeaderSecChUaMobile          = "Sec-Ch-Ua-Mobile"
	HeaderSecChUaModel           = "Sec-Ch-Ua-Model"
	HeaderSecChUaPlatform        = "Sec-Ch-Ua-Platform"
	HeaderSecChUaPlatformVersion = "Sec-Ch-Ua-Platform-Version"
)

// Brand array
// A array of brands that Sec-Ch-Ua prefers to compare.
var PrimaryBrands = []string{
	"Google Chrome",
	"Chrome",
	"Microsoft Edge",
	"Edge",
	"Brave Browser",
	"Brave",
	"Yandex",
	"CocCoc",
}

// This array will be used to try to get the brand name
// if it is not in the PrimaryBrands array.
var SecondaryBrands = []string{
	"Chromium",
}

// Parse User-Agent Client Hints.
//
// Parse the header below.
//   - Sec-Ch-Ua
//   - Sec-Ch-Ua-Arch
//   - Sec-Ch-Ua-Bitness
//   - Sec-Ch-Ua-Full-Version
//   - Sec-Ch-Ua-Full-Version-List
//   - Sec-Ch-Ua-Mobile
//   - Sec-Ch-Ua-Model
//   - Sec-Ch-Ua-Platform
//   - Sec-Ch-Ua-Platform-Version
//
// If the header does not exist, its value will be the empty string, the number 0, or false.
func Parse(headers *http.Header) (*ClientHints, error) {
	// If you can get the full version, use it.
	chUa := headers.Get(HeaderSecChUaFullVersionList)
	if chUa == "" {
		chUa = headers.Get(HeaderSecChUa)
	}
	brand, err := ParseSecChUa(chUa)
	if err != nil {
		return nil, err
	}

	platform, err := ParsePlatform(headers.Get(HeaderSecChUaPlatform))
	if err != nil {
		return nil, err
	}

	platformVersion, err := ParseItem(headers.Get(HeaderSecChUaPlatformVersion))
	if err != nil {
		return nil, err
	}

	isMobile, err := ParseBool(headers.Get(HeaderSecChUaMobile))
	if err != nil {
		return nil, err
	}

	arch, err := ParseItem(headers.Get(HeaderSecChUaArch))
	if err != nil {
		return nil, err
	}

	bitnessStr, err := ParseItem(headers.Get(HeaderSecChUaBitness))
	if err != nil {
		return nil, err
	}
	var bitness int
	if bitnessStr == "" {
		bitness = 0
	} else {
		bitness, err = strconv.Atoi(bitnessStr)
		if err != nil {
			return nil, err
		}
	}

	model, err := ParseItem(headers.Get(HeaderSecChUaModel))
	if err != nil {
		return nil, err
	}

	fullVersion, err := ParseItem(headers.Get(HeaderSecChUaFullVersion))
	if err != nil {
		return nil, err
	}

	return &ClientHints{
		Brand:           *brand,
		Platform:        platform,
		PlatformVersion: platformVersion,
		IsMobile:        isMobile,
		Architecture:    arch,
		Bitness:         bitness,
		Model:           model,
		FullVersion:     fullVersion,
	}, nil
}

// Parse the `Sec-Ch-Ua` header
// Create an array of brand names and their versions,
// plus determine which one of the brands you are using.
func ParseSecChUa(h string) (*Brand, error) {
	if h == "" {
		return &Brand{}, nil
	}

	brands, err := httpsfv.UnmarshalList([]string{h})
	if err != nil {
		return nil, err
	}
	formattedBrandMaps := map[string]string{}
	for _, brand := range brands {
		item, ok := brand.(httpsfv.Item)
		if ok {
			brandName := item.Value.(string)
			brandVersion, ok := item.Params.Get("v")
			if !ok {
				return nil, errors.New("parse failed brand versions")
			}
			brandVersionStr := brandVersion.(string)

			formattedBrandMaps[brandName] = brandVersionStr
		}
	}

	b := &Brand{
		Brands: formattedBrandMaps,
	}

	for _, primaryBrand := range PrimaryBrands {
		for brandName, brandVersion := range formattedBrandMaps {
			if strings.EqualFold(brandName, primaryBrand) {
				b.Brand = brandName
				b.BrandVersion = brandVersion
				return b, nil
			}
		}
	}

	for _, secondaryBrand := range SecondaryBrands {
		for brandName, brandVersion := range formattedBrandMaps {
			if strings.EqualFold(brandName, secondaryBrand) {
				b.Brand = brandName
				b.BrandVersion = brandVersion
				return b, nil
			}
		}
	}

	return b, nil
}

// Prase the `Sec-CH-UA-Platform` header
func ParsePlatform(h string) (Platform, error) {
	platform, err := ParseItem(h)
	if err != nil {
		return "", err
	}

	switch platform {
	case "Android":
		return Android, nil
	case "Chrome OS":
		return ChromeOS, nil
	case "Chromium OS":
		return ChromiumOS, nil
	case "iOS":
		return IOS, nil
	case "Linux":
		return Linux, nil
	case "macOS":
		return MacOS, nil
	case "Windows":
		return Windows, nil
	default:
		return Unknown, nil
	}
}

// Determines if ClientHint is supported.
//
// It is determined by the presence or absence of the `Sec-Ch-Ua` header.
func IsSupportClientHints(headers *http.Header) bool {
	chUa := headers.Get(HeaderSecChUa)

	return chUa != ""
}

func ParseItem(h string) (string, error) {
	if h == "" {
		return "", nil
	}

	item, err := httpsfv.UnmarshalItem([]string{h})
	if err != nil {
		return "", err
	}

	itemStr, ok := item.Value.(string)

	if !ok {
		return "", errors.New("parse failed item")
	}

	return itemStr, nil
}

func ParseBool(h string) (bool, error) {
	if h == "" {
		return false, nil
	}

	item, err := httpsfv.UnmarshalItem([]string{h})
	if err != nil {
		return false, err
	}

	itemBool, ok := item.Value.(bool)

	if !ok {
		return false, errors.New("parse failed item")
	}

	return itemBool, nil
}
