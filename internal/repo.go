package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func ReadTasks() ([]Task, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("os.Getwd err:\n%w", err)
	}

	path := filepath.Join(dir, JsonStoreName)
	_, err = os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		return nil, nil
	}

	// read file
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("os.ReadFile err:\n%w", err)
	}

	// unmarshal from json
	var tasks []Task
	if err := json.Unmarshal(bytes, &tasks); err != nil {
		return nil, fmt.Errorf("json.Unmarshal err:\n%w", err)
	}

	return tasks, nil
}

func ReadTasksWithFp() (*os.File, []Task, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, nil, fmt.Errorf("os.Getwd err:\n%w", err)
	}

	path := filepath.Join(dir, JsonStoreName)
	_, err = os.Stat(path)

	// create file, no defer since return *os.File
	if errors.Is(err, os.ErrNotExist) {
		fp, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0o666)
		if err != nil {
			return fp, nil, fmt.Errorf("create file err:\n%w", err)
		}
		return fp, nil, nil
	}

	// open file, no defer since return *os.File
	fp, err := os.OpenFile(path, os.O_RDWR, 0o666)
	if err != nil {
		return fp, nil, fmt.Errorf("open file err:\n%w", err)
	}

	// read file
	bytes, err := io.ReadAll(fp)
	if err != nil {
		return fp, nil, fmt.Errorf("io.ReadAll err:\n%w", err)
	}

	// unmarshal from json
	var tasks []Task
	if err := json.Unmarshal(bytes, &tasks); err != nil {
		return fp, nil, fmt.Errorf("json.Unmarshal err:\n%w", err)
	}

	return fp, tasks, nil
}

func WriteTasksWithFp(fp *os.File, tasks []Task) error {
	bytes, err := json.Marshal(tasks)
	if err != nil {
		return fmt.Errorf("json.Marshal err:\n%w", err)
	}

	// truncate file
	err = fp.Truncate(0)
	if err != nil {
		return fmt.Errorf("fp.Truncate err:\n%w", err)
	}

	// seek cusor
	_, err = fp.Seek(0, 0)
	if err != nil {
		return fmt.Errorf("fp.Seek err:\n%w", err)
	}

	// write to file
	_, err = fp.Write(bytes)
	if err != nil {
		return fmt.Errorf("fp.Write err:\n%w", err)
	}

	return nil
}
