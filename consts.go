package mime

import "errors"

var ErrInvalidMimeType = errors.New("invalid mime type")

// Prefix definition type
type Prefix string

const (
	// Standard prefix for officical types
	Standard Prefix = ""
	// Private prefix for exclusively private types
	Private Prefix = "x"
	// Vanity prefix for experimental or non public types
	Vanity Prefix = "prs"
	// Vendor prefix for vendor specific unofficial public types
	Vendor Prefix = "vnd"
)

// Suffix definition type
type Suffix string

const (
	None    Suffix = ""
	Ber     Suffix = "ber"
	CBor    Suffix = "cbor"
	CBorSeq Suffix = "cbor-seq"
	Der     Suffix = "der"
	FIS     Suffix = "fastinfoset"
	Json    Suffix = "json"
	JsonSeq Suffix = "json-seq"
	Xml     Suffix = "xml"
	Wbxml   Suffix = "wbxml"
	Zip     Suffix = "zip"
	GZip    Suffix = "gzip"
)

// Type used to define Top levels
type Type string

const (
	Audio       Type = "audio"
	Application Type = "application"
	Chemical    Type = "chemical"
	Example     Type = "example"
	Font        Type = "font"
	Image       Type = "image"
	Message     Type = "message"
	Multipart   Type = "multipart"
	Text        Type = "text"
	Video       Type = "video"
)
