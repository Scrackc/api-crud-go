package model

import "errors"

var (
	// ErrPersonCanNotBeNil la persona no puede ser nula
	ErrPersonCanNotBeNil = errors.New("la persona no puede ser nula")
	// ErrPersonCanNotBeNil la persona no existe
	ErrIdPersonDoesNotExists = errors.New("la persona no existe")
)
