package Auxiliary_package

import "os"

func Open() string {
	_, err := os.OpenFile("TaskTracker.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return Error{Message: err}.ErrOpenFile()
	}
	return "ok"
}
