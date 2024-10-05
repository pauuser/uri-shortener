package router_errors

import (
	"errors"
	"fmt"
)

var (
	BadRequest           = errors.New("bad request")
	BadRequestNoLink     = fmt.Errorf("%w: no link", BadRequest)
	BadRequestInvalidTtl = fmt.Errorf("%w: invalid ttl", BadRequest)

	NotFound       = errors.New("not found")
	NotFoundNoLink = fmt.Errorf("%w: no link", NotFound)

	InternalServerError = errors.New("internal server error")
)
