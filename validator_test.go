package strut

import (
	"testing"

	"gopkg.in/nowk/assert.v2"
)

func TestStrutInvalid(t *testing.T) {
	type Person struct {
		Name    string `length_of:"2:" format_of:"^foo$"`
		Email   string `presence_of:"true"`
		Address string
	}

	val := NewValidator()
	val.Checks("presence_of", invalidated)
	val.Checks("format_of", invalidated)
	val.Checks("length_of", invalidated)

	a := Person{}
	fields, err := val.Validates(a)
	assert.Nil(t, err)
	assert.False(t, fields.Valid())
	assert.Equal(t, 2, fields.Len())

	fname := fields.Get("Name")
	femail := fields.Get("Email")
	assert.Equal(t, 2, len(fname))
	assert.Equal(t, 1, len(femail))
}

func TestStrutValid(t *testing.T) {
	type Person struct {
		Name string `is_valid:"true"`
	}

	val := NewValidator()
	val.Checks("is_valid", validated)

	a := Person{}
	fields, err := val.Validates(a)
	assert.Nil(t, err)
	assert.True(t, fields.Valid())
}

func TestStrutError(t *testing.T) {
	type Post struct {
		Body string `length_of:"2"`
	}

	val := NewValidator()
	val.Checks("length_of", errored)

	a := Post{}
	fields, err := val.Validates(a)
	assert.Nil(t, fields)
	assert.Equal(t, "Whoops", err.Error())
}

func TestStrutOnlyValidatesStructType(t *testing.T) {
	val := NewValidator()
	val.Checks("presence_of", invalidated)

	_, err := val.Validates("str")
	assert.Equal(t, err.Error(), "not a struct: cannot validate type of 'string'")
}
