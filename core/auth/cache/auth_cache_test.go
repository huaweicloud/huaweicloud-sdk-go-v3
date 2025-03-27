package cache

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestPutAndGet(t *testing.T) {
	c := GetCache()

	t.Run("basic read", func(t *testing.T) {
		key := "test_key"
		value := "test_value"

		c.PutAuth(key, value)
		retrieved, ok := c.GetAuth(key)

		assert.True(t, ok, "The existing keys should be found")
		assert.Equal(t, value, retrieved, "The value read should be consistent with the value written")
	})

	t.Run("Read non-existent key", func(t *testing.T) {
		_, ok := c.GetAuth("non_existent_key")
		assert.False(t, ok, "Non-existent keys should return false")
	})
}

func TestConcurrentAccess(t *testing.T) {
	c := GetCache()
	key := "concurrent_key"
	value := "race_test_value"

	var wg sync.WaitGroup
	wg.Add(2)

	// Concurrent write test
	t.Run("Concurrent write safety", func(t *testing.T) {
		go func() {
			defer wg.Done()
			c.PutAuth(key, value+"1")
		}()

		go func() {
			defer wg.Done()
			c.PutAuth(key, value+"2")
		}()

		wg.Wait()

		// The final value should be the value written last
		result, _ := c.GetAuth(key)
		assert.Contains(t, []string{value + "1", value + "2"}, result, "The result should be one of the two")
	})

	// Mixed reading and writing test
	t.Run("Mixed reading and writing", func(t *testing.T) {
		const workers = 100
		wg.Add(workers * 2)

		for i := 0; i < workers; i++ {
			go func(n int) {
				defer wg.Done()
				c.PutAuth(key, value+string(rune(n)))
			}(i)

			go func(n int) {
				defer wg.Done()
				_, _ = c.GetAuth(key)
			}(i)
		}

		wg.Wait()
		assert.NotPanics(t, func() { c.GetAuth(key) }, "Concurrent read and write should not cause panic")
	})
}
