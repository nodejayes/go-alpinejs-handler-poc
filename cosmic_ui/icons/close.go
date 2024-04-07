package icons

import (
	goalpinejshandler "github.com/nodejayes/go-alpinejs-handler"
	"github.com/nodejayes/go-alpinejs-handler-poc/cosmic_ui/cosmic_ui_global"
)

type CloseIcon struct {
	goalpinejshandler.ViewTools
}

func NewCloseIcon() *CloseIcon {
	return &CloseIcon{}
}

func (ctx *CloseIcon) Name() string {
	return cosmic_ui_global.CreateName("close_icon")
}

func (ctx *CloseIcon) Render() string {
	return `
<svg viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
	<g id="SVGRepo_bgCarrier" stroke-width="0"></g>
	<g id="SVGRepo_tracerCarrier" stroke-linecap="round" stroke-linejoin="round"></g>
	<g id="SVGRepo_iconCarrier">
		<circle opacity="0.5" cx="12" cy="12" r="10" stroke="#1C274C" stroke-width="1.5"></circle>
		<path d="M14.5 9.50002L9.5 14.5M9.49998 9.5L14.5 14.5" stroke="#1C274C" stroke-width="1.5" stroke-linecap="round"></path>
	</g>
</svg>
`
}
