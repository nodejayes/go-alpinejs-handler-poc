package todo_app

import (
	"fmt"
	goalpinejshandler "github.com/nodejayes/go-alpinejs-handler"
	"github.com/nodejayes/go-alpinejs-handler-poc/cosmic_ui"
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
			Value:    "open",
			OnChange: "$store.todo.emit({operation:'toggle',value:id})",
		}),
		DeleteButton: cosmic_ui.NewButton(cosmic_ui.ButtonArguments{
			Content: cosmic_ui.NewText(cosmic_ui.TextArguments{
				Content: "X",
			}),
			OnClick: "$store.todo.emit({operation:'remove',value:id})",
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
		&TodoHandler{},
	}
}

func (ctx *Page) Render() string {
	return `
	<!DOCTYPE html>
	<html>
	<head>
	  <title>Todo App</title>
      {{ template "alpinejs" }}
	  {{ template "alpinejs_handler_lib" }}
	  {{ template "alpinejs_handler_stores" }}
	  {{ .Style "cosmic_ui" }}
	  {{ .Style .ID }}
	</head>
	<body>
	  <div class="app">
		<div class="app-wrapper">
		  <div x-data="{name:''}" class="todo-input">
			<input type="text" x-model="name" />
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
