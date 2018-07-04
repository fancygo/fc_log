package fc_log

import (
	"os"
)

type LogFile struct {
	file   *os.File
	nextCh chan *os.File
}

func NewLogFile(f *os.File) *LogFile {
	return &LogFile{
		file:   f,
		nextCh: make(chan *os.File),
	}
}

func (l *LogFile) Write(b []byte) (int, error) {
	select {
	case newFile := <-l.nextCh:
		old := l.file
		l.file = newFile
		old.Close()
	default:
	}
	n, err := l.file.Write(b)
	return n, err
}

func (l *LogFile) NextFile(f *os.File) {
	l.nextCh <- f
}
