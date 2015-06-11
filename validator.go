package strut

import (
	"fmt"
	"reflect"

	"gopkg.in/gostrut/invalid.v0"
)

type ValidatorFunc func(string, string, *reflect.Value) (invalid.Field, error)

func (v ValidatorFunc) Validate(
	f, t string, r *reflect.Value) (invalid.Field, error) {

	return v(f, t, r)
}

type Validator struct {
	validators map[string]ValidatorFunc
}

// NewValidator returns a new Validator
func NewValidator() *Validator {
	return &Validator{
		validators: make(map[string]ValidatorFunc),
	}
}

// Checks appends validators to be validated against
func (s *Validator) Checks(tagName string, fn ValidatorFunc) {
	s.validators[tagName] = fn
}

// Validates interates through struct fiels and validates where applicable
func (s Validator) Validates(obj interface{}) (invalid.Fields, error) {
	if len(s.validators) == 0 {
		return nil, nil // if no validators
	}

	to := reflect.TypeOf(obj)
	if to.Kind() != reflect.Struct {
		return nil, NonStructError{to.Name()}
	}

	fields := invalid.NewFields()
	len := to.NumField()
	for i := 0; i < len; i++ {
		fld := to.Field(i)
		if "" == fld.Tag {
			continue // if no StructTag
		}

		vo := reflect.ValueOf(obj)
		vf := vo.Field(i)
		for k, v := range s.validators {
			tag := fld.Tag.Get(k)
			if "" == tag {
				continue // if tag value is ""
			}

			f, err := v.Validate(fld.Name, tag, &vf)
			if err != nil {
				return nil, err
			}

			if f != nil {
				fields.Add(fld.Name, f)
			}
		}
	}

	return fields, nil
}

// NonStructError is an error implement for non struct types
type NonStructError struct {
	Name string
}

func (e NonStructError) Error() string {
	return fmt.Sprintf("error: non-struct: cannot validate type of %s", e.Name)
}
