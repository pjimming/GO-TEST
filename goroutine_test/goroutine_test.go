package goroutine_test

import (
	"fmt"
	"testing"
)

func TestForGoroutine(t *testing.T) {
	for i := 0; i < 100; i++ {
		go func() {
			//time.Sleep(500 * time.Millisecond)
			fmt.Println(i)
		}()
	}
}
