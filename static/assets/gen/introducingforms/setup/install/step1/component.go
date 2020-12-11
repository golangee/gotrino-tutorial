package c01s01s01

import (
	"github.com/golangee/dom/router"
	"github.com/golangee/gotrino"
	. "github.com/golangee/gotrino-html"
)

const Path = "/c01s01s01"

func Show(q router.Query) gotrino.Renderable {
	return Div(Text("Hello world4"))
}
