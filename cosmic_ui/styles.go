package cosmic_ui

import (
	"bytes"
	"context"
	"io"
)

func GetStyles() string {
	buf := bytes.NewBuffer([]byte{})
	useGlobalStyle(buf)
	useComponentsStyle(buf)
	return buf.String()
}

func useGlobalStyle(w io.Writer) {
	err := GetGlobalStyle().Render(context.Background(), w)
	if err != nil {
		println(err.Error())
	}
}

func useComponentsStyle(w io.Writer) {
	err := AddButtonStyle().Render(context.Background(), w)
	if err != nil {
		println(err.Error())
	}

	err = ButtonStyle().Render(context.Background(), w)
	if err != nil {
		println(err.Error())
	}

	err = CheckboxStyle().Render(context.Background(), w)
	if err != nil {
		println(err.Error())
	}
}
