package c02s05s02

import (
	"fmt"
	"github.com/golangee/dom/router"
	. "github.com/golangee/gotrino"
	. "github.com/golangee/gotrino-html"
	"github.com/golangee/gotrino-tailwind/button"
	"github.com/golangee/gotrino-tailwind/input"
	"github.com/golangee/property"
)

const Path = "/c02s05s02"

type ContentView struct {
	myModel struct {
		a, b, c, d, e, f string
	}
	previewModel property.String
	View
}

func NewContentView() *ContentView {
	c := &ContentView{}
	c.previewModel.Attach(c.Invalidate)
	return c
}

func (c *ContentView) Render() Node {
	text := property.NewString("test")

	floatTop := property.NewBool(text.Get() != "")
	focus := property.NewBool(false)
	focus.Observe(func(old, new bool) {
		floatTop.Set(new || text.Get() != "")
	})

	return Div(Class("p-10"),
		Form(
			Div(Class("grid grid-cols-1 gap-4"),
				input.NewTextField().BindText(&c.myModel.a),
				input.NewTextField().SetLabel("an empty text").BindText(&c.myModel.b),
				input.NewTextField().SetLabel("prefilled text").SetText("hello world").BindText(&c.myModel.c),
				input.NewTextField().SetLabel("a secret").SetType("password").BindText(&c.myModel.d),
				input.NewTextField().SetLabel("a number").SetType("number").BindText(&c.myModel.e),
				input.NewTextField().SetLabel("a date").SetType("date").BindText(&c.myModel.f),

				// show, that the model is automatically updated
				button.NewButton(func() {
					c.previewModel.Set(fmt.Sprintf("%+v", c.myModel))
				}).SetContent(Text("show model")),
			),
		),


		Div(Text(c.previewModel.Get())),
	)
}

func FromQuery(q router.Query) Renderable {
	return NewContentView()
}
