package domain

import (
	"errors"
	"os"
)

var (
	ErrUnsupportedType = errors.New("unsupported type")
	ErrTypeNotPtr      = errors.New("type must a pointer")
	ErrNotExist        = os.ErrNotExist
)
