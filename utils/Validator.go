package utils

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

type Rules map[string][]string

func NotEmpty() string {
	return "notEmpty"
}

func Verify(st interface{}, roleMap Rules) error {
	compareMap := map[string]bool{
		"lt": true,
		"le": true,
		"eq": true,
		"ne": true,
		"ge": true,
		"gt": true,
	}
	typ := reflect.TypeOf(st)
	val := reflect.ValueOf(st)

	kd := val.Kind()
	if kd != reflect.Struct {
		return errors.New("expect struct")
	}

	num := val.NumField()
	// 遍历结构体的所有字段
	for i := 0; i < num; i++ {
		tagVal := typ.Field(i)
		val := val.Field(i)
		if len(roleMap[tagVal.Name]) > 0 {
			for _, v := range roleMap[tagVal.Name] {
				switch {
				case v == NotEmpty():
					if isBlank(val) {
						return errors.New(tagVal.Name + "值不能为空")
					}
				case v:

				}

			}
		}
	}
}

// 非空校验
// true 空 ;false 非空
func isBlank(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.String:
		return value.Len() == 0
	case reflect.Bool:
		return !value.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return value.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return value.IsNil()
	}
	return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
}

func compareVerify(value reflect.Value, verifyVal string) bool {
	switch value.Kind() {
	case reflect.String, reflect.Array, reflect.Slice:
		compare(value.Int(), verifyVal)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
	case reflect.Float32, reflect.Float64:
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
	default:
		return false
	}

}

func compare(value interface{}, verifyStr string) bool {
	verifyStrArr := strings.Split(verifyStr, "=")
	val := reflect.ValueOf(value)
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		VInt, VErr := strconv.ParseInt(verifyStrArr[1], 10, 64)
		if VErr != nil {
			return false
		}
		switch verifyStrArr[0] {
		case "lt":
			return val.Int() < VInt
		case "le":
			return val.Int() <= VInt
		case "eq":
			return val.Int() == VInt
		case "ne":
			return val.Int() != VInt
		case "gt":
			return val.Int() > VInt
		case "ge":
			return val.Int() >= VInt
		default:
			return false
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:

	}
}
