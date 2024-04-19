package todo_app

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	ID   string `json:"id"`
	Name string `json:"name"`
	Open bool   `json:"open"`
}

func (ctx *Todo) GetContext() string {
	return "todo"
}
