// Package mime provides the basis media type(s)
package mime

import (
	"strings"
	"unicode"
)

// Variety describes a MIME type of form: Type "/" [Prefix "."] Subtype ["+" Suffix] *[";" Parameter]
type Variety struct {
	Type
	Prefix
	Subtype string
	Suffix
	Parameters map[string][]string
}

func (v *Variety) IsValid() bool {
	if v == nil {
		return false
	}
	_, typeOk := RegTypes[v.Type]
	_, preOk := RegPrefix[v.Prefix]
	return typeOk && preOk
}

func (v *Variety) String() string {
	if v == nil {
		return ""
	}
	return ""
}

func MustParse(s string) *Variety {
	if v, e := Parse(s); e == nil {
		return v
	} else {
		panic(e)
	}
}

func Parse(s string) (*Variety, error) {
	if len(s) == 0 {
		return nil, ErrInvalidMimeType
	}
	var (
		e error
		v = &Variety{
			Parameters: map[string][]string{},
		}
		m, p, _ = strings.Cut(s, ";")
	)
	if len(m) == 0 {
		return nil, ErrInvalidMimeType
	}
	e = parseMimePart(m, v)
	if e == nil {
		e = parseParams(p, v)
	}

	if e != nil {
		v = nil
	}

	return v, e
}

func parseMimePart(s string, v *Variety) error {
	if typeIdx := strings.Index(s, "/"); typeIdx > -1 {
		v.Type = Type(s[:typeIdx])
		s = s[typeIdx+1:]
	} else {
		return ErrInvalidMimeType
	}

	if preIdx := strings.Index(s, "."); preIdx > -1 {
		v.Prefix = Prefix(s[:preIdx])
		s = s[preIdx+1:]
	}
	if postIdx := strings.Index(s, "+"); postIdx > -1 {
		v.Subtype = s[:postIdx]
		v.Suffix = Suffix(s[postIdx+1:])
	} else {
		v.Subtype = s
	}
	return nil
}

func parseParams(s string, v *Variety) error {
	s = strings.TrimFunc(s, unicode.IsSpace)
	if s != "" {
		var pairs = strings.Split(s, ";")
		for _, pair := range pairs {
			pair := strings.TrimLeftFunc(pair, unicode.IsSpace)
			if key, value, found := strings.Cut(pair, "="); found && key != "" {
				v.Parameters[key] = append(v.Parameters[key], value)
			} else {
				return ErrInvalidMimeType
			}
		}
	}
	return nil
}
