package main

import(
	"bufio"
	"fmt"
	"os" 
	"strconv" 
	"strings"
)

type Todo struct {
	ID int
	Title string
	Completed bool
}
var todos[]Todo
var nextID=1;

func addTodo(reader *bufio.Reader) {
	fmt.Print("Enter Task: ") 
	title,_:=reader.ReadString('\n')
	title=strings.TrimSpace(title);

	if title=="" {
		fmt.Println("task cannot be empty!")
		return
	
}
todo:=Todo{
	ID: nextID,
	Title: title,
	Completed: false,
}
todos=append(todos,todo)
nextID++;

fmt.Println("Todo Added!")
}
func listTodos() {
	if len(todos)==0 {
		fmt.Println("No Todos Found!")
		return
	}
	for _,todo:=range todos {
		status:=" "
		if todo.Completed {
			status="tick"
		}
		fmt.Printf("%d.[%s] %s\n",todo.ID,status,todo.Title)
	}

}

func completeTodo(reader *bufio.Reader) {
	fmt.Print("Enter Todo ID: ")
	input,_:=reader.ReadString('\n')
	input=strings.TrimSpace(input);

	id,err:=strconv.Atoi(input)
	if err!=nil {
		fmt.Println("Invalid ID")
		return
	}
	for i:=range todos {
		if todos[i].ID==id {
			todos[i].Completed=true;
			fmt.Println("Todo completed")
			return
		}
	}
    fmt.Println("Todo Not found!")
}
func deleteTodo(reader *bufio.Reader) {
	fmt.Print("Enter Todo ID: ")
	input,_:=reader.ReadString('\n')
	input=strings.TrimSpace(input)

	id,err:=strconv.Atoi(input)
	if err!=nil {
		fmt.Println("Invalid ID")
		return
	}
    for i:=range todos{
		if todos[i].ID==id {
			todos=append(todos[:i],todos[i+1:]...)
			fmt.Println("Todo deleted!")
			return
		}
	}
	fmt.Println("Todo Not ")
}
func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n===== TODO MANAGER =====")
		fmt.Println("1. Add Todo")
		fmt.Println("2. List Todos")
		fmt.Println("3. Complete Todo")
		fmt.Println("4. Delete Todo")
		fmt.Println("5. Exit")

		fmt.Print("Choose: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid choice!")
			continue
		}

		switch choice {
		case 1:
			addTodo(reader)

		case 2:
			listTodos()

		case 3:
			completeTodo(reader)

		case 4:
			deleteTodo(reader)

		case 5:
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid choice!")
		}
	}
}
