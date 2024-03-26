// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package cosmic_ui

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

type AddButtonArguments struct {
	Data    string
	Label   string
	OnClick string
}

func AddButtonStyle() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templ.Raw(`
  <style type="text/css">
    button.add-button {
      position: relative;
      width: 150px;
      height: 40px;
      cursor: pointer;
      display: flex;
      align-items: center;
      border: 1px solid #34974d;
      background-color: #3aa856;
    }

    button.add-button, span.add-button-icon, span.add-button-text {
      transition: all 0.3s;
    }

    button.add-button span.add-button-text {
      transform: translateX(30px);
      color: #fff;
      font-weight: 600;
    }

    button.add-button span.add-button-icon {
      position: absolute;
      transform: translateX(109px);
      height: 100%;
      width: 39px;
      background-color: #34974d;
      display: flex;
      align-items: center;
      justify-content: center;
    }

    button.add-button .svg {
      width: 30px;
      stroke: #fff;
    }

    button.add-button:hover {
      background: #34974d;
    }

    button.add-button:hover span.add-button-text {
      color: transparent;
    }

    button.add-button:hover span.add-button-icon {
      width: 148px;
      transform: translateX(0);
    }

    button.add-button:active span.add-button-icon {
      background-color: #2e8644;
    }

    button.add-button:active {
      border: 1px solid #2e8644;
    }
  </style>
  `).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func AddButton(args AddButtonArguments) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var2 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var2 == nil {
			templ_7745c5c3_Var2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<button type=\"button\" class=\"add-button\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if len(args.Data) > 0 {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" x-data=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(args.Data))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" @click=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(args.OnClick))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><span class=\"add-button-text\" x-text=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(args.Label))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"></span> <span class=\"add-button-icon\"><svg xmlns=\"http://www.w3.org/2000/svg\" width=\"24\" viewBox=\"0 0 24 24\" stroke-width=\"2\" stroke-linejoin=\"round\" stroke-linecap=\"round\" stroke=\"currentColor\" height=\"24\" fill=\"none\" class=\"svg\"><line y2=\"19\" y1=\"5\" x2=\"12\" x1=\"12\"></line> <line y2=\"12\" y1=\"12\" x2=\"19\" x1=\"5\"></line></svg></span></button>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}