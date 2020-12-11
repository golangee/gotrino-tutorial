package c02s03s02

import (
	"github.com/golangee/dom"
	"github.com/golangee/dom/router"
	. "github.com/golangee/gotrino"
	. "github.com/golangee/gotrino-html"
	"github.com/golangee/gotrino-tailwind/button"
	"github.com/golangee/gotrino-tailwind/menu"
)

const Path = "/c02s03s02"

type ContentView struct {
	View
}


func NewContentView() *ContentView {
	return &ContentView{}
}

func (c *ContentView) Render() Node {
	btnId := dom.GenerateID()

	return Div(Class("w-screen h-screen grid"),
		With(button.NewTextButton("show menu", func() {
			menu.ShowPopup(btnId, menu.NewMenu(Div(
				menu.NewMenuItem(Text("hey2")),
				menu.NewMenuItem(Text("ho2")),
			)))
		}), AddClass("m-auto"), ID(btnId)),
	)
}

func FromQuery(q router.Query) Renderable {
	return NewContentView()
}
