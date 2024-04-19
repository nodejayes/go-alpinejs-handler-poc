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

func TestInit(t *testing.T) {
	domain := "TestInit"
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

	result, err := contextstore.Get(domain, &TestTodo{}, func(builder contextstore.ConditionBuilder) contextstore.ConditionBuilder {
		return builder.Where(&TestTodo{Name: "T1"})
	})
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
	result, err = contextstore.Get(domain, &TestTodo{}, func(builder contextstore.ConditionBuilder) contextstore.ConditionBuilder {
		return builder.Where(&TestTodo{Name: "Todo1"})
	})
	if err != nil {
		t.Error(err)
	}
	if len(result) > 0 {
		t.Errorf("expect Todo was archived")
	}
	archiveResult, err := contextstore.GetArchive(domain, &TestTodo{}, func(builder contextstore.ConditionBuilder) contextstore.ConditionBuilder {
		return builder.Where(&TestTodo{Name: "Todo1"})
	})
	if err != nil {
		t.Error(err)
	}
	archiveLen := len(archiveResult)
	if archiveLen != 1 {
		t.Errorf("expect Todo was found in archive")
	}
	if archiveResult[0].ID != todo1.ID {
		t.Errorf("expect the archived Todo to found in archive but ID %v not equals %v", archiveResult[0].ID, todo1.ID)
	}

	err = contextstore.Delete(domain, &TestTodo{}, todo1.ID)
	if err != nil {
		t.Error(err)
	}
	result, err = contextstore.Get(domain, &TestTodo{}, func(builder contextstore.ConditionBuilder) contextstore.ConditionBuilder {
		return builder.Where(&TestTodo{Name: "Todo1"})
	})
	if err != nil {
		t.Error(err)
	}
	if len(result) > 0 {
		t.Errorf("expect Todo was deleted")
	}
}

func TestBulkCreate(t *testing.T) {
	domain := "TestBulkCreate"
	todos := []*TestTodo{
		{
			Name: "T1",
			Open: true,
		},
		{
			Name: "T2",
			Open: true,
		},
		{
			Name: "T3",
			Open: true,
		},
	}

	contextstore.Register(&TestTodo{})
	defer contextstore.Clear()
	contextstore.Migrate(domain, true, &TestTodo{})

	inserted, err := contextstore.BulkCreate(domain, &TestTodo{}, todos)
	if err != nil {
		t.Error(err)
	}
	if len(inserted) != len(todos) {
		t.Errorf("%v/%v todos inserted", len(inserted), len(todos))
	}
	for _, todo := range inserted {
		if len(todo.ID) < 1 {
			t.Errorf("missing ID on Todo %v", todo.Name)
		}
	}

	sel, err := contextstore.Get(domain, &TestTodo{}, func(builder contextstore.ConditionBuilder) contextstore.ConditionBuilder {
		return builder.Where("name IN ?", "T1", "T3")
	})
	if err != nil {
		t.Error(err)
	}
	if len(sel) != 2 {
		t.Errorf("%v/%v todos selected", len(sel), 2)
	}
	for _, selection := range sel {
		if selection.Name != "T1" && selection.Name != "T3" {
			t.Errorf("only T1 or T3 must in selection name found %v", selection.Name)
		}
	}
}
