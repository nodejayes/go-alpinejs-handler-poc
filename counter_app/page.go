package counter_app

import (
	"fmt"
	goalpinejshandler "github.com/nodejayes/go-alpinejs-handler"
)

const pageId = "counter"

func style() string {
	return `
	* {
	  font-family: system-ui;
	  font-size: 15px;
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
	div.counter {
	  display: flex;
	  border: 1px solid grey;
	  border-radius: 4px;
	  background-color: wheat;
	  padding: 3px 6px;
	  max-width: 120px;
	  align-items: center;
	}
	div.counter span {
	  font-size: bold;
	  margin-left: 4px;
	}
	button {
	  background-color: burlywood;
	  border: 1px solid black;
	  border-radius: 4px;
	  min-width: 29px;
	  min-height: 29px;
	  display: flex;
	  justify-content: center;
	  align-items: center;
	  cursor: pointer;
	  margin: 0 4px;
	}
	button:hover {
	  background-color: wheat;
	}
`
}

type Page struct {
	goalpinejshandler.ViewTools
	ID string
}

func NewPage() *Page {
	goalpinejshandler.RegisterStyle(pageId, style())
	return &Page{
		ID: pageId,
	}
}

func (ctx *Page) Name() string {
	return pageId
}

func (ctx *Page) Route() string {
	return fmt.Sprintf("/%s", ctx.Name())
}

func (ctx *Page) Handlers() []goalpinejshandler.ActionHandler {
	return []goalpinejshandler.ActionHandler{
		&CounterHandler{},
	}
}

func (ctx *Page) Render() string {
	return `
	<!DOCTYPE html>
	  <html>
		<head>
		  <title>Counter App</title>
		  {{ template "alpinejs" }}
		  {{ template "alpinejs_handler_lib" }}
		  {{ template "alpinejs_handler_stores" }}
		  {{ .Style .ID }}
		</head>
		<body>
		  <div class="app">
			<button x-data x-init="$store.counter.emit({operation:'get'})" @click="$store.counter.emit({operation:'sub',value:1})">-</button>
			<div class="counter" x-data="$store.counter.state">
			  <p>Counter:</p>
			  <span x-text="value"></span>
			</div>
			<button x-data @click="$store.counter.emit({operation:'add',value:1})">+</button>
		  </div>
		</body>
	  </html>
`
}
