package str_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubStr(t *testing.T) {
	ast := assert.New(t)

	domain := "qnydns.com"
	str := "node.northeast.midsrc.test.qnydns.com."

	ast.Equal("node.northeast.midsrc.test", str[:len(str)-len(domain)-2])
}

func TestSliceStr(t *testing.T) {
	ast := assert.New(t)

	str := "abcdefghhhh"
	res := ""
	page := (len(str) + 2) / 3
	for i := 0; i < page; i++ {
		start := i * 3
		end := (i + 1) * 3
		if end > len(str) {
			end = len(str)
		}

		res += str[start:end]
	}
	ast.Equal(str, res)
}
