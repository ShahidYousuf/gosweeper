package pkg

import "github.com/google/uuid"

type UIDGen interface {
	New() string
}

type uidgen struct{}

func (u uidgen) New() string {
	return uuid.New().String()
}

func New() UIDGen {
	return &uidgen{}
}
