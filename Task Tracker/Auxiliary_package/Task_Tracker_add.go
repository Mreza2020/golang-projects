package Auxiliary_package

import (
	"encoding/json"
	"os"
)

type TaskAddSt struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

func TaskAdd(id string, description string, status string, CreatedAt string) string {
	Tr := TaskAddSt{
		Id:          id,
		Description: description,
		Status:      status,
		CreatedAt:   CreatedAt,
		UpdatedAt:   "",
	}

	fileData, err := os.ReadFile("./TaskTracker.json")
	if err != nil {
		Error{Message: err}.ErrOpenFile()
	}
	var tasks []TaskAddSt

	if len(fileData) > 0 {
		err = json.Unmarshal(fileData, &tasks)
		if err != nil {
			Error{Message: err}.MarshalIndent()
		}
	}
	tasks = append(tasks, Tr)

	newData, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		Error{Message: err}.MarshalIndent()
	}

	err = os.WriteFile("TaskTracker.json", newData, 0644)
	if err != nil {
		return Error{Message: err}.ErrWriteFile()
	}

	return "ok"
}
