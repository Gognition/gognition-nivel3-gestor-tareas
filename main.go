package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Task struct {
	ID          int
	Title       string
	Description string
	Done        bool
	DueDate     time.Time
}

func (t *Task) MarkAsDone() {
	t.Done = true
}

func AddTask(tasks []Task, newTask Task) []Task {
	return append(tasks, newTask)
}

func ListTasks(tasks []Task) {
	for _, task := range tasks {
		status := "Pendiente"
		if task.Done {
			status = "Completada"
		}
		fmt.Printf("[%d] %s - %s (Vence: %s) - %s\n",
			task.ID, task.Title, task.Description,
			task.DueDate.Format("02/01/2006"), status)
	}
}

func SaveTasks(tasks []Task) error {
	file, err := os.Create("tasks.json")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(tasks)
}

func LoadTasks() ([]Task, error) {
	file, err := os.Open("tasks.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var tasks []Task
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)
	return tasks, err
}

func main() {
	tasks, err := LoadTasks()
	if err != nil {
		fmt.Println("No se pudieron cargar las tareas:", err)
		tasks = []Task{}
	}

	for {
		fmt.Println("\n1. Añadir tarea")
		fmt.Println("2. Listar tareas")
		fmt.Println("3. Marcar tarea como completada")
		fmt.Println("4. Salir")
		fmt.Print("Elige una opción: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var title, desc string
			fmt.Print("Título de la tarea: ")
			fmt.Scanln(&title)
			fmt.Print("Descripción: ")
			fmt.Scanln(&desc)
			newTask := Task{
				ID:          len(tasks) + 1,
				Title:       title,
				Description: desc,
				Done:        false,
				DueDate:     time.Now().AddDate(0, 0, 7), // Vence en 7 días
			}
			tasks = AddTask(tasks, newTask)
		case 2:
			ListTasks(tasks)
		case 3:
			var id int
			fmt.Print("ID de la tarea completada: ")
			fmt.Scanln(&id)
			for i := range tasks {
				if tasks[i].ID == id {
					tasks[i].MarkAsDone()
					break
				}
			}
		case 4:
			err := SaveTasks(tasks)
			if err != nil {
				fmt.Println("Error al guardar las tareas:", err)
			}
			fmt.Println("¡Hasta luego!")
			return
		default:
			fmt.Println("Opción no válida")
		}
	}
}
