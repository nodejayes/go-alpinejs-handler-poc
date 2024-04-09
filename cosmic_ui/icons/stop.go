package icons

import (
	"fmt"
	goalpinejshandler "github.com/nodejayes/go-alpinejs-handler"
	"github.com/nodejayes/go-alpinejs-handler-poc/cosmic_ui/cosmic_ui_global"
)

type (
	StopIcon struct {
		goalpinejshandler.ViewTools
		Color string
	}
	StopIconArguments struct {
		Color string
	}
)

func NewStopIcon(args StopIconArguments) *StopIcon {
	if len(args.Color) < 1 {
		args.Color = "#000000"
	}
	return &StopIcon{
		Color: args.Color,
	}
}

func (ctx *StopIcon) Name() string {
	return cosmic_ui_global.CreateName("stop_icon")
}

func (ctx *StopIcon) Render() string {
	return fmt.Sprintf(`
<svg viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
	<g id="SVGRepo_bgCarrier" stroke-width="0"></g>
	<g id="SVGRepo_tracerCarrier" stroke-linecap="round" stroke-linejoin="round"></g>
	<g id="SVGRepo_iconCarrier">
		<circle cx="12" cy="12" r="10" stroke="%[1]s" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"></circle>
		<path d="M5 19L19 5" stroke="%[1]s" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"></path>
	</g>
</svg>
`, ctx.Color)
}
