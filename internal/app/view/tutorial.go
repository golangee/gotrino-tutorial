package view

import (
	. "github.com/golangee/gotrino"
	. "github.com/golangee/gotrino-html"
)

type TutorialModel struct {
	Steps []TutorialStep
}

type TutorialStep struct {
	Caption     string
	Description string
	Content     TutorialContent
}

type TutorialContent struct {
	Code  string
	Image string
	Video string
}

type Tutorial struct {
	View
	model TutorialModel
}

func NewTutorial(model TutorialModel) *Tutorial {
	return &Tutorial{
		model: model,
	}
}

func (c *Tutorial) Render() Node {
	return Div(Class("container mx-auto pt-20 pb-8 px-6 grid md:grid-cols-2 gap-6 grid-cols-1 max-w-5xl"),
		ForEach(len(c.model.Steps), func(i int) Renderable {
			step := c.model.Steps[i]
			return Yield(
				Div(
					Div(Class("border-l-8 p-6 rounded-lg hover:border-primary bg-gray-100 transition-colors"),
						P(Class("text-sm font-medium pb-2"), Text(step.Caption)),
						P(Text(step.Description)),
					),
				),

				Div(
					P(Text("bild/video/etc")),
					Text(step.Content.Code),
				),
			)

		}),
	)
}


