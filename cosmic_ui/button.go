package cosmic_ui

import (
	"fmt"
	goalpinejshandler "github.com/nodejayes/go-alpinejs-handler"
	"github.com/nodejayes/go-alpinejs-handler-poc/cosmic_ui/cosmic_ui_global"
)

func buttonStyle() string {
	return `
	button.button {
	  background: rgba(207, 205, 205, 0.425);
	  border: none;
	  display: inline-block;
	  font-size: 15px;
	  font-weight: 600;
	  text-transform: uppercase;
	  cursor: pointer;
	  transform: skew(-21deg);
	}
	
	button.button span {
	  display: inline-block;
	  transform: skew(21deg);
	}
	
	button.button::before {
	  content: '';
	  position: absolute;
	  top: 0;
	  bottom: 0;
	  right: 100%;
	  left: 0;
	  background: #3aa856;
	  opacity: 0;
	  z-index: -1;
	  transition: all 0.5s;
	}
	
	button.button:hover {
	  color: #fff;
	}
	
	button.button:hover::before {
	  left: 0;
	  right: 0;
	  opacity: 1;
	}`
}

type (
	ButtonArguments struct {
		Data    string
		Content goalpinejshandler.Component
		Width   string
		Height  string
		OnClick string
	}
	Button struct {
		goalpinejshandler.ViewTools
		Args ButtonArguments
	}
)

func NewButton(args ButtonArguments) *Button {
	goalpinejshandler.RegisterStyle(cosmic_ui_global.PackageID, buttonStyle())
	return &Button{
		Args: args,
	}
}

func (ctx *Button) Name() string {
	return cosmic_ui_global.CreateName("button")
}

func (ctx *Button) Render() string {
	bindingData := ""
	if len(ctx.Args.Data) > 0 {
		bindingData = fmt.Sprintf(` x-data="%s"`, ctx.Args.Data)
	}
	onClickHandler := ""
	if len(ctx.Args.OnClick) > 0 {
		onClickHandler = fmt.Sprintf(` @click="%s"`, ctx.Args.OnClick)
	}
	styleWidth := "150px"
	if len(ctx.Args.Width) > 0 {
		styleWidth = ctx.Args.Width
	}
	styleHeight := "40px"
	if len(ctx.Args.Height) > 0 {
		styleHeight = ctx.Args.Height
	}
	styles := fmt.Sprintf(` style="height:%[1]s;width:%[2]s"`, styleHeight, styleWidth)

	return fmt.Sprintf(`
<button %[1]s class="button"%[2]s%[3]s>
	{{ .Paint .Args.Content }}
</button>`, styles, bindingData, onClickHandler)
}
