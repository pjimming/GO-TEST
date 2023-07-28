package interface_test

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

var tmpMap = map[string]interface{}{
	"float":  1.1,
	"int":    80,
	"string": "abc",
	"bool":   true,
}

func TestFloatToInt(t *testing.T) {
	ast := assert.New(t)
	if v, ok := tmpMap["float"].(int); ok {
		ast.Equal(v, 1)
	} else {
		log.Println("can not trans", v)
	}
}
