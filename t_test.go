package strut

import "fmt"
import "reflect"
import "github.com/gostrut/invalid"

type tField struct {
	name   string
	tagStr string
	msg    string
}

func (f tField) Name() string {
	return f.name
}

func (f tField) Validator() string {
	return ""
}

func (f tField) Error() string {
	return fmt.Sprintf("%s %s; tag: `%s`", f.name, f.msg, f.tagStr)
}

func invalidated(name, tagStr string, vo *reflect.Value) (invalid.Field, error) {
	f := tField{
		name:   name,
		tagStr: tagStr,
		msg:    "is invalid",
	}

	return f, nil
}

func errored(name, tagStr string, vo *reflect.Value) (invalid.Field, error) {
	return nil, fmt.Errorf("Whoops")
}

func validated(_, _ string, _ *reflect.Value) (invalid.Field, error) {
	return nil, nil
}
