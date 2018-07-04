package fc_log

type Level int
type Color int

const (
	LV_TRACE Level = iota
	LV_DEBUG
	LV_INFO
	LV_WARN
	LV_ERROR
	LV_FATAL
)

const (
	NoColor Color = iota
	Black
	Red
	Green
	Yellow
	Blue
	Purple
	DarkGreen
	White
)

var logColorPrefix = map[Color]string{
	NoColor:   "",
	Black:     "\x1b[030m",
	Red:       "\x1b[031m",
	Green:     "\x1b[032m",
	Yellow:    "\x1b[033m",
	Blue:      "\x1b[034m",
	Purple:    "\x1b[035m",
	DarkGreen: "\x1b[036m",
	White:     "\x1b[037m",
}

var logColorSuffix = "\x1b[0m"

var lvStr = map[Level]string{
	LV_TRACE: "[TRACE]",
	LV_DEBUG: "[DEBUG]",
	LV_INFO:  "[INFO]",
	LV_WARN:  "[WARN]",
	LV_ERROR: "[ERROR]",
	LV_FATAL: "[FATAL]",
}

var lvColor = map[Level]Color{
	LV_TRACE: NoColor,
	LV_DEBUG: Green,
	LV_INFO:  Blue,
	LV_WARN:  Yellow,
	LV_ERROR: Purple,
	LV_FATAL: Red,
}

func getColorPrefixByLv(lv Level) string {
	c := lvColor[lv]
	s := logColorPrefix[c]
	return s
}

func getFlagStrByLv(lv Level) string {
	return lvStr[lv]
}
