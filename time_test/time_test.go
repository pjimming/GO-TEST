package time_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestParseDateTime(t *testing.T) {
	ast := assert.New(t)
	str := "2/3/23 10:00"
	_, err := time.Parse("1/2/06 15:04", str)
	ast.Nil(err)
}
