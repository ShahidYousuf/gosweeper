package pkg

import "errors"

var (
	NotFound         = errors.New("not_found")
	IllegalOperation = errors.New("illegal_operation")
	InvalidInput     = errors.New("invalid_input")
	Internal         = errors.New("internal")
)
