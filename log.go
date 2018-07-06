package fc_log

import (
	"fmt"
	"github.com/FancyGo/fc_sys"
	"log"
	"os"
	"path"
	"time"
)

type Logger struct {
	buf        []byte
	name       string
	srcStr     string
	curLv      Level
	lv         Level
	outputMode int

	stdLogger *log.Logger

	logFlag      int
	fileLogger   *log.Logger
	filePtr      *os.File
	fileWriter   *LogFile
	fileIdx      int
	fileSplit    string
	fileInterval int
	fileSize     int
}

func NewLogger(name string, lv Level, mode int) (*Logger, error) {
	l := &Logger{
		name:         name,
		buf:          make([]byte, 0),
		lv:           lv,
		logFlag:      log.Lshortfile | log.LstdFlags | log.Lmicroseconds,
		outputMode:   mode,
		fileInterval: LOG_INTERVAL_DAY,
		fileSize:     LOG_DEF_MAX_SIZE,
	}
	if err := l.InitConsole(); err != nil {
		return l, err
	}
	if err := l.InitFile(); err != nil {
		return l, err
	}
	return l, nil
}

func (l *Logger) InitConsole() error {
	l.stdLogger = log.New(os.Stderr, "", l.logFlag)
	return nil
}

func (l *Logger) InitFile() error {
	var err error
	l.fileSplit = l.getSplitTag(time.Now())
	l.filePtr, err = l.findFile(-1)
	if err != nil {
		return err
	}
	l.fileWriter = NewLogFile(l.filePtr)
	l.fileLogger = log.New(l.fileWriter, "", l.logFlag)

	go l.fileCheck()
	return nil
}

func (l *Logger) fileCheck() {
	for {
		time.Sleep(time.Second * 1)
		curTm := time.Now()
		newSplit := l.getSplitTag(curTm)
		if newSplit == l.fileSplit {
			st, err := l.filePtr.Stat()
			if err != nil {
				l.Default("logger file get stat fail err = %v\n", err)
				continue
			}
			fsize := st.Size()
			l.Info("fileCheck size = %d, filesize = %d", fsize, l.fileSize)
			if fsize > int64(l.fileSize) {
				if l.fileIdx == LOG_MAX_FILE_IDX {
					l.Default("logger file num max = %d\n", l.fileIdx)
					continue
				} else {
					l.fileIdx++
					l.Default("new logger file idx = %d\n", l.fileIdx)
				}
			} else {
				continue
			}
		} else {
			l.fileSplit = newSplit
			l.fileIdx = 0
		}

		newFilePtr, err := l.findFile(l.fileIdx)
		if err != nil {
			l.Default("new logger file idx = %d, spilt = %v err\n", l.fileIdx, l.fileSplit)
			continue
		}
		l.filePtr = newFilePtr
		l.fileWriter.NextFile(l.filePtr)
		l.Default("new logger file\n")
	}
}

func (l *Logger) getSplitTag(tm time.Time) string {
	if l.fileInterval == LOG_INTERVAL_DAY {
		return tm.Format("0102")
	} else if l.fileInterval == LOG_INTERVAL_HOUR {
		return tm.Format("010215")
	} else if l.fileInterval == LOG_INTERVAL_MIN {
		return tm.Format("01021504")
	}
	return tm.Format("0102")
}

func (l *Logger) findFile(idx int) (*os.File, error) {
	findIdx := idx
	if idx < 0 {
		for i := LOG_MAX_FILE_IDX; i >= 1; i-- {
			fileName := fmt.Sprintf("%v.%v.%02d.log", l.name, l.fileSplit, i)
			filePath := path.Join(fc_sys.GetLogDir(), fileName)
			filePtr, err := os.OpenFile(filePath, os.O_RDONLY, os.ModePerm)
			if err != nil {
				continue
			}
			findIdx = i
			filePtr.Close()
			break
		}
		if findIdx == -1 {
			findIdx = 0
		}
	}
	fileName := fmt.Sprintf("%v.%v.%02d.log", l.name, l.fileSplit, findIdx)
	filePath := path.Join(fc_sys.GetLogDir(), fileName)
	filePtr, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		return nil, err
	}
	l.fileIdx = findIdx
	return filePtr, nil
}

func (l *Logger) Log(lv Level, str string) {
	l.srcStr = str
	l.curLv = lv

	if l.curLv < l.lv {
		return
	}

	l.buf = l.buf[:0]
	fmtStr(l)

	if l.checkOutputMode(LOG_OUTPUT_STD) {
		l.stdLogger.Output(3, string(l.buf))
	}
	if l.checkOutputMode(LOG_OUTPUT_FILE) {
		l.fileLogger.Output(3, string(l.srcStr))
	}
}

func (l *Logger) checkOutputMode(m int) bool {
	return (l.outputMode & m) > 0
}

func (l *Logger) Trace(str string, v ...interface{}) {
	l.Log(LV_TRACE, fmt.Sprintf(str, v...))
}

func (l *Logger) Debug(str string, v ...interface{}) {
	l.Log(LV_DEBUG, fmt.Sprintf(str, v...))
}

func (l *Logger) Info(str string, v ...interface{}) {
	l.Log(LV_INFO, fmt.Sprintf(str, v...))
}

func (l *Logger) Warn(str string, v ...interface{}) {
	l.Log(LV_WARN, fmt.Sprintf(str, v...))
}

func (l *Logger) Err(str string, v ...interface{}) {
	l.Log(LV_ERROR, fmt.Sprintf(str, v...))
}

func (l *Logger) Fatal(str string, v ...interface{}) {
	l.Log(LV_FATAL, fmt.Sprintf(str, v...))
}

func (l *Logger) Default(str string, v ...interface{}) {
	fmt.Printf(str, v...)
}

func Sys(str string, v ...interface{}) {
	fmt.Printf(str, v...)
}
func Sysln(v ...interface{}) {
	fmt.Println(v...)
}
