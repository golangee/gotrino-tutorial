package app

import (
	"github.com/golangee/dom/router"
	"github.com/golangee/gotrino-tutorial/internal/index"
	c01s01s01 "github.com/golangee/gotrino-tutorial/internal/tutorial/01-chapter-setup/01-section-setup/01-setup"
	c02s03s01 "github.com/golangee/gotrino-tutorial/internal/tutorial/02-chapter-essentials/03-section-modals/01-dialog"
	c02s03s02 "github.com/golangee/gotrino-tutorial/internal/tutorial/02-chapter-essentials/03-section-modals/02-menu"
	. "github.com/golangee/gotrino"
	. "github.com/golangee/gotrino-html"
	c02s04s01 "github.com/golangee/gotrino-tutorial/internal/tutorial/02-chapter-essentials/04-section-table/01-table"
	"github.com/golangee/log"
	"github.com/golangee/log/ecs"
)

type Application struct {
	router *router.Router
	log    log.Logger
}

func NewApplication() *Application {
	a := &Application{
		router: router.NewRouter(),
		log:    log.NewLogger(ecs.Log("application")),
	}

	a.router.AddRoute("/", a.apply(tutorialOverview))
	a.router.AddRoute(c01s01s01.Path, a.applyNone(c01s01s01.Show))
	a.router.AddRoute(c02s03s01.Path, a.applyNone(c02s03s01.FromQuery))
	a.router.AddRoute(c02s03s02.Path, a.applyNone(c02s03s02.FromQuery))
	a.router.AddRoute(c02s04s01.Path, a.applyNone(c02s04s01.FromQuery))

	for _, chapter := range index.Tutorial.Fragments {
		for _, section := range chapter.Fragments {
			path := "/" + index.Tutorial.ID() + "/" + chapter.ID() + "/" + section.ID()
			a.router.AddRoute(path, a.apply(tutorialStepview))
		}
	}

	a.router.
		SetUnhandledRouteAction(a.apply(func(query router.Query) Renderable {
			return Div(Class("pt-20 text-center"),
				Span(Class("border shadow-lg p-6 m-auto bg-red-300"), Text("unmatched route to "+query.Path())),
			)
		}))

	return a
}

func (a *Application) applyNone(f func(query router.Query) Renderable) func(query router.Query) {
	return func(query router.Query) {
		RenderBody(f(query))
	}
}

func (a *Application) apply(f func(query router.Query) Renderable) func(query router.Query) {
	return func(query router.Query) {
		RenderBody(a.page(query, f(query)))
	}
}

func (a *Application) index(q router.Query) Renderable {
	return P(Text("Welcome to the forms tutorial"))
}

func (a *Application) Run() {
	a.router.Start()
	select {}
}
