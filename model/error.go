package model

type Error struct {}

func (e *Error) ErrNotFound() string {
	return "Error: NotFound"
}

func (e *Error) ErrCannotConvType() string {
	return "Error: CannotConvType"
}