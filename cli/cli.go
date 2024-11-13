package cli

import (
	"errors"
	"flag"
	"fmt"
	"strconv"
	"strings"
	"todo-cli-go/todo"
)

type CmdFlags struct {
	Add    string
	Delete int
	Toggle int
	Edit   string
	List   bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add a new todo specified by the title")
	flag.IntVar(&cf.Delete, "delete", -1, "Delete the todo at the specified index")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Toggle the todo at the specified index")
	flag.StringVar(&cf.Edit, "edit", "", "Edit the todo by index and specify the new title. id:new_title")
	flag.BoolVar(&cf.List, "list", false, "List all todos")

	flag.Parse()

	return &cf

}

func (cf *CmdFlags) Execute(todos *todo.Todos) error {
	switch {
	case cf.List:
		todos.Print()
	case cf.Add != "":
		todos.Add(cf.Add)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			return errors.New("invalid edit command. Please use id:new_title")
			//os.Exit(1)
		}
		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("invalid index for editing")
			return err
			//os.Exit(1)
		}
		todos.Edit(index, parts[1])
	case cf.Delete != -1:
		todos.Delete(cf.Delete)
	case cf.Toggle != -1:
		todos.Toggle(cf.Toggle)
	default:
		return errors.New("invalid command")
		//os.Exit(1)
	}

	return nil
}
