package model

type ErrNotFound struct {}

func (e *ErrNotFound) Error() string {
	return "Error: NotFound"
}

type ErrCannotConvType struct {}

func (e *ErrCannotConvType) Error() string {
	return "Error: CannotConvType"
}