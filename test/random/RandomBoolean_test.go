package test_random

import (
	"testing"

	"github.com/winstarshl/pip-services3-commons-go-vgo/random"
	"github.com/stretchr/testify/assert"
)

func TestChance(t *testing.T) {
	value := random.RandomBoolean.Chance(5, 10)
	assert.True(t, value || !value)

	value = random.RandomBoolean.Chance(5, 5)
	assert.True(t, value || !value)

	value = random.RandomBoolean.Chance(0, 0)
	assert.True(t, !value)

	value = random.RandomBoolean.Chance(-1, 0)
	assert.True(t, !value)

	value = random.RandomBoolean.Chance(-1, -1)
	assert.True(t, !value)
}
