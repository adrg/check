package check

import (
	"errors"
	"fmt"
	"reflect"
	"time"
)

func equal(x, y interface{}) bool {
	return reflect.DeepEqual(x, y)
}

func isEmpty(x interface{}) bool {
	if x == nil {
		return true
	}

	v := reflect.ValueOf(x)
	switch v.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Interface, reflect.Ptr:
		return isEmpty(v.Elem().Interface())
	}

	return reflect.DeepEqual(x, reflect.Zero(reflect.TypeOf(x)).Interface())
}

func toInt64(x interface{}) (int64, error) {
	if x == nil {
		return 0, errors.New("cannot convert nil to type int64")
	}
	v := reflect.ValueOf(x)

	kind := v.Kind()
	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int(), nil
	}

	return 0, fmt.Errorf("cannot convert `%v` to type int64", kind)
}

func toUint64(x interface{}) (uint64, error) {
	if x == nil {
		return 0, errors.New("cannot convert nil to type uint64")
	}
	v := reflect.ValueOf(x)

	kind := v.Kind()
	switch kind {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint(), nil
	}

	return 0, fmt.Errorf("cannot convert `%v` to type uint64", kind)
}

func toFloat64(x interface{}) (float64, error) {
	if x == nil {
		return 0, errors.New("cannot convert nil to type float64")
	}
	v := reflect.ValueOf(x)

	kind := v.Kind()
	switch kind {
	case reflect.Float32, reflect.Float64:
		return v.Float(), nil
	}

	return 0, fmt.Errorf("cannot convert `%v` to type float64", kind)
}

func toString(x interface{}) (string, error) {
	if x == nil {
		return "", errors.New("cannot convert nil to type string")
	}
	v := reflect.ValueOf(x)

	kind := v.Kind()
	if kind == reflect.String {
		return v.String(), nil
	}

	return "", fmt.Errorf("cannot convert `%v` to type string", kind)
}

func toTime(x interface{}) (time.Time, error) {
	if x == nil {
		return time.Time{}, errors.New("cannot convert nil to type time.Time")
	}

	v, ok := x.(time.Time)
	if !ok {
		return time.Time{}, fmt.Errorf("cannot convert `%v` to time.Time", reflect.TypeOf(x))
	}

	return v, nil
}
