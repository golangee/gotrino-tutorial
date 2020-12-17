package c02s05s02

import (
	"fmt"
	"github.com/golangee/dom"
	"github.com/golangee/dom/router"
	. "github.com/golangee/gotrino"
	gt "github.com/golangee/gotrino"
	. "github.com/golangee/gotrino-html"
	"github.com/golangee/gotrino-tailwind/button"
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
				NewTextField().BindText(&c.myModel.a),
				NewTextField().SetLabel("an empty text").BindText(&c.myModel.b),
				NewTextField().SetLabel("prefilled text").SetText("hello world").BindText(&c.myModel.c),
				NewTextField().SetLabel("a secret").SetType("password").BindText(&c.myModel.d),
				NewTextField().SetLabel("a number").SetType("number").BindText(&c.myModel.e),
				NewTextField().SetLabel("a date").SetType("date").BindText(&c.myModel.f),
			),
		),
		// print model
		button.NewButton(func() {
			c.previewModel.Set(fmt.Sprintf("%+v", c.myModel))
		}).SetContent(Text("show model")),


		Div(Text(c.previewModel.Get())),
	)
}

type TextField struct {
	text      property.String
	label     property.String
	floatTop  property.Bool
	focus     property.Bool
	inputType property.String
	name      property.String
	View
}

func NewTextField() *TextField {
	c := &TextField{}
	c.label.Attach(c.Invalidate)
	c.inputType.Set("text")
	c.inputType.Attach(c.Invalidate)
	c.name.Set(dom.GenerateID())
	c.name.Attach(c.Invalidate)
	c.focus.Observe(func(old, new bool) {
		c.floatTop.Set(new || c.text.Get() != "")
	})
	c.text.Observe(func(old, new string) {
		// date: non-empty text to fix non-empty placeholders in chrome and ff - Safari is different
		if c.inputType.Get() == "date" {
			c.floatTop.Set(true)
			return
		}

		c.floatTop.Set(new != "")
	})
	c.inputType.Observe(func(old, new string) {
		// date: non-empty text to fix non-empty placeholders in chrome and ff - Safari is different
		if new == "date" {
			c.floatTop.Set(true)
		}
	})
	return c
}

func (c *TextField) TextProperty() *property.String {
	return &c.text
}

func (c *TextField) NameProperty() *property.String {
	return &c.name
}

func (c *TextField) SetName(t string) *TextField {
	c.NameProperty().Set(t)
	return c
}

func (c *TextField) SetText(t string) *TextField {
	c.TextProperty().Set(t)
	return c
}

func (c *TextField) BindText(dst *string) *TextField {
	c.TextProperty().Bind(dst)
	return c
}

func (c *TextField) LabelProperty() *property.String {
	return &c.label
}

func (c *TextField) SetLabel(t string) *TextField {
	c.LabelProperty().Set(t)
	return c
}

func (c *TextField) TypeProperty() *property.String {
	return &c.inputType
}

func (c *TextField) SetType(t string) *TextField {
	c.TypeProperty().Set(t)
	return c
}

func (c *TextField) Render() Node {
	labelFocus := "scale-75 -translate-y-4 z-0 ml-3 px-1 py-0 bg-surface"

	return Div(Class("box-border rounded outline relative focus-within:border-primary"),
		IfCond(&c.focus,
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
				c.focus.Set(true)
			}),
			AddEventListener("blur", func() {
				c.focus.Set(false)
			}),
			Observe(&c.text.Property, func(e dom.Element) gt.Modifier {
				e.Set("value", c.text.Get())
				return nil
			}),
			InsideDom(func(e dom.Element) {
				e.AddEventListener("input", false, func() {
					c.text.Set(e.Get("value").(string))
				})
			}),
			Type(c.inputType.Get()), Name(c.name.Get()), Placeholder(" "),
		),
		Label(Class("absolute top-0 p-4 text-lg -z-1 duration-300 origin-0 transform"),
			Style("pointer-events", "none"),
			IfCond(&c.floatTop,
				Modifiers(
					AddClass(labelFocus),
					Style("left", "-0.75rem"),
				),
				Modifiers(
					RemoveClass(labelFocus),
					Style("left", "inherit"),
				)),
			For(c.name.Get()), Text(c.label.Get()),
		),
	)
}

func FromQuery(q router.Query) Renderable {
	return NewContentView()
}
