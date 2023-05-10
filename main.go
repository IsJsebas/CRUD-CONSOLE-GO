package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	task "mongo-golang/tasks"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.OpenFile("tasks.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	var tasks []task.Task

	info, err := file.Stat()
	if err != nil {
		panic(err)
	}

	if info.Size() != 0 {
		bytes, err := io.ReadAll(file)
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(bytes, &tasks)
		if err != nil {
			panic(err)
		}
	} else {
		tasks = []task.Task{}
	}

	if len(os.Args) < 2 {
		printUsage()
	}

	switch os.Args[1] {
	case "read":
		task.ListTask(tasks)

	case "create":
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Nombre del paciente que desea agregar.")
		Patient, _ := reader.ReadString('\n')
		Patient = strings.TrimSpace(Patient)

		tasks = task.AddTask(tasks, Patient)
		task.SaveTasks(file, tasks)

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Debes indicar un ID válido para eliminar.")
			return
		}
		ID, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("El ID debe ser un número.")
			return
		}

		tasks = task.DeleteTask(tasks, ID)
		task.SaveTasks(file, tasks)

	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Debes indicar un ID y un nombre válido para actualizar.")
			return
		}
		var patient = os.Args[3]

		ID, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("El ID debe ser un número.")
			return
		}

		tasks = task.UpdateTask(tasks, ID, patient)
		fmt.Println(tasks)
		task.SaveTasks(file, tasks)

	}
}
func printUsage() {
	fmt.Println("Tienes estas opciones: go-clid-crud [create|read|update|delete]")
}
