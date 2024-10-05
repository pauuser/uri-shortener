package usecase_errors

import (
	"errors"
	"fmt"
)

var (
	NotFoundError     = errors.New("link not found")
	LinkNotFoundError = fmt.Errorf("link not found", NotFoundError)
)
