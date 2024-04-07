package icons

import (
	goalpinejshandler "github.com/nodejayes/go-alpinejs-handler"
	"github.com/nodejayes/go-alpinejs-handler-poc/cosmic_ui/cosmic_ui_global"
)

type SuccessIcon struct {
	goalpinejshandler.ViewTools
}

func NewSuccessIcon() *SuccessIcon {
	return &SuccessIcon{}
}

func (ctx *SuccessIcon) Name() string {
	return cosmic_ui_global.CreateName("success_icon")
}

func (ctx *SuccessIcon) Render() string {
	return `
<svg fill="#000000" viewBox="0 0 36 36" version="1.1" preserveAspectRatio="xMidYMid meet" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink">
	<g id="SVGRepo_bgCarrier" stroke-width="0"></g>
	<g id="SVGRepo_tracerCarrier" stroke-linecap="round" stroke-linejoin="round"></g>
	<g id="SVGRepo_iconCarrier">
		<title>success-standard-line</title>
		<path class="clr-i-outline clr-i-outline-path-1" d="M18,2A16,16,0,1,0,34,18,16,16,0,0,0,18,2Zm0,30A14,14,0,1,1,32,18,14,14,0,0,1,18,32Z"></path>
		<path class="clr-i-outline clr-i-outline-path-2" d="M28,12.1a1,1,0,0,0-1.41,0L15.49,23.15l-6-6A1,1,0,0,0,8,18.53L15.49,26,28,13.52A1,1,0,0,0,28,12.1Z"></path>
		<rect x="0" y="0" width="36" height="36" fill-opacity="0"></rect>
	</g>
</svg>
`
}
