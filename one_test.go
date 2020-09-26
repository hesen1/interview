package main

import (
	"testing"
	"sync"
)

func TestPseudoEncrypt(t *testing.T) {
	set := NewSet()
	max := 1000000
	g := sync.WaitGroup{}
	g.Add(max)

	for i := 0; i < max; i++ {
		go func(i int) {
			v := PseudoEncrypt(int64(i))

			if v < 0 {
				t.Fatalf(`小于0`)
				g.Done()
			}

			if set.Has(v) {
				t.Fatalf(`重复`)
				g.Done()
			}

			set.Add(v)
			g.Done()
		}(i)
	}
	g.Wait()
}
