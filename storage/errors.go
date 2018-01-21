package storage

import "github.com/pkg/errors"

var (
	// For errors it is better to be exportable. It will be handy for client
	ErrNoBookFound = errors.New("requested book doesn't exist")
)
