package main

import (
	"github.com/nodejayes/go-alpinejs-handler-poc/cosmic_ui/cosmic_ui_global"
	"github.com/nodejayes/go-alpinejs-handler-poc/counter_app"
	"github.com/nodejayes/go-alpinejs-handler-poc/todo_app"
	"log"
	"net/http"

	"github.com/nodejayes/go-alpinejs-handler"
)

func getPages() []goalpinejshandler.Page {
	return []goalpinejshandler.Page{
		counter_app.NewPage(),
		todo_app.NewPage(),
	}
}

func getConfig() goalpinejshandler.Config {
	return goalpinejshandler.Config{
		ActionUrl:         "/action",
		EventUrl:          "/events",
		ClientIDHeaderKey: "clientId",
		Pages:             getPages(),
	}
}

func main() {
	config := getConfig()
	router := http.NewServeMux()
	goalpinejshandler.Register(router, &config)
	cosmic_ui_global.RegisterGlobalStyles()
	err := http.ListenAndServe(":40000", router)
	if err != nil {
		log.Fatal(err)
	}
}
