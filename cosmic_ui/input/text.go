package input

import (
	"fmt"

	goalpinejshandler "github.com/nodejayes/go-alpinejs-handler"
	"github.com/nodejayes/go-alpinejs-handler-poc/cosmic_ui/cosmic_ui_global"
)

func textInputStyle() string {
	return `
	div.cosmic-ui-input-wrapper {
		display: flex;
		align-items: center;
		margin: var(--baseMargin);
		height: var(--baseElementHeight);
	}
	`
}

type (
	TextInput struct {
		goalpinejshandler.ViewTools
		args TextInputArguments
	}
	TextInputArguments struct {
		Placeholder string
		Model       string
		Label       string
		Hint        string
		Hide        bool
		OnChange    string
	}
)

func NewTextInput(args TextInputArguments) *TextInput {
	goalpinejshandler.RegisterStyle(cosmic_ui_global.PackageID, textInputStyle())
	return &TextInput{
		args: args,
	}
}

func (ctx *TextInput) Name() string {
	return cosmic_ui_global.CreateName("text_input")
}

func (ctx *TextInput) Render() string {
	placeholder := ""
	if len(ctx.args.Placeholder) > 0 {
		placeholder = fmt.Sprintf(` placeholder="%s"`, ctx.args.Placeholder)
	}
	model := ""
	if len(ctx.args.Model) > 0 {
		model = fmt.Sprintf(` x-model="%s"`, ctx.args.Model)
	}
	inputType := "text"
	if ctx.args.Hide {
		inputType = "password"
	}
	hint := ""
	if len(ctx.args.Hint) > 0 {
		hint = fmt.Sprintf(` x-text="%s"`, ctx.args.Hint)
	}
	label := ""
	if len(ctx.args.Label) > 0 {
		label = fmt.Sprintf(` x-text="%s"`, ctx.args.Label)
	}
	changeEvent := ""
	if len(ctx.args.OnChange) > 0 {
		changeEvent = fmt.Sprintf(` @change="%s"`, ctx.args.OnChange)
	}
	return fmt.Sprintf(`
	<div class="cosmic-ui-input-wrapper">
		<span class="cosmic-ui-input-label"%[5]s></span>
		<input type="%[3]s"%[1]s%[2]s%[6]s />
		<span class="cosmic-ui-input-hint"%[4]s></span>
	</div>
	`, placeholder, model, inputType, hint, label, changeEvent)
}
