package errors

import "errors"

var (
	ErrGameNotInitialized   = errors.New("game not initialized")
	ErrGameAlreadyDestroyed = errors.New("game already destroyed")
)
