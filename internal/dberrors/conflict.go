package dberrors

type ConflictErrors struct{}

func (e *ConflictErrors) Error() string {
	return "attempted to create a record with an existing key"
}
