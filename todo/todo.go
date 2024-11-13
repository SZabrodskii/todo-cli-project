package todo

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
	"strconv"
	"time"
)

type Todo struct {
	Title       string `validate:"required"`
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time // nil if not completed - that is why the pointer
}

type Todos []Todo

var validate = validator.New()

func (todos *Todos) Add(title string) error {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CreatedAt:   time.Now(), // TODO: use time.Now() instead of hardcoded time
		CompletedAt: nil,
	}
	if err := validate.Struct(todo); err != nil {
		return errors.New("invalid todo")
	}
	*todos = append(*todos, todo)
	return nil
}

func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("invalid index")
		fmt.Println(err)
		return err
	}

	return nil
}

func (todos *Todos) Delete(index int) error {
	t := *todos
	if err := t.validateIndex(index); err != nil {
		return err
	}

	*todos = append(t[:index], t[index+1:]...)
	return nil
}

func (todos *Todos) Toggle(index int) error {
	t := *todos
	if err := t.validateIndex(index); err != nil {
		return err
	}

	isCompleted := t[index].Completed
	if isCompleted {
		t[index].CompletedAt = nil
	} else {
		completionTime := time.Now()
		t[index].CompletedAt = &completionTime
	}
	t[index].Completed = !isCompleted

	return nil
}

func (todos *Todos) Edit(index int, title string) error {
	t := *todos
	if err := t.validateIndex(index); err != nil {
		return err
	}

	t[index].Title = title

	return nil

}

func (todos *Todos) Print() {
	tw := table.NewWriter()
	tw.SetOutputMirror(os.Stdout)
	tw.AppendHeader(table.Row{"#", "Title", "Completed", "Created At", "Completed At"})

	for index, todo := range *todos {
		completedAt := "N/A"
		if todo.CompletedAt != nil {
			completedAt = todo.CompletedAt.Format(time.RFC1123)
		}
		tw.AppendRow(table.Row{strconv.Itoa(index + 1), todo.Title, todo.Completed, todo.CreatedAt.Format(time.RFC1123), completedAt})
	}
	tw.Render()
}
