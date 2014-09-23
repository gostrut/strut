package strut

import "fmt"
import "reflect"
import "github.com/gostrut/invalid"
import "github.com/gostrut/validator"

type strut struct {
	validators map[string]validator.Func
}

// NewValidator returns a new strut
func NewValidator() *strut {
	return &strut{
		validators: make(map[string]validator.Func),
	}
}

// Checks appends validators to be validated against
func (s *strut) Checks(tagName string, fn validator.Func) {
	s.validators[tagName] = fn
}

// Validates interates through struct fiels and validates where applicable
func (s strut) Validates(obj interface{}) (invalid.Fields, error) {
	if len(s.validators) == 0 {
		return nil, nil // if no validators
	}

	var invf invalid.Fields
	var to = reflect.TypeOf(obj)
	if to.Kind() != reflect.Struct {
		return nil, NonStructError{to.Name()}
	}

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
				invf = append(invf, f)
			}
		}
	}

	return invf, nil
}

// NonStructError is an error implement for non struct types
type NonStructError struct {
	Name string
}

func (e NonStructError) Error() string {
	return fmt.Sprintf("error: non-struct: cannot validate type of %s", e.Name)
}
