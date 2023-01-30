// Package mime provides the basis media type(s)
package libmime

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidMimeType = errors.New("invalid mime type")

// Prefix definition type
type Prefix string

const (
	// Standard prefix for officical types
	Standard = Prefix("")
	// Private prefix for exclusively private types
	Private = Prefix("x")
	// Vanity prefix for experimental or non public types
	Vanity = Prefix("prs")
	// Vendor prefix for vendor specific unofficial public types
	Vendor = Prefix("vnd")
)

// MIME type prefix register
var Prefixes = map[Prefix]struct{}{
	Standard: {},
	Private:  {},
	Vanity:   {},
	Vendor:   {},
}

type Suffix string

const (
	Ber     = Suffix("ber")
	CBor    = Suffix("cbor")
	CBorSeq = Suffix("cbor-seq")
	Der     = Suffix("der")
	FIS     = Suffix("fastinfoset")
	Json    = Suffix("json")
	JsonSeq = Suffix("json-seq")
	Xml     = Suffix("xml")
	Wbxml   = Suffix("wbxml")
	Zip     = Suffix("zip")
	GZip    = Suffix("gzip")
)

// Type definition type
type Type string

// MIME types official registed (one unofficial) by IANA as of Dec 2020
const (
	Audio Type = "audio"
	App   Type = "application"
	// common but unofficial top-level name
	Chemical Type = "chemical"
	Example  Type = "example"
	Font     Type = "font"
	Image    Type = "image"
	Message  Type = "message"
	MultiPt  Type = "multipart"
	Text     Type = "text"
	Video    Type = "video"
)

// MIME type register
var Types = map[Type]struct{}{
	Audio:    {},
	App:      {},
	Chemical: {},
	Example:  {},
	Font:     {},
	Image:    {},
	Message:  {},
	MultiPt:  {},
	Text:     {},
	Video:    {},
}

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
	_, typeOk := Types[v.Type]
	_, preOk := Prefixes[v.Prefix]
	return typeOk && preOk
}

func (v *Variety) String() string {
	if v == nil {
		return ""
	}
	return ""
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
