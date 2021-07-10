package webfw

import (
	"errors"
	"reflect"
	"strings"
)

func Validator(s interface{}) error {
	// 获取类型、值
	typeof := reflect.TypeOf(s)
	val := reflect.ValueOf(s)
	//将指针类型转成非指针类型
	if val.Kind() == reflect.Ptr && !val.IsNil() {
		val = val.Elem()
		typeof = typeof.Elem()
	}
	//非struct类型返回
	if val.Kind() != reflect.Struct {
		return errors.New("unsupport type")
	}
	for i := 0; i < typeof.NumField(); i++ {

		if typeof.Field(i).Tag.Get("validate") == "required" {
			if !NotBlank(val.Field(i)) {
				err := errors.New("Field " + typeof.Field(i).Name + " is nil")
				return err
			}

		}

	}
	return nil

}

func NotBlank(field reflect.Value) bool {

	switch field.Kind() {
	case reflect.String:
		return len(strings.TrimSpace(field.String())) > 0
	case reflect.Chan, reflect.Map, reflect.Slice, reflect.Array:
		return field.Len() > 0
	case reflect.Ptr, reflect.Interface, reflect.Func:
		return !field.IsNil()
	default:
		return field.IsValid() && field.Interface() != reflect.Zero(field.Type()).Interface()
	}
}
