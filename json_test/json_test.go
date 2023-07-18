package json_test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrListToObject(t *testing.T) {
	ast := assert.New(t)

	array := []string{"apple", "banana", "cherry"}

	// 将字符串数组转换为JSON字符串
	jsonString, err := json.Marshal(array)
	ast.Nil(err)

	// 输出JSON字符串
	ast.Equal(string(jsonString), `["apple","banana","cherry"]`)
}

func TestObjectToStrList(t *testing.T) {
	ast := assert.New(t)

	jsonStr := `["apple","banana","cherry"]`

	res := make([]string, 0)
	err := json.Unmarshal([]byte(jsonStr), &res)
	ast.Nil(err)
	ast.Equal(res, []string{"apple", "banana", "cherry"})
}
