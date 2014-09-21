package strut

import "testing"
import "github.com/nowk/assert"

func TestStrut(t *testing.T) {
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
	f1 := fields[0]
	f2 := fields[1]
	f3 := fields[2]
	assert.Nil(t, err)
	assert.False(t, fields.Valid())
	assert.Equal(t, 3, fields.Len())
	assert.Equal(t, "Name is invalid; tag: `^foo$`", f1.Error())
	assert.Equal(t, "Name is invalid; tag: `2:`", f2.Error())
	assert.Equal(t, "Email is invalid; tag: `true`", f3.Error())
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
