package validate

import (
	"strings"

	"github.com/winstarshl/pip-services3-commons-go-vgo/reflect"
)

type AtLeastOneExistsRule struct {
	properties []string
}

func NewAtLeastOneExistsRule(properties ...string) *AtLeastOneExistsRule {
	return &AtLeastOneExistsRule{
		properties: properties,
	}
}

func (c *AtLeastOneExistsRule) Validate(path string, schema ISchema, value interface{}) []*ValidationResult {
	name := path
	if name == "" {
		name = "value"
	}
	found := 0

	for _, property := range c.properties {
		propertyValue := reflect.ObjectReader.GetProperty(value, property)
		if propertyValue != nil {
			found++
		}
	}

	if found == 0 {
		return []*ValidationResult{
			NewValidationResult(
				path,
				Error,
				"VALUE_NULL",
				name+" must have at least one property from "+strings.Join(c.properties, ","),
				c.properties,
				nil,
			),
		}
	}

	return nil
}
