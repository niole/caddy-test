package mymodule

import (
	"log"
	"net/http"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
)

func init() {
	caddy.RegisterModule(MyHandler{})
}

type MyHandler struct{}

func (MyHandler) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.mymodule",
		New: func() caddy.Module { return new(MyHandler) },
	}
}

func (h MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
	// Your custom logic hereA

	log.Println("NIOLE")

	return next.ServeHTTP(w, r)
}
