package c02s05s01

import (
	"github.com/golangee/dom"
	"github.com/golangee/dom/router"
	. "github.com/golangee/gotrino"
	gt "github.com/golangee/gotrino"
	. "github.com/golangee/gotrino-html"
	"github.com/golangee/property"
)

const Path = "/c02s05s01"

type ContentView struct {
	View
}

func NewContentView() *ContentView {
	return &ContentView{}
}

func (c *ContentView) Render() Node {
	text := property.NewString("test")

	floatTop := property.NewBool(text.Get() != "")
	focus := property.NewBool(false)
	focus.Observe(func(old, new bool) {
		floatTop.Set(new || text.Get() != "")
	})

	labelFocus := "scale-75 -translate-y-4 z-0 ml-3 px-1 py-0 bg-surface"
	return Div(Class("p-10"),
		Form(
			// actual text field
			Div(Class("box-border rounded outline relative focus-within:border-primary"),
				IfCond(focus,
					Modifiers(
						RemoveClass("hover:border-black border"),
						AddClass("border-2"),
					),
					Modifiers(
						AddClass("hover:border-black border"),
						RemoveClass("border-2"),
					)),
				Input(Class("block p-4 w-full text-lg appearance-none focus:outline-none bg-transparent"),
					AddEventListener("focus", func() {
						focus.Set(true)
					}),
					AddEventListener("blur", func() {
						focus.Set(false)
					}),
					Observe(&text.Property, func(e dom.Element) gt.Modifier {
						e.Set("value", text.Get())
						return nil
					}),
					InsideDom(func(e dom.Element) {
						e.AddEventListener("input", false, func() {
							text.Set(e.Get("value").(string))
						})
					}),
					Type("text"), Name("myfield"), Placeholder(" "),
				),
				Label(Class("absolute top-0 p-4 text-lg -z-1 duration-300 origin-0 transform"),
					Style("pointer-events", "none"),
					IfCond(floatTop,
						Modifiers(
							AddClass(labelFocus),
							Style("left", "-0.75rem"),
						),
						Modifiers(
							RemoveClass(labelFocus),
							Style("left", "inherit"),
						)),
					For("myfield"), Text("Username"),
				),
			),
		),

	)

}

func FromQuery(q router.Query) Renderable {
	return NewContentView()
}
