package main

import (
	"fmt"

	"github.com/vroomy/common"
)

// Init is called when Vroomy initializes the plugin
func Init(env map[string]string) (err error) {
	return
}

// BasicAuth will provide a basic authentication layer
func BasicAuth(args ...string) (h common.Handler, err error) {
	var (
		username string
		password string
		realm    string
	)

	switch len(args) {
	case 2:
		username = args[0]
		password = args[1]

	case 3:
		username = args[0]
		password = args[1]
		realm = args[3]

	default:
		err = fmt.Errorf("invalid number of arguments, expected 2 or 3 and received %d", len(args))
		return
	}

	h = newHandler(username, password, realm)
	return
}

func newHandler(username, password, realm string) common.Handler {
	realm = fmt.Sprintf(`Basic realm="%s"`, realm)
	return func(ctx common.Context) {
		req := ctx.Request()
		u, p, ok := req.BasicAuth()
		switch {
		case !ok:
			w := ctx.Writer()
			w.Header().Set("WWW-Authenticate", realm)
			ctx.WriteString(401, "text/plain", "Please provide login credentials\n")
			return
		case u != username:
		case p != password:

		default:
			// Access is permitted!
			return
		}

		ctx.WriteString(401, "text/plain", "Insufficient access\n")
	}
}
