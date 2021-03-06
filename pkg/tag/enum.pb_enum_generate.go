// Code generated by github.com/infraboard/mcube
// DO NOT EDIT

package tag

import (
	"bytes"
	"fmt"
	"strings"
)

// ParseScopeTypeFromString Parse ScopeType from string
func ParseScopeTypeFromString(str string) (ScopeType, error) {
	key := strings.Trim(string(str), `"`)
	v, ok := ScopeType_value[strings.ToUpper(key)]
	if !ok {
		return 0, fmt.Errorf("unknown ScopeType: %s", str)
	}

	return ScopeType(v), nil
}

// Equal type compare
func (t ScopeType) Equal(target ScopeType) bool {
	return t == target
}

// IsIn todo
func (t ScopeType) IsIn(targets ...ScopeType) bool {
	for _, target := range targets {
		if t.Equal(target) {
			return true
		}
	}

	return false
}

// MarshalJSON todo
func (t ScopeType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(strings.ToUpper(t.String()))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

// UnmarshalJSON todo
func (t *ScopeType) UnmarshalJSON(b []byte) error {
	ins, err := ParseScopeTypeFromString(string(b))
	if err != nil {
		return err
	}
	*t = ins
	return nil
}

// ParseValueFromFromString Parse ValueFrom from string
func ParseValueFromFromString(str string) (ValueFrom, error) {
	key := strings.Trim(string(str), `"`)
	v, ok := ValueFrom_value[strings.ToUpper(key)]
	if !ok {
		return 0, fmt.Errorf("unknown ValueFrom: %s", str)
	}

	return ValueFrom(v), nil
}

// Equal type compare
func (t ValueFrom) Equal(target ValueFrom) bool {
	return t == target
}

// IsIn todo
func (t ValueFrom) IsIn(targets ...ValueFrom) bool {
	for _, target := range targets {
		if t.Equal(target) {
			return true
		}
	}

	return false
}

// MarshalJSON todo
func (t ValueFrom) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(strings.ToUpper(t.String()))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

// UnmarshalJSON todo
func (t *ValueFrom) UnmarshalJSON(b []byte) error {
	ins, err := ParseValueFromFromString(string(b))
	if err != nil {
		return err
	}
	*t = ins
	return nil
}
