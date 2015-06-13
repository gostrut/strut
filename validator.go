package strut

import (
	"fmt"
	"reflect"

	"gopkg.in/gostrut/strut.v1/invalid"
)

// ValidatorFunc is the validator func signature. The first 2 strings are the
// Field name and Field tag (what is within "", eg. json:"...") value
// respectively
type ValidatorFunc func(string, string, *reflect.Value) (invalid.Field, error)

func (v ValidatorFunc) Validate(
	n, t string, r *reflect.Value) (invalid.Field, error) {

	return v(n, t, r)
}

type validators map[string]ValidatorFunc

type Validator struct {
	validators validators
}

func NewValidator() *Validator {
	return &Validator{
		validators: make(validators),
	}
}

// Add adds a ValidatorFunc for a given tag
func (v *Validator) Add(n string, fn ValidatorFunc) {
	v.validators[n] = fn
}

func (v *Validator) validate(
	f reflect.StructField, r reflect.Value) ([]invalid.Field, error) {

	t := f.Tag
	if t == "" {
		return nil, nil
	}

	i := make([]invalid.Field, 0, len(v.validators))

	for k, fn := range v.validators {
		tval := t.Get(k)
		if tval == "" {
			continue // if tag value is ""
		}

		f, err := fn.Validate(f.Name, tval, &r)
		if err != nil {
			return nil, err
		}
		if f != nil {
			i = append(i, f)
		}
	}

	return i, nil
}

// Check checks the given struct against the available validators, returning
// an collection of invalid fields or error
func (v *Validator) Check(o interface{}) (invalid.Fields, error) {
	if len(v.validators) == 0 {
		return nil, nil // if no validators
	}

	r := reflect.ValueOf(o)

	t := r.Type()
	if t.Kind() != reflect.Struct {
		return nil,
			fmt.Errorf("not a struct: cannot validate type of '%s'", t.Name())
	}

	e := make(invalid.Fields)

	i := 0
	n := t.NumField()
	for ; i < n; i++ {
		f, err := v.validate(t.Field(i), r.Field(i))
		if err != nil {
			return nil, err
		}

		e.Add(f...)
	}

	return e, nil
}
