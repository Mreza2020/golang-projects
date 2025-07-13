package Auxiliary_package

import (
	"encoding/json"
	"os"
	"time"
)

func TaskUpdate(nameActivity string, status string) string {
	fileData, err := os.ReadFile("./TaskTracker.json")
	if err != nil {
		Error{Message: err}.ErrOpenFile()

	}

	var data []map[string]interface{}

	err = json.Unmarshal(fileData, &data)
	if err != nil {
		Error{Message: err}.MarshalIndent()
	}

	update := ""
	for i, idMap := range data {
		if activity, ok := idMap["id"].(string); ok && activity == nameActivity {
			data[i]["status"] = status
			now := time.Now()
			data[i]["updatedAt"] = now.Format("2006-01-02 15:04:05")
			update = "ok Update"
			break
		}
	}

	updatedJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		Error{Message: err}.MarshalIndent()
	}

	err = os.WriteFile("./TaskTracker.json", updatedJSON, 0644)
	if err != nil {
		Error{Message: err}.ErrWriteFile()
	}

	return update

}
