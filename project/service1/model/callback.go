package model

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// CallbackRequest is a JSON request to /callback route.
type CallbackRequest struct {
	ObjectIDs []uint `json:"object_ids" validate:"required,notempty,comma,max=200"`
}

func validateStruct(s interface{}) error {
	val := reflect.ValueOf(s)
	if val.Kind() != reflect.Struct {
		return errors.New("validateStruct: input is not a struct")
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := val.Type().Field(i)
		tag := fieldType.Tag.Get("validate")

		if tag == "" {
			continue
		}

		tagRules := strings.Split(tag, ",")
		var maxCount int
		var hasMaxCount bool

		for _, rule := range tagRules {
			if strings.HasPrefix(rule, "max=") {
				max, err := strconv.Atoi(strings.TrimPrefix(rule, "max="))
				if err != nil {
					return fmt.Errorf("invalid max value: %v", err)
				}
				maxCount = max
				hasMaxCount = true
			}
		}

		for _, rule := range tagRules {
			switch rule {
			case "required":
				if isEmptyValue(field) {
					return fmt.Errorf("field '%s' is required", fieldType.Name)
				}
			case "notempty":
				if field.Kind() == reflect.String && len(strings.TrimSpace(field.String())) == 0 {
					return fmt.Errorf("field '%s' cannot be empty", fieldType.Name)
				}
			case "comma":
				if err := validateCommaSeparatedInts(field.String()); err != nil {
					return fmt.Errorf("field '%s' %v", fieldType.Name, err)
				}
			}
		}
		if hasMaxCount {
			if err := validateMaxCount(field.String(), maxCount); err != nil {
				return fmt.Errorf("field '%s' %v", fieldType.Name, err)
			}
		}
	}

	return nil
}

func validateMaxCount(s string, max int) error {
	elements := strings.Split(s, ",")
	if len(elements) > max {
		return fmt.Errorf("must contain no more than %d elements, but got %d", max, len(elements))
	}
	return nil
}

func validateCommaSeparatedInts(s string) error {
	elements := strings.Split(s, ",")
	if len(elements) == 0 {
		return errors.New("must contain at least one element")
	}

	for i, elem := range elements {
		elem = strings.TrimSpace(elem)
		if _, err := strconv.Atoi(elem); err != nil {
			return fmt.Errorf("element '%s' at position %d is not a valid integer", elem, i)
		}
	}

	return nil
}

func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String, reflect.Array, reflect.Slice, reflect.Map, reflect.Chan:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return false
}
