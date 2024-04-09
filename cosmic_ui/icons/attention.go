package icons

import (
	"fmt"
	goalpinejshandler "github.com/nodejayes/go-alpinejs-handler"
	"github.com/nodejayes/go-alpinejs-handler-poc/cosmic_ui/cosmic_ui_global"
)

type (
	AttentionIcon struct {
		goalpinejshandler.ViewTools
		Color string
	}
	AttentionIconArguments struct {
		Color string
	}
)

func NewAttentionIcon(args AttentionIconArguments) *AttentionIcon {
	if len(args.Color) < 1 {
		args.Color = "#000000"
	}
	return &AttentionIcon{
		Color: args.Color,
	}
}

func (ctx *AttentionIcon) Name() string {
	return cosmic_ui_global.CreateName("attention_icon")
}

func (ctx *AttentionIcon) Render() string {
	return fmt.Sprintf(`
<svg viewBox="0 0 24 24" role="img" xmlns="http://www.w3.org/2000/svg" aria-labelledby="errorIconTitle" stroke="%[1]s" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" fill="none" color="%[1]s">
	<g id="SVGRepo_bgCarrier" stroke-width="0"></g>
	<g id="SVGRepo_tracerCarrier" stroke-linecap="round" stroke-linejoin="round"></g>
	<g id="SVGRepo_iconCarrier">
		<title id="errorIconTitle">Error</title>
		<path d="M12 8L12 13"></path>
		<line x1="12" y1="16" x2="12" y2="16"></line>
		<circle cx="12" cy="12" r="10"></circle>
	</g>
</svg>
`, ctx.Color)
}
