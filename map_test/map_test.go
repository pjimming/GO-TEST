package map_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMapSlice(t *testing.T) {
	ast := assert.New(t)

	adj := make(map[string][]string)
	adj["a"] = append(adj["a"], "111")
	adj["a"] = append(adj["a"], "222")
	ast.NotEmpty(adj["a"])
	for _, item := range adj["a"] {
		fmt.Println(item)
	}
}
