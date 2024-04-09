package toaster

import (
	"fmt"
	goalpinejshandler "github.com/nodejayes/go-alpinejs-handler"
	"github.com/nodejayes/go-alpinejs-handler-poc/cosmic_ui"
	"github.com/nodejayes/go-alpinejs-handler-poc/cosmic_ui/cosmic_ui_global"
	"github.com/nodejayes/go-alpinejs-handler-poc/cosmic_ui/icons"
)

func style() string {
	return `
div.cosmic-ui-toaster-wrapper {
	position: absolute;
    bottom: var(--basePadding);
    right: var(--basePadding);
}
div.cosmic-ui-toaster-wrapper div.message {
	display: flex;
    align-items: center;
	border-radius: var(--borderRadius);
	margin: var(--baseMargin);
	padding: var(--basePadding);
}
div.cosmic-ui-toaster-wrapper div.message.danger {
	border: var(--borderSize) solid var(--dangerColor);
	border-left: 5px solid var(--dangerColor);
}
div.cosmic-ui-toaster-wrapper div.message.warning {
	border: var(--borderSize) solid var(--warningColor);
	border-left: 5px solid var(--warningColor);
}
div.cosmic-ui-toaster-wrapper div.message.success {
	border: var(--borderSize) solid var(--successColor);
	border-left: 5px solid var(--successColor);
}
div.cosmic-ui-toaster-wrapper div.message div.icon {
	width: 25px;
	height: 25px;
	margin: var(--baseMargin);
}
div.cosmic-ui-toaster-wrapper div.message div.close {
	width: 25px;
	height: 25px;
	margin: var(--baseMargin);
}
.cosmic-ui-toaster-animate-in-base {
	transition: transform 300ms ease-in;
}
.cosmic-ui-toaster-animate-out-base {
	transition: transform 300ms ease-out;
}
.cosmic-ui-toaster-animate-in {
	translate: transformX(0);
}
.cosmic-ui-toaster-animate-out {
	translate: transformX(200%);
}
`
}

type Toaster struct {
	goalpinejshandler.ViewTools
	SuccessIcon   goalpinejshandler.Component
	AttentionIcon goalpinejshandler.Component
	StopIcon      goalpinejshandler.Component
	CloseButton   goalpinejshandler.Component
}

func NewToaster() *Toaster {
	goalpinejshandler.RegisterStyle(cosmic_ui_global.PackageID, style())
	return &Toaster{
		SuccessIcon:   icons.NewSuccessIcon(icons.SuccessIconArguments{Color: "green"}),
		AttentionIcon: icons.NewAttentionIcon(icons.AttentionIconArguments{Color: "orange"}),
		StopIcon:      icons.NewStopIcon(icons.StopIconArguments{Color: "red"}),
		CloseButton: cosmic_ui.NewButton(cosmic_ui.ButtonArguments{
			Content: cosmic_ui.NewText(cosmic_ui.TextArguments{
				Content: "X",
			}),
			Width:   "25px",
			Height:  "25px",
			OnClick: fmt.Sprintf(`$store.%[1]s.emit({operation:'remove',value:msg})`, name),
		}),
	}
}

func (ctx *Toaster) Name() string {
	return cosmic_ui_global.CreateName("toaster")
}

func (ctx *Toaster) Render() string {
	return fmt.Sprintf(`
<div class="cosmic-ui-toaster-wrapper" x-data="$store.%[1]s.state">
	<template x-for="msg in messages" :key="msg.id">
		<div class="message" :class="msg.typ"
			 x-show="msg.open"
             x-init="$nextTick(()=>{msg.open=true;$store.%[1]s.emit({operation:'animation_start',value:msg})})"
             x-transition.duration.300>
			<div class="icon">
				<template x-if="msg.typ === 'danger'">
					{{ .Paint .StopIcon }}
				</template>
				<template x-if="msg.typ === 'warning'">
					{{ .Paint .AttentionIcon }}
				</template>
				<template x-if="msg.typ === 'success'">
					{{ .Paint .SuccessIcon }}
				</template>
			</div>
			<div class="label">
				<span x-text="msg.message"></span>
			</div>
			<div class="close">
				{{ .Paint .CloseButton }}
			</div>
		</div>
	</template>
</div>
`, name)
}
