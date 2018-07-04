package fc_log

import (
	_ "fmt"
)

func fmtColorBegin(l *Logger) {
	addFmtStr(l, getColorPrefixByLv(l.curLv))
}

func fmtColorEnd(l *Logger) {
	addFmtStr(l, logColorSuffix)
}

func fmtLv(l *Logger) {
	addFmtStr(l, getFlagStrByLv(l.curLv))
	addFmtByte(l, ' ')
}

func fmtSrcStr(l *Logger) {
	addFmtStr(l, l.srcStr)
}

func fmtStrLine(l *Logger) {
	strLen := len(l.srcStr)
	if (strLen > 0 && l.srcStr[strLen-1] != '\n') || strLen == 0 {
		addFmtByte(l, '\n')
	}
}

func fmtStr(l *Logger) {
	fmtColorBegin(l)
	fmtLv(l)
	fmtSrcStr(l)
	fmtColorEnd(l)
	fmtStrLine(l)
}

func addFmtStr(l *Logger, s string) {
	l.buf = append(l.buf, s...)
}

func addFmtByte(l *Logger, b byte) {
	l.buf = append(l.buf, b)
}
