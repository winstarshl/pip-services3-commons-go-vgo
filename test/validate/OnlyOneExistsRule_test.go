package test_validate

import (
	"testing"

	"github.com/winstarshl/pip-services3-commons-go-vgo/validate"
	"github.com/stretchr/testify/assert"
)

func TestOnlyOneExistsRule(t *testing.T) {
	obj := &TestClass{}

	schema := validate.NewSchema().WithRule(validate.NewOnlyOneExistsRule("missingField", "stringField1", "nullField"))
	results := schema.Validate(obj)
	assert.Equal(t, 0, len(results))

	schema = validate.NewSchema().WithRule(validate.NewOnlyOneExistsRule("missingField", "stringField1", "intField"))
	results = schema.Validate(obj)
	assert.Equal(t, 1, len(results))

	schema = validate.NewSchema().WithRule(validate.NewOnlyOneExistsRule("missingField", "nullField"))
	results = schema.Validate(obj)
	assert.Equal(t, 1, len(results))
}
