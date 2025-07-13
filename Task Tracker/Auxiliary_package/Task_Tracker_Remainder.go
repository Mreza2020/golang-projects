package Auxiliary_package

import (
	"encoding/json"
	"os"
)

func Remainder() []TaskAddSt {
	file, err := os.ReadFile("./TaskTracker.json")
	if err != nil {
		Error{Message: err}.ErrOpenFile()
	}
	var (
		data          []TaskAddSt
		dataRemainder []TaskAddSt
	)
	err = json.Unmarshal(file, &data)
	if err != nil {
		Error{Message: err}.MarshalIndent()
	}
	for _, value1 := range data {
		if value1.Status != "no" {
			dataRemainder = append(dataRemainder, value1)
		}
	}
	return dataRemainder

}
