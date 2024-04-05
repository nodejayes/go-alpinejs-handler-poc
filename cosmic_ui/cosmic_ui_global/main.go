package cosmic_ui_global

import (
	"fmt"
	goalpinejshandler "github.com/nodejayes/go-alpinejs-handler"
)

const PackageID = "cosmic_ui"

func RegisterGlobalStyles() {
	goalpinejshandler.RegisterStyle(PackageID, `
	:root {
		--fontFamily: system-ui;
		--fontSize: 15px;
		--borderColor: #34974d;
		--borderSize: 1px;
		--primaryColor: #3aa856;
		--secondaryColor: #34974d;
		--textColor: #fff;
		--alternativeColor: #2e8644;
		--baseMargin: 5px 8px;
	}
	* {
	  font-family: var(--fontFamily);
	  font-size: var(--fontSize);
	  margin: 0;
	  padding: 0;
	}
	html, body {
	  width: 100vw;
	  height: 100vh;
	}
	div.app {
	  display: flex;
	  width: 100vw;
	  height: 100vh;
	  align-items: center;
	  justify-content: center;
	}
	div.todo-input {
	  display: flex;
	}
	span.todo-display {
	  display: flex;
	}`)
}

func CreateName(name string) string {
	return fmt.Sprintf("%s_%s", PackageID, name)
}
