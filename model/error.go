package model

type ErrNotFound struct {}
type ErrCannotConvType struct {}


func (e *ErrNotFound) Error() string {
	return "Error: NotFound"
}

func (e *ErrCannotConvType) Error() string {
	return "Error: CannotConvType"
}