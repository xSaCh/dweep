package util

import (
	"reflect"
)

func MissingStructFields(st any) []string {
	stV := reflect.ValueOf(st)

	missing := []string{}
	for i := 0; i < stV.NumField(); i++ {
		field1 := stV.Field(i)

		if isEmpty(field1) {
			missing = append(missing, stV.Type().Field(i).Name)
		}

	}
	return missing
}

func isEmpty(val reflect.Value) bool {
	switch val.Kind() {
	case reflect.String:
		return val.String() == ""
	case reflect.Slice, reflect.Array:
		return val.Len() == 0
	case reflect.Map:
		return val.Len() == 0
	case reflect.Ptr:
		return val.IsNil()
	case reflect.Struct:
		return val.Interface() == reflect.Zero(val.Type()).Interface()
	default:
		return val.Interface() == reflect.Zero(val.Type()).Interface()
	}
}
