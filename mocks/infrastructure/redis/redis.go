package mocks

import (
	"github.com/alicebob/miniredis/v2"
)

func NewRedisMock() *miniredis.Miniredis {
	s, err := miniredis.Run()

	if err != nil {
		panic(err)
	}

	return s
}
