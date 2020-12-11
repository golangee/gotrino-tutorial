package t01helloworld

import (
	"github.com/golangee/dom/router"
	"github.com/golangee/gotrino"
	html "github.com/golangee/gotrino-html"
)

const Path = "/tutorial/01-helloworld"

func FromQuery(router.Query) gotrino.Renderable {
	return html.Span(html.Text("hello world"))
}
