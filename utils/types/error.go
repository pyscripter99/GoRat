package types

import "fmt"

type Error struct {
	Namespace string
	Err       error
}

func (e *Error) String() string {
	return fmt.Sprintf("%s error: %s", e.Namespace, e.Err.Error())
}
