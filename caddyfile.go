package mymodule

import (
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/caddy/v2/caddyconfig/httpcaddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
)

func init() {
	httpcaddyfile.RegisterHandlerDirective("mymodule", parseCaddyfile)
}

func parseCaddyfile(h httpcaddyfile.Helper) (caddyhttp.MiddlewareHandler, error) {
	handler := new(MyHandler)
	err := handler.UnmarshalCaddyfile(h.Dispenser)
	return handler, err
}

func (h *MyHandler) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		// No same-line options are supported
		if d.NextArg() {
			return d.ArgErr()
		}

		for d.NextBlock(0) {
			switch d.Val() {
			case "address":
				if !d.NextArg() {
					return d.ArgErr()
				}
				//				h.Address = d.Val()
			default:
				return d.ArgErr()
			}
		}
	}
	return nil
}
