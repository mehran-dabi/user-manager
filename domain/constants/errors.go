package constants

import "fmt"

var (
	ErrUserNotFound = fmt.Errorf("user not found")
	ErrHasNoChanges = fmt.Errorf("the information has no changes")
)
