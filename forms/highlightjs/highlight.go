package highlightjs

import (
	"github.com/golangee/dom"
	. "github.com/golangee/gotrino"
	. "github.com/golangee/gotrino-html"
	"github.com/golangee/property"
)

type CodeView struct {
	code property.String
	lang property.String
	View
}

func NewCode() *CodeView {
	c := &CodeView{}
	c.code.Attach(c.Invalidate)
	return c
}

func (c *CodeView) CodeProperty() *property.String {
	return &c.code
}

func (c *CodeView) LangProperty() *property.String {
	return &c.lang
}

func (c *CodeView) Render() Node {
	return Pre(Class(c.lang.Get()),
		Code(
			Text(c.code.Get()),
		),
		// order is important here, because otherwise round the Code block has not been created yet
		InsideDom(func(e dom.Element) {

			dom.GetGlobal().Get("hljs").Call("highlightBlock", e)
			dom.GetGlobal().Get("hljs").Call("lineNumbersBlock", e)
		}),
	)
}
