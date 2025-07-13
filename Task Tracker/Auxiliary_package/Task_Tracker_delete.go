package Auxiliary_package

import (
	"encoding/json"
	"os"
)

func TaskDelete(nameActivity string) string {
	fileData, err := os.ReadFile("./TaskTracker.json")
	if err != nil {
		Error{Message: err}.ErrOpenFile()
	}

	var (
		data          []map[string]interface{}
		updatedPeople []map[string]interface{}
	)
	err = json.Unmarshal(fileData, &data)
	if err != nil {
		Error{Message: err}.MarshalIndent()
	}

	for _, idMap := range data {
		if name, ok := idMap["id"].(string); ok && name != nameActivity {
			updatedPeople = append(updatedPeople, idMap)
		}
	}

	updatedJSON, err := json.MarshalIndent(updatedPeople, "", "  ")
	if err != nil {
		Error{Message: err}.MarshalIndent()
	}

	err = os.WriteFile("./TaskTracker.json", updatedJSON, 0644)
	if err != nil {
		Error{Message: err}.ErrWriteFile()
	}
	return "ok"

}
