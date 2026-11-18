package internal

import (
	"log"
	"slices"
	"sort"
	"time"
)

func HandleAdd(description string) int {
	fp, tasks, err := ReadTasksWithFp()
	if err != nil {
		log.Printf("ReadTasksWithFp err:\n%v", err)
		return -1
	}
	defer func() {
		if err := fp.Close(); err != nil {
			log.Printf("fp.Close err:\n%v", err)
		}
	}()

	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i].Id > tasks[j].Id
	})

	newId := 1
	if len(tasks) > 0 {
		newId = tasks[0].Id
	}
	newTask := Task{
		Id:          newId,
		Description: description,
		Status:      Status("todo"),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	tasks = append(tasks, newTask)

	err = WriteTasksWithFp(fp, tasks)
	if err != nil {
		log.Printf("WriteTasksWithFp err:\n%v", err)
	}

	return newId
}

func HandleUpdate(id int, newTaskDescription string) {
	fp, tasks, err := ReadTasksWithFp()
	if err != nil {
		log.Printf("ReadTasksWithFp err:\n%v", err)
	}
	defer func() {
		if err := fp.Close(); err != nil {
			log.Printf("fp.Close err:\n%v", err)
		}
	}()

	idx := slices.IndexFunc(tasks, func(task Task) bool {
		return task.Id == id
	})
	if idx == -1 {
		return
	}

	tasks[idx].Description = newTaskDescription
	tasks[idx].UpdatedAt = time.Now()

	err = WriteTasksWithFp(fp, tasks)
	if err != nil {
		log.Printf("WriteTasksWithFp err:\n%v", err)
	}
}

func HandleDelete(id int) {
	fp, tasks, err := ReadTasksWithFp()
	if err != nil {
		log.Printf("ReadTasksWithFp err:\n%v", err)
	}
	defer func() {
		if err := fp.Close(); err != nil {
			log.Printf("fp.Close err:\n%v", err)
		}
	}()

	idx := slices.IndexFunc(tasks, func(task Task) bool {
		return task.Id == id
	})
	if idx == -1 {
		return
	}

	tasks = slices.Delete(tasks, idx, idx)

	err = WriteTasksWithFp(fp, tasks)
	if err != nil {
		log.Printf("WriteTasksWithFp err:\n%v", err)
	}
}

func HandleMark(status Status, id int) {
	fp, tasks, err := ReadTasksWithFp()
	if err != nil {
		log.Printf("ReadTasksWithFp err:\n%v", err)
	}
	defer func() {
		if err := fp.Close(); err != nil {
			log.Printf("fp.Close err:\n%v", err)
		}
	}()

	idx := slices.IndexFunc(tasks, func(task Task) bool {
		return task.Id == id
	})
	if idx == -1 {
		return
	}

	tasks[idx].Status = status
	tasks[idx].UpdatedAt = time.Now()

	err = WriteTasksWithFp(fp, tasks)
	if err != nil {
		log.Printf("WriteTasksWithFp err:\n%v", err)
	}
}

func HandleListAll() []Task {
	tasks, err := ReadTasks()
	if err != nil {
		log.Printf("ReadTasks err:\n%v", err)
	}

	return tasks
}

func HandleListByStatus(status Status) []Task {
	tasks, err := ReadTasks()
	if err != nil {
		log.Printf("ReadTasks err:\n%v", err)
	}

	return slices.DeleteFunc(tasks, func(t Task) bool {
		return t.Status != status
	})
}
