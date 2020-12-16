// Package t02hellohtml shows a more complex example which just replicates a Tailwind CSS frontpage demo.
// Please notice, how expressive the notation is by just dot-import
// the forms view package and start coding. Using the dot-import allows to write Go-code nearly as efficient
// as it would be in html/xml.
package t02hellohtml

import (
	"github.com/golangee/dom/router"
	. "github.com/golangee/gotrino"
	. "github.com/golangee/gotrino-html"
)

const Path = "/tutorial/02-hellohtml"

func FromQuery(router.Query) Renderable {
	return Div(Class("rounded overflow-hidden shadow-lg dark:bg-gray-800"),
		Figure(Class("md:flex bg-gray-100 rounded-xl p-8 md:p-0"),
			Img(
				Class("object-contain w-32 h-32 md:w-48 md:h-auto md:rounded-none rounded-full mx-auto"),
				Src("https://content-prod.worldiety.de/v2/media/images/wdy_200414_torben.d48be593.png"),
				Width("384"),
				Height("512"),
			),
			Div(
				Class("pt-6 md:p-8 text-center md:text-left space-y-4"),
				Blockquote(
					P(
						Class("text-lg font-semibold"),
						Text("“Forms is efficient to write and uses a modern WASM client side rendering approach. Using Tailwind CSS allows easy customization of html and building components.”"),
					),
				),
				Figcaption(Class("font-medium"),
					Div(Class("text-yellow-400"),
						Text("Torben Schinke"),
					),
					Div(Class("text-gray-500"),
						Text("Author of golangee/forms"),
					),
				),
			),
		),
	)
}
