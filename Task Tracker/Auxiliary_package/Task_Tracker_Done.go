package Auxiliary_package

import (
	"encoding/json"
	"os"
)

func Done() []TaskAddSt {
	file, err := os.ReadFile("./TaskTracker.json")
	if err != nil {
		Error{Message: err}.ErrOpenFile()
	}
	var (
		data     []TaskAddSt
		dataDone []TaskAddSt
	)
	err = json.Unmarshal(file, &data)
	if err != nil {
		Error{Message: err}.MarshalIndent()
	}
	for _, v := range data {
		if v.Status == "no" {
			dataDone = append(dataDone, v)
		}
	}
	return dataDone
}
