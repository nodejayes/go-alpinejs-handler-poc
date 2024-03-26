package main

import (
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/nodejayes/go-alpinejs-handler"
	handler_counter "github.com/nodejayes/go-alpinejs-handler-poc/counter_app/handler"
	template_counter "github.com/nodejayes/go-alpinejs-handler-poc/counter_app/template"
	handler_todo "github.com/nodejayes/go-alpinejs-handler-poc/todo_app/handler"
	template_todo "github.com/nodejayes/go-alpinejs-handler-poc/todo_app/template"
)

func getHandler() []goalpinejshandler.ActionHandler {
	return []goalpinejshandler.ActionHandler{
		&handler_counter.Counter{},
		&handler_todo.Todo{},
	}
}

func getConfig() goalpinejshandler.Config {
	return goalpinejshandler.Config{
		ActionUrl:         "/action",
		EventUrl:          "/events",
		ClientIDHeaderKey: "clientId",
		Handlers:          getHandler(),
	}
}

func main() {
	config := getConfig()
	router := http.NewServeMux()
	goalpinejshandler.Register(router, &config)

	router.Handle("/counter", templ.Handler(template_counter.Index(goalpinejshandler.HeadScripts())))
	router.Handle("/todo", templ.Handler(template_todo.Index(goalpinejshandler.HeadScripts())))

	err := http.ListenAndServe(":40000", router)
	if err != nil {
		log.Fatal(err)
	}
}
