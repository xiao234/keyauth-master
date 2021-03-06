// Code generated by github.com/infraboard/mcube
// DO NOT EDIT

package department

import (
	"bytes"
	"fmt"
	"strings"
)

// ParseApplicationFormStatusFromString Parse ApplicationFormStatus from string
func ParseApplicationFormStatusFromString(str string) (ApplicationFormStatus, error) {
	key := strings.Trim(string(str), `"`)
	v, ok := ApplicationFormStatus_value[strings.ToUpper(key)]
	if !ok {
		return 0, fmt.Errorf("unknown ApplicationFormStatus: %s", str)
	}

	return ApplicationFormStatus(v), nil
}

// Equal type compare
func (t ApplicationFormStatus) Equal(target ApplicationFormStatus) bool {
	return t == target
}

// IsIn todo
func (t ApplicationFormStatus) IsIn(targets ...ApplicationFormStatus) bool {
	for _, target := range targets {
		if t.Equal(target) {
			return true
		}
	}

	return false
}

// MarshalJSON todo
func (t ApplicationFormStatus) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(strings.ToUpper(t.String()))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

// UnmarshalJSON todo
func (t *ApplicationFormStatus) UnmarshalJSON(b []byte) error {
	ins, err := ParseApplicationFormStatusFromString(string(b))
	if err != nil {
		return err
	}
	*t = ins
	return nil
}
