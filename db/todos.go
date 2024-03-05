package db

import (
	"embed"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Todos []*Todo

func NewTodos() *Todos {
	return &Todos{}
}

var List = FillTodos()

func (l *Todos) Add(description string) *Todo {
	todo := NewTodo(description, false)
	*l = append(*l, todo)
	return todo

}

func (l *Todos) Remove(id uuid.UUID) {
	index := l.indexOf(id)
	if index == -1 {
		return
	}
	*l = append((*l)[:index], (*l)[index+1:]...)
}

// Update updates a todo in the list
func (l *Todos) Update(id uuid.UUID, completed bool, description string) *Todo {
	index := l.indexOf(id)
	if index == -1 {
		return nil
	}
	todo := (*l)[index]
	todo.Update(completed, description)

	return todo
}

func (l *Todos) Search(search string) []*Todo {
	list := make([]*Todo, 0)
	for _, todo := range *l {
		if strings.Contains(todo.Description, search) {
			list = append(list, todo)
		}
	}
	return list
}

func (l *Todos) All() []*Todo {
	list := make([]*Todo, len(*l))
	copy(list, *l)
	return list
}

func (l *Todos) Get(id uuid.UUID) *Todo {
	index := l.indexOf(id)
	if index == -1 {
		return nil
	}
	return (*l)[index]
}

func (l *Todos) Reorder(ids []uuid.UUID) []*Todo {
	newTodos := make([]*Todo, len(ids))
	for i, id := range ids {
		newTodos[i] = (*l)[l.indexOf(id)]
	}
	copy(*l, newTodos)
	return newTodos
}

func (l *Todos) indexOf(id uuid.UUID) int {
	for i, todo := range *l {
		if todo.ID == id {
			return i
		}
	}
	return -1
}

type JSONTodos struct {
	Todos []*Todo `json:"todos"`
}

//go:embed todos.json
var f embed.FS

// Fill the todo List with json data
func FillTodos() *Todos {
	file, _ := f.ReadFile("todos.json")
	data := &JSONTodos{}
	json.Unmarshal([]byte(file), &data)

	todos := NewTodos()
	for _, todo := range data.Todos {
		todos.AddWithId(todo.Description, todo.Completed, todo.ID)
	}

	return todos
}

func (t *Todos) PrintTodos() {
	for _, todo := range t.All() {
		fmt.Println(todo)
	}
}

func (l *Todos) AddWithId(description string, completed bool, id uuid.UUID) *Todo {
	todo := &Todo{
		Description: description,
		ID:          id,
		Completed:   completed,
		CreatedAt:   time.Now(),
	}
	*l = append(*l, todo)
	return todo
}
