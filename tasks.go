package task

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	ID      int    `json:"ID"`
	Patient string `json:"patient"`
	Doctor  string `json:"Doctor"`
	Date    string `json:"Date"`
}

func ListTask(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("No hay Tareas")
		return
	}

	for _, task := range tasks {
		fmt.Printf("[%d] %s %s %s\n", task.ID, task.Patient, task.Doctor, task.Date)
	}
}

func AddTask(tasks []Task, Patient string) []Task {
	newTask := Task{
		ID:      GetNextID(tasks),
		Patient: Patient,
		Doctor:  "Juan S",
		Date:    "10/05/2023",
	}
	return append(tasks, newTask)
}

func DeleteTask(tasks []Task, ID int) []Task {
	for i, task := range tasks {
		if task.ID == ID {
			return append(tasks[:i], tasks[i+1:]...)
		}
	}
	return tasks
}

func UpdateTask(tasks []Task, ID int, patient string) []Task {

	for i, t := range tasks {
		if t.ID == ID {
			tasks[i].Patient = patient
		}
	}
	return tasks
}

func SaveTasks(file *os.File, tasks []Task) {
	bytes, err := json.Marshal(tasks)

	if err != nil {
		panic(err)
	}

	_, err = file.Seek(0, 0)

	if err != nil {
		panic(err)
	}

	err = file.Truncate(0)

	if err != nil {
		panic(err)
	}

	writer := bufio.NewWriter(file)
	_, err = writer.Write(bytes)
	if err != nil {
		panic(err)
	}

	err = writer.Flush()
	if err != nil {
		panic(err)
	}
}

func GetNextID(tasks []Task) int {
	if len(tasks) == 0 {
		return 1
	}
	return tasks[len(tasks)-1].ID + 1
}
