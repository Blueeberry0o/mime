package mime

import "strings"

// MIME type prefix register
var RegPrefix = map[Prefix]struct{}{
	Standard: {},
	Private:  {},
	Vanity:   {},
	Vendor:   {},
}

// MIME type suffix register
var RegSuffix = map[Suffix]struct{}{
	Ber:     {},
	CBor:    {},
	CBorSeq: {},
	Der:     {},
	FIS:     {},
	Json:    {},
	JsonSeq: {},
	Xml:     {},
	Wbxml:   {},
	Zip:     {},
	GZip:    {},
}

// MIME top type register
var RegTypes = map[Type]struct{}{
	Audio:       {},
	Application: {},
	Chemical:    {},
	Example:     {},
	Font:        {},
	Image:       {},
	Message:     {},
	Multipart:   {},
	Text:        {},
	Video:       {},
}

// MIME type Register
var RegMIME = map[string]*MIME{
	"audio/3gpp":               MustParse("audio/3gpp"),
	"video/3gpp":               MustParse("video/3gpp"),
	"audio/3gpp2":              MustParse("audio/3gpp2"),
	"video/3gpp2":              MustParse("video/3gpp2"),
	"audio/aac":                MustParse("audio/aac"),
	"application/x-abiword":    MustParse("application/x-abiword"),
	"application/x-freearc":    MustParse("application/x-freearc"),
	"image/avif":               MustParse("image/avif"),
	"video/x-msvideo":          MustParse("video/x-msvideo"),
	"application/octet-stream": MustParse("application/octet-stream"),
	"image/bmp":                MustParse("image/bmp"),
	"application/x-bzip":       MustParse("application/x-bzip"),
	"application/x-bzip2":      MustParse("application/x-bzip2"),
	"application/x-cdf":        MustParse("application/x-cdf"),
	"application/x-csh":        MustParse("application/x-csh"),
	"text/css":                 MustParse("text/css"),
	"text/csv":                 MustParse("text/csv"),
	"application/msword":       MustParse("application/msword"),
	"application/epub+zip":     MustParse("application/epub+zip"),
	"application/gzip":         MustParse("application/gzip"),
	"image/gif":                MustParse("image/gif"),
	"text/html":                MustParse("text/html"),
	"text/htm":                 MustParse("text/html"),
	"text/calendar":            MustParse("text/calendar"),
	"image/vnd.microsoft.icon": MustParse("image/vnd.microsoft.icon"),
	"application/java-archive": MustParse("application/java-archive"),
	"image/jpeg":               MustParse("image/jpeg"),
	"image/jpg":                MustParse("image/jpeg"),
	"application/json":         MustParse("application/json"),
	"application/ld+json":      MustParse("application/ld+json"),
	"audio/x-midi":             MustParse("audio/x-midi"),
	"audio/midi":               MustParse("audio/midi"),
	"text/javascript":          MustParse("text/javascript"),
	"audio/mpeg":               MustParse("audio/mpeg"),
	"video/mp4":                MustParse("video/mp4"),
	"video/mpeg":               MustParse("video/mpeg"),
	"audio/ogg":                MustParse("audio/ogg"),
	"video/ogg":                MustParse("video/ogg"),
	"application/ogg":          MustParse("application/ogg"),
	"audio/opus":               MustParse("audio/opus"),
	"font/otf":                 MustParse("font/otf"),
	"image/png":                MustParse("image/png"),
	"application/pdf":          MustParse("application/pdf"),
	"application/x-httpd-php":  MustParse("application/x-httpd-php"),
	"application/vnd.rar":      MustParse("application/vnd.rar"),
	"application/rtf":          MustParse("application/rtf"),
	"application/x-sh":         MustParse("application/x-sh"),
	"image/svg+xml":            MustParse("image/svg+xml"),
	"application/x-tar":        MustParse("application/x-tar"),
	"image/tiff":               MustParse("image/tiff"),
	"video/mp2t":               MustParse("video/mp2t"),
	"font/ttf":                 MustParse("font/ttf"),
	"text/plain":               MustParse("text/plain"),
	"audio/wav":                MustParse("audio/wav"),
	"audio/webm":               MustParse("audio/webm"),
	"video/webm":               MustParse("video/webm"),
	"image/webp":               MustParse("image/webp"),
	"font/woff":                MustParse("font/woff"),
	"font/woff2":               MustParse("font/woff2"),
	"application/xml":          MustParse("application/xml"),
	"application/zip":          MustParse("application/zip"),
	"application/vnd.visio":    MustParse("application/vnd.visio"),
	"application/xhtml+xml":    MustParse("application/xhtml+xml"),

	"application/vnd.ms-excel":            MustParse("application/vnd.ms-excel"),
	"application/vnd.amazon.ebook":        MustParse("application/vnd.amazon.ebook"),
	"application/vnd.ms-fontobject":       MustParse("application/vnd.ms-fontobject"),
	"application/vnd.ms-powerpoint":       MustParse("application/vnd.ms-powerpoint"),
	"application/vnd.mozilla.xul+xml":     MustParse("application/vnd.mozilla.xul+xml"),
	"application/vnd.apple.installer+xml": MustParse("application/vnd.apple.installer+xml"),

	"application/x-7z-compressed": MustParse("application/x-7z-compressed"),

	"application/vnd.oasis.opendocument.text":         MustParse("application/vnd.oasis.opendocument.text"),
	"application/vnd.oasis.opendocument.spreadsheet":  MustParse("application/vnd.oasis.opendocument.spreadsheet"),
	"application/vnd.oasis.opendocument.presentation": MustParse("application/vnd.oasis.opendocument.presentation"),

	"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet":         MustParse("application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"),
	"application/vnd.openxmlformats-officedocument.wordprocessingml.document":   MustParse("application/vnd.openxmlformats-officedocument.wordprocessingml.document"),
	"application/vnd.openxmlformats-officedocument.presentationml.presentation": MustParse("application/vnd.openxmlformats-officedocument.presentationml.presentation"),
}

// Extensions register and related MIME-keys
var RegExts = map[string][]string{
	".3gp":    {"audio/3gpp", "video/3gpp"},
	".3g2":    {"audio/3gpp2", "video/3gpp2"},
	".7z":     {"application/x-7z-compressed"},
	".aac":    {"audio/aac"},
	".abw":    {"application/x-abiword"},
	".arc":    {"application/x-freearc"},
	".avif":   {"image/avif"},
	".avi":    {"video/x-msvideo"},
	".azw":    {"application/vnd.amazon.ebook"},
	".bin":    {"application/octet-stream"},
	".bmp":    {"image/bmp"},
	".bz":     {"application/x-bzip"},
	".bz2":    {"application/x-bzip2"},
	".cda":    {"application/x-cdf"},
	".csh":    {"application/x-csh"},
	".css":    {"text/css"},
	".csv":    {"text/csv"},
	".doc":    {"application/msword"},
	".docx":   {"application/vnd.openxmlformats-officedocument.wordprocessingml.document"},
	".eot":    {"application/vnd.ms-fontobject"},
	".epub":   {"application/epub+zip"},
	".gz":     {"application/gzip"},
	".gif":    {"image/gif"},
	".htm":    {"text/html"},
	".html":   {"text/html"},
	".ico":    {"image/vnd.microsoft.icon"},
	".ics":    {"text/calendar"},
	".jar":    {"application/java-archive"},
	".jpeg":   {"image/jpeg"},
	".jpg":    {"image/jpeg"},
	".json":   {"application/json"},
	".jsonld": {"application/ld+json"},
	".mid":    {"audio/midi", "audio/x-midi"},
	".midi":   {"audio/midi", "audio/x-midi"},
	".mjs":    {"text/javascript"},
	".mp3":    {"audio/mpeg"},
	".mp4":    {"video/mp4"},
	".mpeg":   {"video/mpeg"},
	".mpkg":   {"application/vnd.apple.installer+xml"},
	".odp":    {"application/vnd.oasis.opendocument.presentation"},
	".ods":    {"application/vnd.oasis.opendocument.spreadsheet"},
	".odt":    {"application/vnd.oasis.opendocument.text"},
	".oga":    {"audio/ogg"},
	".ogv":    {"video/ogg"},
	".ogx":    {"application/ogg"},
	".opus":   {"audio/opus"},
	".otf":    {"font/otf"},
	".png":    {"image/png"},
	".pdf":    {"application/pdf"},
	".php":    {"application/x-httpd-php"},
	".ppt":    {"application/vnd.ms-powerpoint"},
	".pptx":   {"application/vnd.openxmlformats-officedocument.presentationml.presentation"},
	".rar":    {"application/vnd.rar"},
	".rtf":    {"application/rtf"},
	".sh":     {"application/x-sh"},
	".svg":    {"image/svg+xml"},
	".tar":    {"application/x-tar"},
	".tif":    {"image/tiff"},
	".tiff":   {"image/tiff"},
	".ts":     {"video/mp2t"},
	".ttf":    {"font/ttf"},
	".txt":    {"text/plain"},
	".vsd":    {"application/vnd.visio"},
	".wav":    {"audio/wav"},
	".weba":   {"audio/webm"},
	".webm":   {"video/webm"},
	".webp":   {"image/webp"},
	".woff":   {"font/woff"},
	".woff2":  {"font/woff2"},
	".xhtml":  {"application/xhtml+xml"},
	".xls":    {"application/vnd.ms-excel"},
	".xlsx":   {"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"},
	".xml":    {"application/xml"},
	".xul":    {"application/vnd.mozilla.xul+xml"},
	".zip":    {"application/zip"},
}

// GetMimeExt returns the extension for the mapped mime type or string empty if not found
func GetMimeExt(m *MIME) string {
	if m != nil {
		key := m.String()
		for ext, mimeKeys := range RegExts {
			for _, regMime := range mimeKeys {
				if strings.EqualFold(key, regMime) {
					return ext
				}
			}
		}
	}
	return ""
}

// GetMimeExtOrDef returns the extension for the mapped mime type or the fallback value if not found
func GetMimeExtOrDef(m *MIME, def string) string {
	if m != nil {
		if ext := GetMimeExt(m); ext != "" {
			return ext
		}
	}
	return def
}

func GetStringExt(s string) string {
	if m, ok := RegMIME[s]; ok {
		return GetMimeExt(m)
	}
	return ""
}
