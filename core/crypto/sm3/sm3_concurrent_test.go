package sm3

import (
	"sync"
	"testing"
)

func TestSm3ConcurrentSum(t *testing.T) {
	data := []byte("concurrent sm3 test")
	expected := Sm3Sum(data)
	var wg sync.WaitGroup
	concurrency := 100

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			h := Sm3Sum(data)
			if len(h) != 32 {
				t.Error("Sm3Sum returned wrong length")
			}
		}()
	}
	wg.Wait()
	_ = expected
}

func TestSm3ConcurrentIncremental(t *testing.T) {
	var wg sync.WaitGroup
	concurrency := 50

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			h := New()
			h.Write([]byte("part1"))
			h.Write([]byte("part2"))
			result := h.Sum(nil)
			if len(result) != 32 {
				t.Error("Sum returned wrong length")
			}
		}()
	}
	wg.Wait()
}
