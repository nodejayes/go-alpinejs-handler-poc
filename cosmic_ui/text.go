package cosmic_ui

import (
	goalpinejshandler "github.com/nodejayes/go-alpinejs-handler"
	"github.com/nodejayes/go-alpinejs-handler-poc/cosmic_ui/cosmic_ui_global"
)

type (
	TextArguments struct {
		Content   string
		ClassName string
	}
	Text struct {
		goalpinejshandler.ViewTools
		Content   string
		ClassName string
	}
)

func NewText(args TextArguments) *Text {
	return &Text{
		Content:   args.Content,
		ClassName: args.ClassName,
	}
}

func (ctx *Text) Name() string {
	return cosmic_ui_global.CreateName("text")
}

func (ctx *Text) Render() string {
	return `<span class="{{ .ClassName }}" x-text="'{{ .Content }}'"></span>`
}
