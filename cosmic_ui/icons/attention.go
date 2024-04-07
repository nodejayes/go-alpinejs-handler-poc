package icons

import (
	goalpinejshandler "github.com/nodejayes/go-alpinejs-handler"
	"github.com/nodejayes/go-alpinejs-handler-poc/cosmic_ui/cosmic_ui_global"
)

type AttentionIcon struct {
	goalpinejshandler.ViewTools
}

func NewAttentionIcon() *AttentionIcon {
	return &AttentionIcon{}
}

func (ctx *AttentionIcon) Name() string {
	return cosmic_ui_global.CreateName("attention_icon")
}

func (ctx *AttentionIcon) Render() string {
	return `
<svg viewBox="0 0 24 24" role="img" xmlns="http://www.w3.org/2000/svg" aria-labelledby="errorIconTitle" stroke="#000000" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" fill="none" color="#000000">
	<g id="SVGRepo_bgCarrier" stroke-width="0"></g>
	<g id="SVGRepo_tracerCarrier" stroke-linecap="round" stroke-linejoin="round"></g>
	<g id="SVGRepo_iconCarrier">
		<title id="errorIconTitle">Error</title>
		<path d="M12 8L12 13"></path>
		<line x1="12" y1="16" x2="12" y2="16"></line>
		<circle cx="12" cy="12" r="10"></circle>
	</g>
</svg>
`
}
