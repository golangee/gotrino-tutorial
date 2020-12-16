package c02s04s01

import (
	"github.com/golangee/dom/router"
	. "github.com/golangee/gotrino"
	. "github.com/golangee/gotrino-html"
)

const Path = "/c02s04s01"

type ContentView struct {
	View
}

func NewContentView() *ContentView {
	return &ContentView{}
}

func (c *ContentView) Render() Node {
	return Div(Class("p-4 bg-gray-100 min-w-full"),
		// frame around table
		Div(Class("border border-gray-200 rounded bg-white overflow-auto"),
			// actual table
			Table(Class(""),
				Thead(Class("font-medium"),
					Tr(
						Th(Class("w-1/3 text-left py-3 px-4 font-semibold text-sm"),
							Text("first name"),
						),
						Th(Class("w-1/3 text-left py-3 px-4 font-semibold text-sm"),
							Text("last name"),
						),
						Th(Class("text-left py-3 px-4 font-semibold text-sm"),
							Text("id"),
						),

						Th(Class("text-left py-3 px-4 font-semibold text-sm"),
							Text("mail"),
						),
					),
				),

				Tbody(Class("text-gray-700"),
					Tr(Class("hover:bg-gray-50 border-t border-gray-200"),
						Td(Class("w-1/3 text-left py-3 px-4"),
							Text("Goph"),
						),
						Td(Class("w-1/3 text-left py-3 px-4"),
							Text("O'Neil"),
						),
						Td(Class("text-left py-3 px-4"),
							Text("1234"),
						),
						Td(Class("text-left py-3 px-4"),
							Text("gopher.oneil@sgc.com"),
						),
					),

					Tr(Class("hover:bg-gray-50 border-t border-gray-200"),
						Td(Class("w-1/3 text-left py-3 px-4"),
							Text("Gopha"),
						),
						Td(Class("w-1/3 text-left py-3 px-4"),
							Text("Carter"),
						),
						Td(Class("text-left py-3 px-4"),
							Text("2345"),
						),
						Td(Class("text-left py-3 px-4"),
							Text("gopha.carter@sgc.com"),
						),
					),


					Tr(Class("hover:bg-gray-50 border-t border-gray-200"),
						Td(Class("w-1/3 text-left py-3 px-4"),
							Text("Gophel"),
						),
						Td(Class("w-1/3 text-left py-3 px-4"),
							Text("Jackson"),
						),
						Td(Class("text-left py-3 px-4"),
							Text("3456"),
						),
						Td(Class("text-left py-3 px-4"),
							Text("gophel.jackson@sgc.com"),
						),
					),
				),
			),
		),
	)
}

func FromQuery(q router.Query) Renderable {
	return NewContentView()
}
