package storage

import "github.com/pkg/errors"

var (
	// For errors it is better to be exportable. It will be handy for client
	ErrNoBookFound = errors.New("requested book doesn't exist")
	ErrMissedGenre = errors.New("bad input, missing values for field Genres")
	ErrMissedPages = errors.New("bad input, missing values for field Pages")
	ErrMissedPrice = errors.New("bad input, missing values for field Price")
	ErrMissedTitle = errors.New("bad input, missing values for field Title")
)

