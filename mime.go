// Package mime provides the basis media type(s)
package mime

import (
	"strings"
	"unicode"
)

// MIME in form: Type "/" [Prefix "."] Subtype ["+" Suffix] *[";" Parameter]
type MIME struct {
	Type
	Prefix
	Subtype string
	Suffix
	Parameters map[string][]string
}

func (m *MIME) IsValid() bool {
	if m == nil {
		return false
	}
	// TBD: check if other fields are defined as well
	_, typeOk := RegTypes[m.Type]
	_, preOk := RegPrefix[m.Prefix]
	return typeOk && preOk
}

func (m *MIME) String() string {
	if m == nil {
		return ""
	}
	var s = string(m.Type)
	if m.Prefix != Standard {
		s += string(m.Prefix) + "."
	}
	s += string(m.Subtype)
	if m.Suffix != None {
		s += "+" + string(m.Suffix)
	}
	for key, vals := range m.Parameters {
		if len(key) > 0 && len(vals) > 0 {
			for _, v := range vals {
				s += ";" + key + "=" + v
			}
		}
	}
	return s
}

// MustParse panics in case of a parse error instead of returning the error
func MustParse(s string) *MIME {
	if v, e := Parse(s); e == nil {
		return v
	} else {
		panic(e)
	}
}

// Parse returns on success the mime.Variety representaiton of the passed string mime otherwise an error
func Parse(s string) (*MIME, error) {
	if len(s) == 0 {
		return nil, ErrInvalidMimeType
	}
	var (
		e error
		v = &MIME{
			Parameters: map[string][]string{},
		}
		m, p, _ = strings.Cut(s, ";")
	)
	if len(m) == 0 {
		return nil, ErrInvalidMimeType
	}
	e = parseTree(m, v)
	if e == nil {
		e = parseParams(p, v)
	}

	if e != nil {
		v = nil
	}
	return v, e
}

func parseTree(s string, v *MIME) error {
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

func parseParams(s string, v *MIME) error {
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
