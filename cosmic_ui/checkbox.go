package cosmic_ui

import (
	"fmt"
	goalpinejshandler "github.com/nodejayes/go-alpinejs-handler"
	"github.com/nodejayes/go-alpinejs-handler-poc/cosmic_ui/cosmic_ui_global"
)

func checkboxStyle() string {
	return `
	.checkbox-wrapper input[type="checkbox"] {
      display: none;
    }

    .checkbox-wrapper .terms-label {
      cursor: pointer;
      display: flex;
      align-items: center;
    }

    .checkbox-wrapper .terms-label .label-text {
      margin-left: 10px;
    }

    .checkbox-wrapper .checkbox-svg {
      width: 30px;
      height: 30px;
    }

    .checkbox-wrapper .checkbox-box {
      fill: rgba(207, 205, 205, 0.425);
      stroke: #3aa856;
      stroke-dasharray: 800;
      stroke-dashoffset: 800;
      transition: stroke-dashoffset 0.6s ease-in;
    }

    .checkbox-wrapper .checkbox-tick {
      stroke: #3aa856;
      stroke-dasharray: 172;
      stroke-dashoffset: 172;
      transition: stroke-dashoffset 0.6s ease-in;
    }

    .checkbox-wrapper input[type="checkbox"]:checked + .terms-label .checkbox-box,
      .checkbox-wrapper input[type="checkbox"]:checked + .terms-label .checkbox-tick {
      stroke-dashoffset: 0;
    }`
}

type (
	CheckboxArguments struct {
		Data           string
		ID             string
		Label          string
		Value          string
		OnChange       string
		CheckboxHeight int
		CheckboxWidth  int
	}
	Checkbox struct {
		goalpinejshandler.ViewTools
		Args CheckboxArguments
	}
)

func NewCheckbox(args CheckboxArguments) *Checkbox {
	goalpinejshandler.RegisterStyle(cosmic_ui_global.PackageID, checkboxStyle())
	return &Checkbox{
		Args: args,
	}
}

func (ctx *Checkbox) Name() string {
	return cosmic_ui_global.CreateName("checkbox")
}

func (ctx *Checkbox) Render() string {
	bindingData := ""
	if len(ctx.Args.Data) > 0 {
		bindingData = fmt.Sprintf(` x-data="%s"`, ctx.Args.Data)
	}
	if ctx.Args.CheckboxWidth <= 0 {
		ctx.Args.CheckboxWidth = 200
	}
	if ctx.Args.CheckboxHeight <= 0 {
		ctx.Args.CheckboxHeight = 200
	}

	return fmt.Sprintf(`
<div class="checkbox-wrapper"%[1]s>
    <input :id="{{ .Args.ID }}" :name="{{ .Args.ID }}" type="checkbox" :value="{{ .Args.Value }}" @change="{{ .Args.OnChange }}" />
    <label class="terms-label" :for="{{ .Args.ID }}">
      <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 {{ .Args.CheckboxWidth }} {{ .Args.CheckboxHeight }}" class="checkbox-svg">
        <mask fill="white" id="path-1-inside-1_476_5-37">
          <rect height="{{ .Args.CheckboxHeight }}" width="{{ .Args.CheckboxWidth }}"></rect>
        </mask>
        <rect mask="url(#path-1-inside-1_476_5-37)" stroke-width="40" class="checkbox-box" height="{{ .Args.CheckboxHeight }}" width="{{ .Args.CheckboxWidth }}"></rect>
        <path stroke-width="15" d="M52 111.018L76.9867 136L149 64" class="checkbox-tick"></path>
      </svg>
      <span class="label-text" x-text="{{ .Args.Label }}"></span>
    </label>
  </div>`, bindingData)
}
