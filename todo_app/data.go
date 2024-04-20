package todo_app

import (
	"github.com/google/uuid"
	contextstore "github.com/nodejayes/context-store"
	"gorm.io/gorm"
)

func init() {
	contextstore.Register(&Todo{})
}

type Todo struct {
	gorm.Model
	ID   string `json:"id"`
	Name string `json:"name"`
	Open bool   `json:"open"`
}

func (ctx *Todo) GetContext() string {
	return "todos"
}

func (ctx *Todo) TableName() string {
	return "todo"
}

func (ctx *Todo) BeforeCreate(tx *gorm.DB) (err error) {
	ctx.ID = uuid.NewString()
	return
}
