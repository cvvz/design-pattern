package singleton

import (
	"sync"
)

type Singleton struct {
	val int
}

var (
	lazySingle *Singleton
	once       sync.Once
)

func GetLazyInstance() *Singleton {
	once.Do(
		func() {
			lazySingle = new(Singleton)
		},
	)
	return lazySingle
}
