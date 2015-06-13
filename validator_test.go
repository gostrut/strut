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
	val.Add("presence_of", invalidated)
	val.Add("format_of", invalidated)
	val.Add("length_of", invalidated)

	a := Person{}
	fields, err := val.Check(a)
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
	val.Add("is_valid", validated)

	a := Person{}
	fields, err := val.Check(a)
	assert.Nil(t, err)
	assert.True(t, fields.Valid())
}

func TestStrutError(t *testing.T) {
	type Post struct {
		Body string `length_of:"2"`
	}

	val := NewValidator()
	val.Add("length_of", errored)

	a := Post{}
	fields, err := val.Check(a)
	assert.Nil(t, fields)
	assert.Equal(t, "Whoops", err.Error())
}

func TestStrutOnlyValidatesStructType(t *testing.T) {
	val := NewValidator()
	val.Add("presence_of", invalidated)

	_, err := val.Check("str")
	assert.Equal(t, err.Error(), "not a struct: cannot validate type of 'string'")
}
