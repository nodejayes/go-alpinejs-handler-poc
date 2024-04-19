package todo_app

import (
	"fmt"

	di "github.com/nodejayes/generic-di"
	goalpinejshandler "github.com/nodejayes/go-alpinejs-handler"
	"github.com/nodejayes/go-alpinejs-handler-poc/cosmic_ui"
	"github.com/nodejayes/go-alpinejs-handler-poc/cosmic_ui/input"
	"github.com/nodejayes/go-alpinejs-handler-poc/cosmic_ui/toaster"
)

const pageId = "todo"

func style() string {
	return ``
}

type Page struct {
	goalpinejshandler.ViewTools
	ID           string
	AddButton    goalpinejshandler.Component
	TodoCheckbox goalpinejshandler.Component
	DeleteButton goalpinejshandler.Component
	Toaster      goalpinejshandler.Component
	ActionInput  goalpinejshandler.Component
}

func NewPage() *Page {
	goalpinejshandler.RegisterStyle(pageId, style())
	return &Page{
		ID: pageId,
		AddButton: cosmic_ui.NewAddButton(cosmic_ui.AddButtonArguments{
			Label:   "hinzufügen",
			Width:   "150px",
			Height:  "40px",
			OnClick: "$store.todo.emit({operation:'add',value:{id:'',name:name,open:false}})",
		}),
		TodoCheckbox: cosmic_ui.NewCheckbox(cosmic_ui.CheckboxArguments{
			ID:       "id",
			Label:    "name",
			Value:    "!open",
			OnChange: "$store.todo.emit({operation:'toggle',value:id})",
		}),
		DeleteButton: cosmic_ui.NewButton(cosmic_ui.ButtonArguments{
			Content: cosmic_ui.NewText(cosmic_ui.TextArguments{
				Content: "X",
			}),
			OnClick: "$store.todo.emit({operation:'remove',value:id})",
		}),
		Toaster: toaster.NewToaster(),
		ActionInput: input.NewTextInput(input.TextInputArguments{
			Placeholder: "Todo",
			Model:       "name",
		}),
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
		di.Inject[TodoHandler](),
		di.Inject[toaster.Handler](),
	}
}

func (ctx *Page) Render() string {
	return `
	<!DOCTYPE html>
	<html>
	<head>
	  <title>Todo App</title>
		<link rel="icon" type="image/x-icon" href="/icons/todo_favicon.ico">
    {{ template "alpinejs" }}
	  {{ template "alpinejs_handler_lib" }}
	  {{ template "alpinejs_handler_stores" }}
	  {{ .Style "cosmic_ui" }}
	  {{ .Style .ID }}
	</head>
	<body>
	  <div class="app">
		{{ .Paint .Toaster }}
		<div class="app-wrapper">
		  <div x-data="{name:''}" class="todo-input">
				{{ .Paint .ActionInput }}
				{{ .Paint .AddButton }}
		  </div>
		  <div x-data="$store.todo.state" x-init="$store.todo.emit({operation:'get'})" class="todo-list">
			<template x-for="todo in todos">
			  <span x-data="todo" class="todo-display">
			  {{ .Paint .TodoCheckbox }}
			  {{ .Paint .DeleteButton }}
			  </span>
			</template>
		  </div>
		</div>
	  </div>
	</body>
	</html>
`
}
