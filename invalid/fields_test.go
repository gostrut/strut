package invalid

import (
	"fmt"
	"testing"

	"gopkg.in/nowk/assert.v2"
)

type tField struct {
	name      string
	validator string
}

func (f tField) Name() string {
	return f.name
}

func (f tField) Validator() string {
	return f.validator
}

func (f tField) Error() string {
	return fmt.Sprintf("invalid %s: %s", f.name, f.validator)
}

func TestFieldsImplementsError(t *testing.T) {
	f := Fields{
		"Name": []Field{
			tField{"Name", "PresenceOf"},
			tField{"Name", "FormatOf"},
		},
		"Email": []Field{
			tField{"Email", "PresenceOf"},
		},
	}

	assert.Equal(t, `invalid fields:
  * Name:
    - invalid Name: PresenceOf
    - invalid Name: FormatOf
  * Email:
    - invalid Email: PresenceOf`, f.Error())
}
