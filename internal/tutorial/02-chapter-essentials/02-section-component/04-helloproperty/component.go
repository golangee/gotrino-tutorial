// Package t04helloproperty demonstrates the usage of a property and the automatic re-rendering. The lifecycle
// is quite complex behind the scenes, because the property.Property invalidates its attached parent
// (a property.Invalidator which is introduced by embedding a view.View). IfCond performing a rendering (e.g.
// by using view.WithElement), the view.Component is observed and each future observed invalidation (e.g. caused
// by either calling view.View#Invalidate or by property.Property#Set), releases the lastly
// created view.Node (which created a dom.Element behind the scenes). Note, that each dom.Element#Release
// will also detach automatically all connected view.Component's.
//
// In short, an invalidation will always cause a full view.Component#Render cycle and replaces any attached
// dom.Element. Therefore, you cannot modify the dom.Element directly.
package t04helloproperty

import (
	"time"

	"github.com/golangee/dom/router"
	. "github.com/golangee/gotrino"
	. "github.com/golangee/gotrino-html"
	"github.com/golangee/property"
)

const Path = "/tutorial/04-helloproperty"

type ContentView struct {
	msg property.String
	View
}

func NewContentView() *ContentView {
	c := &ContentView{}
	c.msg.Set("hello property").Attach(c.Invalidate)
	return c
}

func (c *ContentView) Render() Node {
	return Span(Text(c.msg.Get()), AddClickListener(func() {
		c.msg.Set("time is " + time.Now().String())
	}))
}

func FromQuery(router.Query) Renderable {
	return NewContentView()
}
