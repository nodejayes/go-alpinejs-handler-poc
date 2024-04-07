package icons

import (
	goalpinejshandler "github.com/nodejayes/go-alpinejs-handler"
	"github.com/nodejayes/go-alpinejs-handler-poc/cosmic_ui/cosmic_ui_global"
)

type StopIcon struct {
	goalpinejshandler.ViewTools
}

func NewStopIcon() *StopIcon {
	return &StopIcon{}
}

func (ctx *StopIcon) Name() string {
	return cosmic_ui_global.CreateName("stop_icon")
}

func (ctx *StopIcon) Render() string {
	return `
<svg viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
	<g id="SVGRepo_bgCarrier" stroke-width="0"></g>
	<g id="SVGRepo_tracerCarrier" stroke-linecap="round" stroke-linejoin="round"></g>
	<g id="SVGRepo_iconCarrier">
		<circle cx="12" cy="12" r="10" stroke="#000000" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"></circle>
		<path d="M5 19L19 5" stroke="#000000" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"></path>
	</g>
</svg>
`
}
