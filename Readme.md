# strut

[![Build Status](https://travis-ci.org/gostrut/strut.svg?branch=master)](https://travis-ci.org/gostrut/strut)
[![GoDoc](https://godoc.org/gopkg.in/gostrut/strut.v1?status.svg)](http://godoc.org/gopkg.in/gostrut/strut.v1)

Validate your struct

## Install

    go get gopkg.in/gostrut/strut.v1

## Example

    type Person struct {
      Name    string `length_of:"2:" format_of:"^\w+$"`
      Email   string `presence_of:"true"`
      Address string
    }

    val := NewValidator()
    val.Add("presence_of", presenceof.Validator)
    val.Add("format_of", formatof.Validator)
    val.Add("length_of", lengthof.Validator)

    p := Person{}
    fields, err := val.Check(p)
    if err != nil {
      // handler validation error
    }

    if !fields.Valid() {
      // handle invalid struct
    }

## License

MIT
