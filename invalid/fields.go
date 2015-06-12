package invalid

// Field is an interface for an invalidated field
type Field interface {
	// Name is the field name
	Name() string

	// Validator is the name of the validator used
	Validator() string

	// Error is the validation error received from the validation
	Error() string
}

// Fields type is a map of []Field to a string (the field name)
type Fields map[string][]Field

func (f Fields) Len() int {
	return len(f)
}

func (f Fields) Valid() bool {
	return f.Len() == 0
}

func (f Fields) Get(name string) []Field {
	return f[name]
}

func (f Fields) Add(o ...Field) {
	for _, v := range o {
		k := v.Name()

		f[k] = append(f[k], v)
	}
}
