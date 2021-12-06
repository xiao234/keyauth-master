// Code generated by github.com/infraboard/mcube
// DO NOT EDIT

package token

import (
	"bytes"
	"fmt"
	"strings"
)

// ParseGrantTypeFromString Parse GrantType from string
func ParseGrantTypeFromString(str string) (GrantType, error) {
	key := strings.Trim(string(str), `"`)
	v, ok := GrantType_value[strings.ToUpper(key)]
	if !ok {
		return 0, fmt.Errorf("unknown GrantType: %s", str)
	}

	return GrantType(v), nil
}

// Equal type compare
func (t GrantType) Equal(target GrantType) bool {
	return t == target
}

// IsIn todo
func (t GrantType) IsIn(targets ...GrantType) bool {
	for _, target := range targets {
		if t.Equal(target) {
			return true
		}
	}

	return false
}

// MarshalJSON todo
func (t GrantType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(strings.ToUpper(t.String()))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

// UnmarshalJSON todo
func (t *GrantType) UnmarshalJSON(b []byte) error {
	ins, err := ParseGrantTypeFromString(string(b))
	if err != nil {
		return err
	}
	*t = ins
	return nil
}

// ParseTokenTypeFromString Parse TokenType from string
func ParseTokenTypeFromString(str string) (TokenType, error) {
	key := strings.Trim(string(str), `"`)
	v, ok := TokenType_value[strings.ToUpper(key)]
	if !ok {
		return 0, fmt.Errorf("unknown TokenType: %s", str)
	}

	return TokenType(v), nil
}

// Equal type compare
func (t TokenType) Equal(target TokenType) bool {
	return t == target
}

// IsIn todo
func (t TokenType) IsIn(targets ...TokenType) bool {
	for _, target := range targets {
		if t.Equal(target) {
			return true
		}
	}

	return false
}

// MarshalJSON todo
func (t TokenType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(strings.ToUpper(t.String()))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

// UnmarshalJSON todo
func (t *TokenType) UnmarshalJSON(b []byte) error {
	ins, err := ParseTokenTypeFromString(string(b))
	if err != nil {
		return err
	}
	*t = ins
	return nil
}

// ParseBlockTypeFromString Parse BlockType from string
func ParseBlockTypeFromString(str string) (BlockType, error) {
	key := strings.Trim(string(str), `"`)
	v, ok := BlockType_value[strings.ToUpper(key)]
	if !ok {
		return 0, fmt.Errorf("unknown BlockType: %s", str)
	}

	return BlockType(v), nil
}

// Equal type compare
func (t BlockType) Equal(target BlockType) bool {
	return t == target
}

// IsIn todo
func (t BlockType) IsIn(targets ...BlockType) bool {
	for _, target := range targets {
		if t.Equal(target) {
			return true
		}
	}

	return false
}

// MarshalJSON todo
func (t BlockType) MarshalJSON() ([]byte, error) {
	b := bytes.NewBufferString(`"`)
	b.WriteString(strings.ToUpper(t.String()))
	b.WriteString(`"`)
	return b.Bytes(), nil
}

// UnmarshalJSON todo
func (t *BlockType) UnmarshalJSON(b []byte) error {
	ins, err := ParseBlockTypeFromString(string(b))
	if err != nil {
		return err
	}
	*t = ins
	return nil
}
