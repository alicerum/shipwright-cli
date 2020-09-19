package util

import "errors"

// APIVerb verb name, to be called against any given resource.
type APIVerb string

const (
	Create APIVerb = "create"
	Update APIVerb = "update"
	Delete APIVerb = "delete"
)

// ErrUnknownVerb for unknown verbs, not exported as a constant.
var ErrUnknownVerb = errors.New("unknown API verb")
