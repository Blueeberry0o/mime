package mime

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

// MIME Register
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

	"application/x-7z-compressed": MustParse("application/x-7z-compressed"),

	"application/vnd.ms-excel":            MustParse("application/vnd.ms-excel"),
	"application/vnd.amazon.ebook":        MustParse("application/vnd.amazon.ebook"),
	"application/vnd.ms-fontobject":       MustParse("application/vnd.ms-fontobject"),
	"application/vnd.ms-powerpoint":       MustParse("application/vnd.ms-powerpoint"),
	"application/vnd.mozilla.xul+xml":     MustParse("application/vnd.mozilla.xul+xml"),
	"application/vnd.apple.installer+xml": MustParse("application/vnd.apple.installer+xml"),

	"application/vnd.oasis.opendocument.text":         MustParse("application/vnd.oasis.opendocument.text"),
	"application/vnd.oasis.opendocument.spreadsheet":  MustParse("application/vnd.oasis.opendocument.spreadsheet"),
	"application/vnd.oasis.opendocument.presentation": MustParse("application/vnd.oasis.opendocument.presentation"),

	"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet":         MustParse("application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"),
	"application/vnd.openxmlformats-officedocument.wordprocessingml.document":   MustParse("application/vnd.openxmlformats-officedocument.wordprocessingml.document"),
	"application/vnd.openxmlformats-officedocument.presentationml.presentation": MustParse("application/vnd.openxmlformats-officedocument.presentationml.presentation"),
}

// Extension register / mapping
var RegExtension = map[string][]*MIME{
	".3gp":    {RegMIME["audio/3gpp"], RegMIME["video/3gpp"]},
	".3g2":    {RegMIME["audio/3gpp2"], RegMIME["video/3gpp2"]},
	".7z":     {RegMIME["application/x-7z-compressed"]},
	".aac":    {RegMIME["audio/aac"]},
	".abw":    {RegMIME["application/x-abiword"]},
	".arc":    {RegMIME["application/x-freearc"]},
	".avif":   {RegMIME["image/avif"]},
	".avi":    {RegMIME["video/x-msvideo"]},
	".azw":    {RegMIME["application/vnd.amazon.ebook"]},
	".bin":    {RegMIME["application/octet-stream"]},
	".bmp":    {RegMIME["image/bmp"]},
	".bz":     {RegMIME["application/x-bzip"]},
	".bz2":    {RegMIME["application/x-bzip2"]},
	".cda":    {RegMIME["application/x-cdf"]},
	".csh":    {RegMIME["application/x-csh"]},
	".css":    {RegMIME["text/css"]},
	".csv":    {RegMIME["text/csv"]},
	".doc":    {RegMIME["application/msword"]},
	".docx":   {RegMIME["application/vnd.openxmlformats-officedocument.wordprocessingml.document"]},
	".eot":    {RegMIME["application/vnd.ms-fontobject"]},
	".epub":   {RegMIME["application/epub+zip"]},
	".gz":     {RegMIME["application/gzip"]},
	".gif":    {RegMIME["image/gif"]},
	".htm":    {RegMIME["text/html"]},
	".html":   {RegMIME["text/html"]},
	".ico":    {RegMIME["image/vnd.microsoft.icon"]},
	".ics":    {RegMIME["text/calendar"]},
	".jar":    {RegMIME["application/java-archive"]},
	".jpeg":   {RegMIME["image/jpeg"]},
	".jpg":    {RegMIME["image/jpeg"]},
	".json":   {RegMIME["application/json"]},
	".jsonld": {RegMIME["application/ld+json"]},
	".mid":    {RegMIME["audio/midi"], RegMIME["audio/x-midi"]},
	".midi":   {RegMIME["audio/midi"], RegMIME["audio/x-midi"]},
	".mjs":    {RegMIME["text/javascript"]},
	".mp3":    {RegMIME["audio/mpeg"]},
	".mp4":    {RegMIME["video/mp4"]},
	".mpeg":   {RegMIME["video/mpeg"]},
	".mpkg":   {RegMIME["application/vnd.apple.installer+xml"]},
	".odp":    {RegMIME["application/vnd.oasis.opendocument.presentation"]},
	".ods":    {RegMIME["application/vnd.oasis.opendocument.spreadsheet"]},
	".odt":    {RegMIME["application/vnd.oasis.opendocument.text"]},
	".oga":    {RegMIME["audio/ogg"]},
	".ogv":    {RegMIME["video/ogg"]},
	".ogx":    {RegMIME["application/ogg"]},
	".opus":   {RegMIME["audio/opus"]},
	".otf":    {RegMIME["font/otf"]},
	".png":    {RegMIME["image/png"]},
	".pdf":    {RegMIME["application/pdf"]},
	".php":    {RegMIME["application/x-httpd-php"]},
	".ppt":    {RegMIME["application/vnd.ms-powerpoint"]},
	".pptx":   {RegMIME["application/vnd.openxmlformats-officedocument.presentationml.presentation"]},
	".rar":    {RegMIME["application/vnd.rar"]},
	".rtf":    {RegMIME["application/rtf"]},
	".sh":     {RegMIME["application/x-sh"]},
	".svg":    {RegMIME["image/svg+xml"]},
	".tar":    {RegMIME["application/x-tar"]},
	".tif":    {RegMIME["image/tiff"]},
	".tiff":   {RegMIME["image/tiff"]},
	".ts":     {RegMIME["video/mp2t"]},
	".ttf":    {RegMIME["font/ttf"]},
	".txt":    {RegMIME["text/plain"]},
	".vsd":    {RegMIME["application/vnd.visio"]},
	".wav":    {RegMIME["audio/wav"]},
	".weba":   {RegMIME["audio/webm"]},
	".webm":   {RegMIME["video/webm"]},
	".webp":   {RegMIME["image/webp"]},
	".woff":   {RegMIME["font/woff"]},
	".woff2":  {RegMIME["font/woff2"]},
	".xhtml":  {RegMIME["application/xhtml+xml"]},
	".xls":    {RegMIME["application/vnd.ms-excel"]},
	".xlsx":   {RegMIME["application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"]},
	".xml":    {RegMIME["application/xml"]},
	".xul":    {RegMIME["application/vnd.mozilla.xul+xml"]},
	".zip":    {RegMIME["application/zip"]},
}
