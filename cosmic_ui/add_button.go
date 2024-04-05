package cosmic_ui

import (
	"fmt"
	goalpinejshandler "github.com/nodejayes/go-alpinejs-handler"
	"github.com/nodejayes/go-alpinejs-handler-poc/cosmic_ui/cosmic_ui_global"
)

func addButtonStyle() string {
	return `
	button.add-button {
	  position: relative;
	  cursor: pointer;
	  display: flex;
	  align-items: center;
	  border: var(--borderSize) solid var(--borderColor);
	  background-color: var(--primaryColor);
	  margin: var(--baseMargin);
	}
	
	button.add-button, span.add-button-icon, span.add-button-text {
	  transition: all 0.3s;
	}
	
	button.add-button span.add-button-text {
	  color: var(--textColor);
	  font-weight: 600;
	  width: calc(100% - 45px);
	  padding: 0 5px;
	  margin: 5px 0;
	  white-space: nowrap;
	  overflow: hidden;
	  text-overflow: ellipsis;
	}
	
	button.add-button span.add-button-icon {
	  position: absolute;
	  height: 100%;
	  width: 35px;
	  background-color: var(--secondaryColor);
	  color: var(--textColor);
	  display: flex;
	  align-items: center;
	  justify-content: center;
	  right: 0;
	}
	
	button.add-button .svg {
	  width: 35px;
	  stroke: vat(--textColor);
	}
	
	button.add-button:hover {
	  background: var(--secondaryColor);
	}
	
	button.add-button:hover span.add-button-text {
	  color: transparent;
	}
	
	button.add-button:hover span.add-button-icon {
	  width: 100%;
	  transform: translateX(0);
	}
	
	button.add-button:active span.add-button-icon {
	  background-color: var(--alternativeColor);
	}
	
	button.add-button:active {
	  border: var(--borderSize) solid var(--alternativeColor);
	}
`
}

type (
	AddButtonArguments struct {
		Data    string
		Label   string
		Width   string
		Height  string
		OnClick string
	}
	AddButton struct {
		goalpinejshandler.ViewTools
		args AddButtonArguments
		Text goalpinejshandler.Component
	}
)

func NewAddButton(args AddButtonArguments) *AddButton {
	goalpinejshandler.RegisterStyle(cosmic_ui_global.PackageID, addButtonStyle())
	return &AddButton{
		args: args,
		Text: NewText(TextArguments{
			Content:   args.Label,
			ClassName: "add-button-text",
		}),
	}
}

func (ctx *AddButton) Name() string {
	return cosmic_ui_global.CreateName("add_button")
}

func (ctx *AddButton) Render() string {
	dataBinding := ""
	if len(ctx.args.Data) > 0 {
		dataBinding = fmt.Sprintf(` x-data="%s"`, ctx.args.Data)
	}
	clickHandler := ""
	if len(ctx.args.OnClick) > 0 {
		clickHandler = fmt.Sprintf(` @click="%s"`, ctx.args.OnClick)
	}
	styleWidth := "150px"
	if len(ctx.args.Width) > 0 {
		styleWidth = ctx.args.Width
	}
	styleHeight := "40px"
	if len(ctx.args.Height) > 0 {
		styleHeight = ctx.args.Height
	}
	styles := fmt.Sprintf(` style="height:%[1]s;width:%[2]s"`, styleHeight, styleWidth)

	return fmt.Sprintf(`
<button type="button" class="add-button"%[1]s%[2]s%[3]s>
	{{ .Paint .Text }}
	<span class="add-button-icon">
	  <svg xmlns="http://www.w3.org/2000/svg" width="24" viewBox="0 0 24 24" stroke-width="2" stroke-linejoin="round" stroke-linecap="round" stroke="currentColor" height="24" fill="none" class="svg">
		<line y2="19" y1="5" x2="12" x1="12"></line>
		<line y2="12" y1="12" x2="19" x1="5"></line>
	  </svg>
	</span>
</button>
`, dataBinding, clickHandler, styles)
}
