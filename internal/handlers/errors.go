package handlers

import "errors"

var (
	ErrInvalidPathParameter = errors.New("path parameter should be integer number")
)
