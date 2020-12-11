package app

import (
	"github.com/golangee/dom/router"
	. "github.com/golangee/gotrino"
	. "github.com/golangee/gotrino-html"
	"github.com/golangee/gotrino-tailwind/appbar"
	"github.com/golangee/gotrino-tailwind/button"
	"github.com/golangee/gotrino-tailwind/icon"
	"github.com/golangee/gotrino-tailwind/layout"
)

func (a *Application) page(q router.Query, content Renderable) Renderable {
	routes := a.router.Routes()

	var bar *appbar.AppBar
	page := Div(
		appbar.NewAppBar().Self(&bar).SetIcon(Img(Class("h-8 w-auto px-6"), Src("https://www.worldiety.de/_nuxt/img/wdy_logo_white.8bf54.svg"), Alt("Logo"))).
			SetTitle(Span(Text("golangee/forms: "+q.Path()))).
			SetDrawerHeader(Img(
				Class("h-auto w-32 mx-auto bg-primary p-2"),
				Alt("Logo"),
				Src("https://www.worldiety.de/_nuxt/img/wdy_logo_white.8bf54.svg"),
			)).

			SetDrawerMain( // a side menu button
				layout.NewVStack(
					ForEach(len(routes), func(i int) Renderable {
						btn := button.NewIconTextButton(icon.FolderOpen, routes[i].Path, func() {
							bar.Close()
							router.Navigate("#" + routes[i].Path)
						})

						if q.Path() == routes[i].Path {
							return Span(Class("bg-primary bg-opacity-10 w-full"), btn) //bg-primary bg-opacity-10 bg-indigo-600 bg-opacity-10
						}

						return With(btn, AddClass("w-full"))
					}),

				),
			).

			SetDrawerFooter(Div(
				Class("fixed bottom-0 w-full"),
				With(
					button.NewIconTextButton(icon.OpenInBrowser, "what is worldiety?", func() {
						router.Navigate("https://worldiety.de")
					}),


					AddClass("w-full"),
				),
			)),

		content,
	)

	return page
}
