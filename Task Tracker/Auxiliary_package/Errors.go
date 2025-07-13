package Auxiliary_package

import (
	"fmt"
)

type Error struct {
	Message error
}

type Errors interface {
	ErrOpenFile() string
	ErrWriteFile() string
	ErrScan() string
}

// ErrOpenFile package for handling errors when opening a file
func (e Error) ErrOpenFile() string {
	return fmt.Sprintf("\n Error opening file: %s\n", e)
}

func (e Error) ErrWriteFile() string {
	return fmt.Sprintf("\n Error WriteFile file: %s\n", e)
}

func (e Error) ErrScan() string {
	if e.Message != nil {
		return fmt.Sprintf("\n Error Invalid input: %s\n", e)
	} else {
		return ""
	}

}

func (e Error) MarshalIndent() string {
	return fmt.Sprintf("\n Error MarshalIndent : %s\n", e)

}
