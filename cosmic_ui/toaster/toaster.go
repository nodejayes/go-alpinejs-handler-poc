package toaster

import (
	"fmt"
	goalpinejshandler "github.com/nodejayes/go-alpinejs-handler"
	"github.com/nodejayes/go-alpinejs-handler-poc/cosmic_ui/cosmic_ui_global"
)

func style() string {
	return `
div.cosmic-ui-toaster-wrapper {
	position: absolute;
    bottom: 5px;
    right: 5px;
}
`
}

type Toaster struct {
	goalpinejshandler.ViewTools
}

func NewToaster() *Toaster {
	goalpinejshandler.RegisterStyle(cosmic_ui_global.PackageID, style())
	return &Toaster{}
}

func (ctx *Toaster) Name() string {
	return cosmic_ui_global.CreateName("toaster")
}

func (ctx *Toaster) Render() string {
	return fmt.Sprintf(`
<div class="cosmic-ui-toaster-wrapper" x-data="$store.%[1]s.state">
	<template x-for="msg in messages">
		<ul>
			<li x-text="msg.message"></li>
		</ul>
	</template>
</div>
`, name)
}
