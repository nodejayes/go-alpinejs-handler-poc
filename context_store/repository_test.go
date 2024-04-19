package contextstore_test

import (
	"testing"

	"github.com/google/uuid"
	contextstore "github.com/nodejayes/go-alpinejs-handler-poc/context_store"
	"gorm.io/gorm"
)

type TestTodo struct {
	gorm.Model
	ID   string `json:"id"`
	Name string `json:"name"`
	Open bool   `json:"open"`
}

func (ctx *TestTodo) GetContext() string {
	return "todos"
}

func (ctx *TestTodo) TableName() string {
	return "test_todos"
}

func (ctx *TestTodo) BeforeCreate(tx *gorm.DB) error {
	ctx.ID = uuid.NewString()
	return nil
}

func Test_Init(t *testing.T) {
	domain := "testinstance"
	usedModels := make([]contextstore.RepositoryContext, 0)
	usedModels = append(usedModels, &TestTodo{})

	contextstore.Register(usedModels...)
	defer contextstore.Clear()

	err := contextstore.Migrate(domain, true, usedModels...)
	if err != nil {
		t.Error(err)
	}

	todo1 := TestTodo{
		Name: "T1",
		Open: true,
	}
	inserted, err := contextstore.Save(domain, &todo1)
	if err != nil {
		t.Error(err)
	}
	if len(inserted.ID) < 1 {
		t.Errorf("invalid created id %v", inserted.ID)
	}
	if inserted.Name != todo1.Name {
		t.Errorf("expect Todo Name to be %v but was %v", todo1.Name, inserted.Name)
	}

	result, err := contextstore.Get(domain, &TestTodo{}, &TestTodo{
		Name: "T1",
	}, 0, 0)
	if err != nil {
		t.Error(err)
	}
	if len(result) != 1 {
		t.Errorf("expect 1 Todo but have %v", len(result))
	}
	if result[0].Name != todo1.Name {
		t.Errorf("expect Todo Name to be %v but was %v", todo1.Name, result[0].Name)
	}

	todo1.Name = "Todo1"
	updated, err := contextstore.Save(domain, &todo1)
	if err != nil {
		t.Error(err)
	}
	if updated.Name != todo1.Name {
		t.Errorf("expect Todo Name to be %v but was %v", todo1.Name, updated.Name)
	}

	err = contextstore.Archive(domain, &TestTodo{}, todo1.ID)
	if err != nil {
		t.Error(err)
	}
	result, err = contextstore.Get(domain, &TestTodo{}, &TestTodo{
		Name: "Todo1",
	}, 0, 0)
	if err != nil {
		t.Error(err)
	}
	if len(result) > 0 {
		t.Errorf("expect Todo was deleted")
	}

	err = contextstore.Delete(domain, &TestTodo{}, todo1.ID)
	if err != nil {
		t.Error(err)
	}
	result, err = contextstore.Get(domain, &TestTodo{}, &TestTodo{
		Name: "Todo1",
	}, 0, 0)
	if err != nil {
		t.Error(err)
	}
	if len(result) > 0 {
		t.Errorf("expect Todo was deleted")
	}
}
