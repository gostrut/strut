# strut

<!-- [![Build Status](https://travis-ci.org/gostrut/strut.svg?branch=master)](https://travis-ci.org/gostrut/strut) -->
[![GoDoc](https://godoc.org/github.com/gostrut/strut?status.svg)](http://godoc.org/github.com/gostrut/strut)

Validate your struct

# Example

    type Person struct {
      Name    string `length_of:"2:" format_of:"^\w+$"`
      Email   string `presence_of:"true"`
      Address string
    }

    val := NewValidator()
    val.Checks("presence_of", presenceof.Validator)
    val.Checks("format_of", formatof.Validator)
    val.Checks("length_of", lengthof.Validator)

    p := Person{}
    fields, err := val.Validates(p)
    if err != nil {
      // handler validation error
    }

    if !fields.Valid() {
      // handle invalid struct
    }

# MIT

License