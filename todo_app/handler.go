package todo_app

import (
	"fmt"
	"net/http"
	"time"

	contextstore "github.com/nodejayes/context-store"
	"github.com/nodejayes/go-alpinejs-handler-poc/cosmic_ui/toaster"

	anythingparsejson "github.com/nodejayes/anything-parse-json"
	di "github.com/nodejayes/generic-di"
	goalpinejshandler "github.com/nodejayes/go-alpinejs-handler"
)

func init() {
	di.Injectable(NewTodoHandler)
}

type (
	TodoHandler struct {
		Toaster *toaster.Handler
	}
	TodoArguments struct {
		Operation string `json:"operation"`
		Value     any    `json:"value"`
	}
)

func NewTodoHandler() *TodoHandler {
	return &TodoHandler{
		Toaster: di.Inject[toaster.Handler](),
	}
}

func (ctx *TodoHandler) GetName() string {
	return "todo"
}

func (ctx *TodoHandler) GetActionType() string {
	return fmt.Sprintf("[%s] operation", ctx.GetName())
}

func (ctx *TodoHandler) GetDefaultState() any {
	return NewTodoState()
}

func (ctx *TodoHandler) OnDestroy(clientID string, tools *goalpinejshandler.Tools) {
	go func() {
		time.Sleep(60 * time.Second)
		if tools.HasConnections(clientID) {
			return
		}
		di.Destroy[State](clientID)
	}()
}

func (ctx *TodoHandler) Handle(msg goalpinejshandler.Message, res http.ResponseWriter, req *http.Request, messagePool *goalpinejshandler.MessagePool, tools *goalpinejshandler.Tools) {
	args, err := anythingparsejson.Parse[TodoArguments](msg.Payload)
	if err != nil {
		return
	}
	clientId := tools.GetClientId(req)
	if len(clientId) < 1 {
		return
	}
	contextstore.Migrate(clientId, false, &Todo{})
	state := di.Inject[State](clientId)

	switch args.Operation {
	case "load":
		if !ctx.loadTodos(clientId, res, req, messagePool, tools, state) {
			return
		}
	case "add":
		if !ctx.addTodo(clientId, args, res, req, messagePool, tools, state) {
			return
		}
	case "remove":
		if !ctx.removeTodo(clientId, args, res, req, messagePool, tools, state) {
			return
		}
	case "toggle":
		if !ctx.toggleTodo(clientId, args, res, req, messagePool, tools, state) {
			return
		}
	}

	messagePool.Add(goalpinejshandler.ChannelMessage{
		ClientFilter: func(client goalpinejshandler.Client) bool {
			return client.ID == clientId
		},
		Message: goalpinejshandler.Message{
			Type:    fmt.Sprintf("[%s] update", ctx.GetName()),
			Payload: state,
		},
	})
}

func (ctx *TodoHandler) loadTodos(clientId string, res http.ResponseWriter, req *http.Request, messagePool *goalpinejshandler.MessagePool, tools *goalpinejshandler.Tools, state *State) bool {
	todos, err := contextstore.Get(clientId, &Todo{}, func(builder contextstore.ConditionBuilder) contextstore.ConditionBuilder {
		return builder
	})
	if err != nil {
		ctx.sendMessage(toaster.DangerType, err.Error(), res, req, messagePool, tools)
		return false
	}
	state.LoadTodos(todos)
	return true
}

func (ctx *TodoHandler) addTodo(clientId string, args TodoArguments, res http.ResponseWriter, req *http.Request, messagePool *goalpinejshandler.MessagePool, tools *goalpinejshandler.Tools, state *State) bool {
	todoToAdd, err := anythingparsejson.Parse[*Todo](args.Value)
	if err != nil {
		return false
	}
	if len(todoToAdd.Name) < 3 {
		ctx.sendMessage(toaster.DangerType, "Activity must have a label with min 3 Chars", res, req, messagePool, tools)
		return false
	}
	todoToAdd.Open = true
	todoToAdd, err = contextstore.Save(clientId, todoToAdd)
	if err != nil {
		ctx.sendMessage(toaster.DangerType, err.Error(), res, req, messagePool, tools)
		return false
	}
	state.Add(todoToAdd)
	ctx.sendMessage(toaster.SuccessType, fmt.Sprintf("Activity %s added", todoToAdd.Name), res, req, messagePool, tools)
	return true
}

func (ctx *TodoHandler) removeTodo(clientId string, args TodoArguments, res http.ResponseWriter, req *http.Request, messagePool *goalpinejshandler.MessagePool, tools *goalpinejshandler.Tools, state *State) bool {
	todoToRemoveId, err := anythingparsejson.Parse[string](args.Value)
	if err != nil {
		ctx.sendMessage(toaster.DangerType, err.Error(), res, req, messagePool, tools)
		return false
	}
	err = contextstore.Delete(clientId, &Todo{}, []string{todoToRemoveId})
	if err != nil {
		ctx.sendMessage(toaster.DangerType, err.Error(), res, req, messagePool, tools)
		return false
	}
	state.Remove(todoToRemoveId)
	return true
}

func (ctx *TodoHandler) toggleTodo(clientId string, args TodoArguments, res http.ResponseWriter, req *http.Request, messagePool *goalpinejshandler.MessagePool, tools *goalpinejshandler.Tools, state *State) bool {
	todoToToggleId, err := anythingparsejson.Parse[string](args.Value)
	if err != nil {
		ctx.sendMessage(toaster.DangerType, err.Error(), res, req, messagePool, tools)
		return false
	}
	state.Toggle(todoToToggleId)
	todo, err := state.Get(todoToToggleId)
	if err != nil {
		ctx.sendMessage(toaster.DangerType, err.Error(), res, req, messagePool, tools)
		return false
	}
	todo, err = contextstore.Save(clientId, todo)
	if err != nil {
		ctx.sendMessage(toaster.DangerType, err.Error(), res, req, messagePool, tools)
		return false
	}
	activeLabel := "open"
	if !todo.Open {
		activeLabel = "finish"
	}
	ctx.sendMessage(toaster.SuccessType, fmt.Sprintf("State of Activity %s was set to %s", todo.Name, activeLabel), res, req, messagePool, tools)
	return true
}

func (ctx *TodoHandler) sendMessage(messageTyp, message string, res http.ResponseWriter, req *http.Request, messagePool *goalpinejshandler.MessagePool, tools *goalpinejshandler.Tools) {
	ctx.Toaster.Handle(goalpinejshandler.Message{
		Type: ctx.Toaster.GetActionType(),
		Payload: toaster.HandlerArguments{
			Operation: "add",
			Value: toaster.Message{
				Typ:     messageTyp,
				Message: message,
			},
		},
	}, res, req, messagePool, tools)
}
