package main

import (
	"testing"
	"sync"
)

func TestSetAdd(t *testing.T) {
	set := NewSet()
	testValue := 1

	set.Add(testValue)

	if !set.Has(testValue) {
		t.Fatal("Set Add: add方法无效")
	}

	if set.Len() != 1 {
		t.Fatal("Set Add: add方法无效")
	}

	set.Clear()

	if set.Has(testValue) {
		t.Fatal("Set Clear: Clear方法无效")
	}
}

func TestSetLen(t *testing.T) {
	var counts int = 1000
	set := NewSet()
	g := sync.WaitGroup{}
	var i int

	for i = 0; i < counts; i++ {
		g.Add(1)
		go func (j int) {
			set.Add(j)
			g.Done()
		}(i)
	}
	g.Wait()

	if set.Len() != counts {
		t.Fatalf("期望: %d 条数据, 实际: %d ", counts, set.Len())
	}

	for i = 0; i < counts; i++ {
		g.Add(1)
		go func (j int) {
			if !set.Has(j) {
				set.Add(j)
			}
			g.Done()
		}(i)
	}
	g.Wait()

	if set.Len() != counts {
		t.Fatalf("期望: %d 条数据, 实际: %d ", counts, set.Len())
	}
}
