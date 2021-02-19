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
	if len(args) != 2 {
		err = fmt.Errorf("invalid number of arguments, expected 2 and received %d", len(args))
		return
	}

	username := args[0]
	password := args[1]
	h = newHandler(username, password)
	return
}

func newHandler(username, password string) common.Handler {
	return func(ctx common.Context) {
		req := ctx.Request()
		u, p, ok := req.BasicAuth()
		switch {
		case !ok:
		case u != username:
		case p != password:

		default:
			// Access is permitted!
			return
		}

		// No access permitted, send no content
		ctx.WriteNoContent()
	}
}
