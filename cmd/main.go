package main

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/nodejayes/go-alpinejs-handler"
	handler_counter "github.com/nodejayes/go-alpinejs-handler-poc/counter_app/handler"
	template_counter "github.com/nodejayes/go-alpinejs-handler-poc/counter_app/template"
	handler_todo "github.com/nodejayes/go-alpinejs-handler-poc/todo_app/handler"
	template_todo "github.com/nodejayes/go-alpinejs-handler-poc/todo_app/template"
)

func getConfig() goalpinejshandler.Config {
	return goalpinejshandler.Config{
		ActionUrl:            "/action",
		EventUrl:             "/events",
		ClientIDHeaderKey:    "clientId",
		Handlers: []goalpinejshandler.ActionHandler{
			&handler_counter.Counter{},
			&handler_todo.Todo{},
		},
	}
}

func main() {
	config := getConfig()
	router := http.NewServeMux()
	goalpinejshandler.Register(router, &config)


	router.Handle("/counter", templ.Handler(template_counter.Index(goalpinejshandler.HeadScripts())))
	router.Handle("/todo", templ.Handler(template_todo.Index(goalpinejshandler.HeadScripts())))
	
	http.ListenAndServe(":40000", router)
}