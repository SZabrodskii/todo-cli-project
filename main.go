package main

import (
	"todo-cli-go/cli"
	"todo-cli-go/storage"
	"todo-cli-go/todo"
)

func main() {
	todos := todo.Todos{}
	storage := storage.NewStorage[todo.Todos]("todos.json")
	defer storage.Close()

	storage.Load(&todos)
	cmdFlags := cli.NewCmdFlags()
	cmdFlags.Execute(&todos)
	storage.Save(todos)
}
