package state

import (
	"slices"

	di "github.com/nodejayes/generic-di"
	"github.com/nodejayes/go-alpinejs-handler-poc/todo_app/data"
)

func init() {
	di.Injectable(NewTodoState)
}

type Todo struct {
	Todos []data.Todo `json:"todos"`
}

func NewTodoState() *Todo {
	return &Todo{
		Todos: make([]data.Todo, 0),
	}
}

func (ctx *Todo) Add(todo data.Todo) {
	ctx.Todos = append(ctx.Todos, todo)
}

func (ctx *Todo) Remove(id string) {
	idx := slices.IndexFunc(ctx.Todos, func(todo data.Todo) bool { return todo.ID == id })
	if idx < 0 {
		return
	}
	ctx.Todos = append(ctx.Todos[:idx], ctx.Todos[idx+1:]...)
}

func (ctx *Todo) Toggle(id string) {
	idx := slices.IndexFunc(ctx.Todos, func(todo data.Todo) bool { return todo.ID == id })
	if idx < 0 {
		return
	}
	ctx.Todos[idx].Open = !ctx.Todos[idx].Open
}
