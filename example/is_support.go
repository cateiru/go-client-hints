package example

import (
	"net/http"

	clienthint "github.com/cateiru/go-client-hints/v2"
)

func Handler2(w http.ResponseWriter, r *http.Request) {
	isSupport := clienthint.IsSupportClientHints(&r.Header)

	if isSupport {
		// ...do something
	}
}
