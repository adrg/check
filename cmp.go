package check

import (
	"errors"
	"fmt"
	"reflect"
	"time"
)

type cmpOp int

const (
	eq cmpOp = iota + 1
	ne
	lt
	lte
	gt
	gte
)

var cmpOps = map[cmpOp]string{
	eq:  "eq",
	ne:  "ne",
	lt:  "lt",
	lte: "lte",
	gt:  "gt",
	gte: "gte",
}

var cmpErrs = map[cmpOp]string{
	eq:  "`%s` comparison failed: `%v` is not equal to `%v`",
	ne:  "`%s` comparison failed: `%v` is equal to `%v`",
	lt:  "`%s` comparison failed: `%v` is not less than `%v`",
	lte: "`%s` comparison failed: `%v` is not less than or equal to `%v`",
	gt:  "`%s` comparison failed: `%v` is not greater than `%v`",
	gte: "`%s` comparison failed: `%v` is not greater than or equal to `%v`",
}

type cmpField struct {
	op   cmpOp
	term interface{}
}

func newCmpField(op cmpOp, term interface{}) (*cmpField, error) {
	if op < eq || op > gte {
		return nil, fmt.Errorf("invalid comparison operator `%d`", op)
	}

	return &cmpField{
		op:   op,
		term: term,
	}, nil
}

func compare(x interface{}, cmp *cmpField) error {
	if cmp == nil {
		return errors.New("comparison field cannot be nil")
	}

	op := cmp.op
	if op < eq || op > gte {
		return fmt.Errorf("invalid comparison operator `%d`", op)
	}
	v := reflect.ValueOf(x)

	kind := v.Kind()
	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return compareInt64(v.Int(), cmp)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return compareUint64(v.Uint(), cmp)
	case reflect.Float32, reflect.Float64:
		return compareFloat64(v.Float(), cmp)
	case reflect.String:
		return compareString(v.String(), cmp)
	case reflect.Struct:
		if t, err := toTime(x); err == nil {
			return compareTime(t, cmp)
		}
	}

	return compareInterface(x, cmp)
}

func compareInt64(x int64, cmp *cmpField) error {
	term, err := toInt64(cmp.term)
	if err != nil {
		return err
	}
	op := cmp.op

	var ok bool
	switch op {
	case eq:
		ok = x == term
	case ne:
		ok = x != term
	case lt:
		ok = x < term
	case lte:
		ok = x <= term
	case gt:
		ok = x > term
	case gte:
		ok = x >= term
	}

	if !ok {
		return fmt.Errorf(cmpErrs[op], cmpOps[op], x, term)
	}

	return nil
}

func compareUint64(x uint64, cmp *cmpField) error {
	term, err := toUint64(cmp.term)
	if err != nil {
		return err
	}
	op := cmp.op

	var ok bool
	switch op {
	case eq:
		ok = x == term
	case ne:
		ok = x != term
	case lt:
		ok = x < term
	case lte:
		ok = x <= term
	case gt:
		ok = x > term
	case gte:
		ok = x >= term
	}

	if !ok {
		return fmt.Errorf(cmpErrs[op], cmpOps[op], x, term)
	}

	return nil
}

func compareFloat64(x float64, cmp *cmpField) error {
	term, err := toFloat64(cmp.term)
	if err != nil {
		return err
	}
	op := cmp.op

	var ok bool
	switch op {
	case eq:
		ok = x == term
	case ne:
		ok = x != term
	case lt:
		ok = x < term
	case lte:
		ok = x <= term
	case gt:
		ok = x > term
	case gte:
		ok = x >= term
	}

	if !ok {
		return fmt.Errorf(cmpErrs[op], cmpOps[op], x, term)
	}

	return nil
}

func compareString(x string, cmp *cmpField) error {
	term, err := toString(cmp.term)
	if err != nil {
		return err
	}
	op := cmp.op

	var ok bool
	switch op {
	case eq:
		ok = x == term
	case ne:
		ok = x != term
	case lt:
		ok = x < term
	case lte:
		ok = x <= term
	case gt:
		ok = x > term
	case gte:
		ok = x >= term
	}

	if !ok {
		return fmt.Errorf(cmpErrs[op], cmpOps[op], x, term)
	}

	return nil
}

func compareTime(x time.Time, cmp *cmpField) error {
	term, err := toTime(cmp.term)
	if err != nil {
		return err
	}
	op := cmp.op

	var ok bool
	switch op {
	case eq:
		ok = x.Equal(term)
	case ne:
		ok = !x.Equal(term)
	case lt:
		ok = x.Before(term)
	case lte:
		ok = x.Before(term) || x.Equal(term)
	case gt:
		ok = x.After(term)
	case gte:
		ok = x.After(term) || x.Equal(term)
	}

	if !ok {
		return fmt.Errorf(cmpErrs[op], cmpOps[op], x, term)
	}

	return nil
}

func compareInterface(x interface{}, cmp *cmpField) error {
	op := cmp.op
	term := cmp.term

	var ok bool
	switch op {
	case eq:
		ok = equals(x, term)
	case ne:
		ok = !equals(x, term)
	default:
		return fmt.Errorf("invalid operation `%s` for values `%v` and `%v`", cmpOps[op], x, term)
	}

	if !ok {
		return fmt.Errorf(cmpErrs[op], cmpOps[op], x, term)
	}

	return nil
}
