// Package global automatically sets http.DefaultTransport to loghttp.DefaultTransport when loaded.
package global

import (
	"net/http"

	"github.com/walf443/go-loghttp"
)

func init() {
	http.DefaultTransport = loghttp.DefaultTransport
}
