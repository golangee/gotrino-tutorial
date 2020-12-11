package view

import (
	. "github.com/golangee/gotrino"
	. "github.com/golangee/gotrino-html"
)

type Step struct {
	View
	caption     string
	description Renderable
}

func NewStep(caption string, description Renderable) *Step {
	return &Step{
		caption:     caption,
		description: description,
	}
}

func (c *Step) Render() Node {
	return Div(Class("border-l-8 p-6 rounded-lg hover:border-primary bg-gray-100 transition-colors"),
		P(Class("text-sm font-medium pb-2"), Text(c.caption)),
		c.description,
	)
}
