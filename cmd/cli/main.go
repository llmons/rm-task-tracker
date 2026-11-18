package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"rm-task-tracker/internal"
	"strconv"
)

func main() {
	switch os.Args[1] {

	// Args[0]  Args[1] Args[2]
	// task-cli add     "description"
	case "add":
		fmt.Printf("Task added successfully (ID: %v)\n", internal.HandleAdd(os.Args[2]))

	// Args[0]  Args[1] Args[2] Args[3]
	// task-cli update  1       "description"
	case "update":
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Printf("strconv.Atoi err:\n%v", err)
		}
		internal.HandleUpdate(id, os.Args[3])

	// Args[0]  Args[1] Args[2]
	// task-cli delete  1
	case "delete":
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Printf("strconv.Atoi err:\n%v", err)
		}
		internal.HandleDelete(id)

	// Args[0]  Args[1]          Args[2]
	// task-cli mark-in-progress 1
	case "mark-in-progress":
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Printf("strconv.Atoi err:\n%v", err)
		}
		internal.HandleMark(internal.InProgress, id)

	// Args[0]  Args[1]   Args[2]
	// task-cli mark-done 1
	case "mark-done":
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Printf("strconv.Atoi err:\n%v", err)
		}
		internal.HandleMark(internal.Done, id)

	// Args[0]  Args[1] Args
	// task-cli list    in-progress
	case "list":
		var tasks []internal.Task
		if len(os.Args) == 2 {
			tasks = internal.HandleListAll()
		} else {
			tasks = internal.HandleListByStatus(internal.Status(os.Args[2]))
		}

		bytes, err := json.MarshalIndent(tasks, "", "\t")
		if err != nil {
			log.Printf("json.MarshalIndent err:\n%v", err)
		}
		fmt.Printf("%s\n", bytes)
	}
}
