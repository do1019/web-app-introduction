package model

type Error struct {
	errmsg string
}

func (e *Error) ErrNotFound() string {
	return "Error: NotFound"
}

// func (e *Error) ErrCannotConvType() string {
// 	return "Error: CannotConvType"
// }