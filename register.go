// Package websockets exist just to register the websockets extension
package websockets

import (
	"github.com/edamame-load-test/xk6-websockets-extended/websockets"
	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/websockets", new(websockets.RootModule))
}
