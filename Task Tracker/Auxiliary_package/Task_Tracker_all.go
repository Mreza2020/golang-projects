package Auxiliary_package

import (
	"encoding/json"
	"os"
)

func AllTask() []map[string]interface{} {
	fileData, err := os.ReadFile("./TaskTracker.json")
	if err != nil {
		Error{Message: err}.ErrOpenFile()
	}
	var (
		data []map[string]interface{}
	)
	err = json.Unmarshal(fileData, &data)
	if err != nil {
		Error{Message: err}.MarshalIndent()
	}
	return data

}
