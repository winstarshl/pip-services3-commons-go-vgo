package test_convert

import (
	"testing"

	"github.com/winstarshl/pip-services3-commons-go-vgo/convert"
	"github.com/stretchr/testify/assert"
)

func TestToDouble(t *testing.T) {
	assert.Nil(t, convert.ToNullableDouble(nil))

	assert.Equal(t, 123., convert.ToDouble(123))
	assert.Equal(t, 123.456, convert.ToDouble(123.456))
	assert.Equal(t, 123., convert.ToDouble("123"))
	assert.Equal(t, 123.456, convert.ToDouble("123.456"))

	assert.Equal(t, 123., convert.ToDoubleWithDefault(nil, 123))
	assert.Equal(t, 0., convert.ToDoubleWithDefault(false, 123))
	assert.Equal(t, 123., convert.ToDoubleWithDefault("ABC", 123))
}
