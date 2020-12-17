package app

import (
	"github.com/golangee/dom/router"
	. "github.com/golangee/gotrino"
	. "github.com/golangee/gotrino-html"
	"github.com/golangee/gotrino-tailwind/progress"
	"github.com/golangee/gotrino-tailwind/tabs"
	"github.com/golangee/gotrino-tutorial/forms/highlightjs"
	"github.com/golangee/gotrino-tutorial/internal/index"
	"github.com/golangee/gotrino-tutorial/nestor"
	"github.com/golangee/wasm-net/fetch"
	"strconv"
	"strings"
)

func tutorialStepview(q router.Query) Renderable {
	pathIds := strings.Split(q.Path(), "/")
	if pathIds[0] == "" {
		pathIds = pathIds[1:]
	}

	var sec *nestor.Fragment
	for _, chapter := range index.Tutorial.Fragments {
		for _, section := range chapter.Fragments {
			if pathIds[0] == index.Tutorial.ID() && pathIds[1] == chapter.ID() && pathIds[2] == section.ID() {
				sec = section
				break
			}
		}
	}

	if sec == nil {
		return Span(Text("section not found: " + q.Path()))
	}

	return Div(Class("container mx-auto pt-20 pb-8 px-6 grid md:grid-cols-2 gap-6 grid-cols-1 max-w-7xl"),
		ForEach(len(sec.Fragments), func(i int) Renderable {
			step := sec.Fragments[i]
			return Yield(

				Div(
					Div(Class("border-l-8 p-6 rounded-lg hover:border-primary bg-gray-100 transition-colors"),
						P(Class("text-sm font-medium pb-2"), Text("Step "+strconv.Itoa(i+1))),
						P(InnerHTML(step.Body)),
					),
				),


				Div(
					tabs.NewTabs().With(func(c *tabs.Tabs) {
						for _, at := range step.Attachments {
							title := at.Title
							switch at.Type {
							case nestor.AtIFrame:
								title = "Preview"
							case nestor.AtSource:
								title = "Source"
							}
							c.AddPane(Text(title), attachmentPane(at))
						}
					}),


				),
			)

		}),
	)
}

func attachmentPane(at *nestor.Attachment) Renderable {
	switch at.Type {
	case nestor.AtIFrame:
		return Div(Class("shadow rounded-md border mt-4 m-auto"), Style("width", "278px"), Style("height", "602px"),
			IFrame(Style("width", "278px"), Style("height", "602px"), //seems to be buggy, because the iframe ignores partially the outer size
				Src("/#"+at.File),
			),
			Div(Class("text-center"), A(Href("/#"+at.File), Target("_blank"), Text("open in new tab"))),
		)
	case nestor.AtImage:
		return Img(Src(at.File))
	case nestor.AtSource:
		codeView := highlightjs.NewCode()
		codeView.LangProperty().Set("go")
		pg := progress.NewInfiniteCircle()
		pg.VisibleProperty().Set(true)

		fetch.Get(at.File, fetch.AsText(func(res string, err error) {
			if err != nil {
				codeView.CodeProperty().Set(err.Error())
				return
			}

			pg.VisibleProperty().Set(false)
			codeView.CodeProperty().Set(res)
		}))

		return Div(Class("max-w-prose overflow-x-auto bg-gray-100"),
			Div(Class("text-center"),
				Span(Class("m-2"),
					Style("display", "inline-block"),
					IfCond(pg.VisibleProperty(), Style("display", "inline-block"), Style("display", "none")),
					pg,
				),
			),

			codeView,
		)

	}
	return Div(Text("type not implemented: " + string(at.Type)))
}
