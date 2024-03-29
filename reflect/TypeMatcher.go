package reflect

import (
	refl "reflect"
	"strings"
	"time"

	"github.com/winstarshl/pip-services3-commons-go-vgo/convert"
)

type TTypeMatcher struct{}

var TypeMatcher = &TTypeMatcher{}

func (c *TTypeMatcher) MatchValue(expectedType interface{}, actualValue interface{}) bool {
	if expectedType == nil {
		return true
	}
	if actualValue == nil {
		panic("Actual value cannot be nil")
	}

	// Check actual value by its type
	return c.MatchType(expectedType, refl.TypeOf(actualValue))
}

func (c *TTypeMatcher) MatchValueByName(expectedType string, actualValue interface{}) bool {
	if expectedType == "" {
		return true
	}
	if actualValue == nil {
		panic("Actual value cannot be nil")
	}

	// Check actual value by its type
	return c.MatchTypeByName(expectedType, refl.TypeOf(actualValue))
}

func (c *TTypeMatcher) MatchType(expectedType interface{}, actualType refl.Type) bool {
	if expectedType == nil {
		return true
	}
	if actualType == nil {
		panic("Actual type cannot be null")
	}

	// Compare for matching types
	if expectedType == actualType {
		return true
	}

	// Extract inner value because Go implementations of Maps and Arrays are wrappers
	innerType, ok := expectedType.(IValueWrapper)
	if ok {
		expectedType = innerType.InnerValue()
	}

	// If expected value is type
	typ, ok1 := expectedType.(refl.Type)
	if ok1 {
		type2 := typ
		// Check pointer type as well
		if type2.Kind() == refl.Ptr {
			type2 = type2.Elem()
		}
		return actualType.AssignableTo(typ) || actualType.AssignableTo(type2)
	}

	// For strings compare string types
	str, ok2 := expectedType.(string)
	if ok2 {
		return c.MatchTypeByName(str, actualType)
	}

	// For typecodes compare them
	typeCode, ok3 := expectedType.(convert.TypeCode)
	if ok3 {
		return convert.TypeConverter.ToTypeCode(actualType) == typeCode
	}

	return false
}

func (c *TTypeMatcher) MatchTypeByName(expectedType string, actualType refl.Type) bool {
	if expectedType == "" {
		return true
	}

	if actualType == nil {
		panic("Actual type cannot be null")
	}

	if actualType.Kind() == refl.Ptr {
		actualType = actualType.Elem()
	}

	expectedType = strings.ToLower(expectedType)
	actualTypeName := strings.ToLower(actualType.Name())
	actualTypeFullName := strings.ToLower(actualType.PkgPath() + "." + actualType.Name())
	actualTypeKind := actualType.Kind()

	if actualTypeName == expectedType || actualTypeFullName == expectedType {
		return true
	}

	if expectedType == "object" {
		return true
	}

	if expectedType == "int" || expectedType == "integer" {
		return actualTypeKind == refl.Int8 ||
			actualTypeKind == refl.Uint8 ||
			actualTypeKind == refl.Int16 ||
			actualTypeKind == refl.Uint16 ||
			actualTypeKind == refl.Int32 ||
			actualTypeKind == refl.Int
	}

	if expectedType == "long" {
		return actualTypeKind == refl.Int8 ||
			actualTypeKind == refl.Uint8 ||
			actualTypeKind == refl.Int16 ||
			actualTypeKind == refl.Uint16 ||
			actualTypeKind == refl.Int32 ||
			actualTypeKind == refl.Uint32 ||
			actualTypeKind == refl.Int64 ||
			actualTypeKind == refl.Uint64 ||
			actualTypeKind == refl.Int ||
			actualTypeKind == refl.Uint
	}

	if expectedType == "float" {
		return actualTypeKind == refl.Float32
	}

	if expectedType == "double" {
		return actualTypeKind == refl.Float32 ||
			actualTypeKind == refl.Float64
	}

	if expectedType == "string" {
		return actualTypeKind == refl.String
	}

	if expectedType == "bool" || expectedType == "boolean" {
		return actualTypeKind == refl.Bool
	}

	if expectedType == "date" || expectedType == "datetime" {
		return actualType == refl.TypeOf(time.Time{})
	}

	if expectedType == "timespan" || expectedType == "duration" {
		return actualTypeKind == refl.Int8 ||
			actualTypeKind == refl.Uint8 ||
			actualTypeKind == refl.Int16 ||
			actualTypeKind == refl.Uint16 ||
			actualTypeKind == refl.Int32 ||
			actualTypeKind == refl.Uint32 ||
			actualTypeKind == refl.Int64 ||
			actualTypeKind == refl.Uint64 ||
			actualTypeKind == refl.Int ||
			actualTypeKind == refl.Uint ||
			actualTypeKind == refl.Float32 ||
			actualTypeKind == refl.Float64 ||
			actualType == refl.TypeOf(time.Duration(1))
	}

	if expectedType == "map" || expectedType == "dict" || expectedType == "dictionary" {
		return actualTypeKind == refl.Map
	}

	if expectedType == "array" || expectedType == "list" {
		return actualTypeKind == refl.Array ||
			actualTypeKind == refl.Slice
	}

	if strings.HasSuffix(expectedType, "[]") {
		if actualTypeKind == refl.Slice || actualTypeKind == refl.Array {
			expectedType = expectedType[:len(expectedType)-2]
			actualType = actualType.Elem()
			return c.MatchTypeByName(expectedType, actualType)
		}
	}

	return false
}
