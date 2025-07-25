package Receiver

import (
	auxiliarypackage "Number_Guessing_Game/Auxiliary_package/Error"
	"fmt"
)

func Receiver(m *string) string {
	_, err := fmt.Scanln(m)
	return auxiliarypackage.Error{Message: err}.ErrScan()
}
