package singleton

import (
	"sync"
	"testing"
)

func TestGetInstance(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(200)
	lazyCh := make(chan struct{})
	hungryCh := make(chan struct{})

	for i := 0; i < 100; i++ {
		go func() {
			<-lazyCh
			t.Logf("lazy: %p\n", GetLazyInstance())
			wg.Done()
			hungryCh <- struct{}{}
		}()

		go func() {
			<-hungryCh
			t.Logf("hungry: %p\n", GetInstance())
			wg.Done()
			lazyCh <- struct{}{}
		}()
	}

	lazyCh <- struct{}{}
	<-lazyCh

	wg.Wait()
}
