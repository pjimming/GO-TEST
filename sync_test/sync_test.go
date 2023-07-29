package sync_test

import (
	"sync"
	"testing"
)

func TestLockTwice(t *testing.T) {
	var mu sync.Mutex
	mu.Lock()
	mu.Lock()
}
