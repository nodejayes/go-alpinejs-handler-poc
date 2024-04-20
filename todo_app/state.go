package todo_app

import (
	"fmt"
	"slices"

	di "github.com/nodejayes/generic-di"
)

func init() {
	di.Injectable(NewTodoState)
}

type State struct {
	Todos []*Todo `json:"todos"`
}

func NewTodoState() *State {
	return &State{
		Todos: make([]*Todo, 0),
	}
}

func (ctx *State) LoadTodos(todos []*Todo) {
	ctx.Todos = todos
}

func (ctx *State) Get(id string) (*Todo, error) {
	idx := slices.IndexFunc(ctx.Todos, func(todo *Todo) bool { return todo.ID == id })
	if idx < 0 || idx > len(ctx.Todos)-1 {
		return nil, fmt.Errorf("todo was not found")
	}
	return ctx.Todos[idx], nil
}

func (ctx *State) Add(todo *Todo) {
	ctx.Todos = append(ctx.Todos, todo)
}

func (ctx *State) Remove(id string) {
	idx := slices.IndexFunc(ctx.Todos, func(todo *Todo) bool { return todo.ID == id })
	if idx < 0 {
		return
	}
	ctx.Todos = append(ctx.Todos[:idx], ctx.Todos[idx+1:]...)
}

func (ctx *State) Toggle(id string) {
	idx := slices.IndexFunc(ctx.Todos, func(todo *Todo) bool { return todo.ID == id })
	if idx < 0 {
		return
	}
	ctx.Todos[idx].Open = !ctx.Todos[idx].Open
}
