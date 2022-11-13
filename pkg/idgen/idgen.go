package idgen

import "github.com/google/uuid"

type generator struct{}

type Generator interface {
	NewUUID() string
}

func New() *generator {
	return &generator{}
}

func (*generator) NewUUID() string {
	return uuid.NewString()
}
