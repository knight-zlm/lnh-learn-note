package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

// FileLogger FileLogger
type FileLogger struct {
	Level       LogLevel
	FilePath    string
	FileName    string
	MaxFileSize int64
	outObj      *os.File
	errObj      *os.File
	// fullPath    string
}

// NewFileLogger NewFileLogger
func NewFileLogger(sl, fp, fn string, maxSize int64) Logger {
	le, _ := parseLogLevel(sl)
	fullpath := path.Join(fp, fn)
	fObj, err := os.OpenFile(fullpath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open file error:%s\n", err.Error())
		panic(err)
	}
	errFullpath := fmt.Sprintf("%s.err", fullpath)
	eObj, err := os.OpenFile(errFullpath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	return &FileLogger{
		Level:       le,
		FilePath:    fp,
		FileName:    fn,
		MaxFileSize: maxSize,
		outObj:      fObj,
		errObj:      eObj,
	}
}

func (l *FileLogger) close() {
	l.errObj.Close()
	l.outObj.Close()
}

func (l *FileLogger) enable(le LogLevel) bool {
	return le >= l.Level
}

// 检查文件大小切割文件
func (l *FileLogger) checkSize(fObj *os.File) *os.File {
	fInfo, err := fObj.Stat()
	if err != nil {
		fmt.Printf("检查文件出问题了 error:%s, fObj:%p", err.Error(), fObj)
		panic(err)
	}
	if fInfo.Size() < l.MaxFileSize {
		return fObj
	}
	// 关闭旧的文件打开新的文件
	now := time.Now()
	curFullPaht := path.Join(l.FilePath, fObj.Name())
	newFullPath := fmt.Sprintf("%s.bak%s", curFullPaht, now.Format("20060102150405000"))
	fObj.Close()
	os.Rename(curFullPaht, newFullPath)
	fObj, err = os.OpenFile(curFullPaht, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("打开文件出问题了 error:%s", err.Error())
		panic(err)
	}
	return fObj
}

// fileLog
func (l *FileLogger) fileLog(lv LogLevel, msg string) {
	data := makeLog(lv, msg)
	if lv >= ERROR {
		// 检查文件大小
		l.errObj = l.checkSize(l.errObj)
		fmt.Fprintf(l.errObj, data)
		// l.errObj.Close()
	}
	// 检查文件大小
	l.outObj = l.checkSize(l.outObj)
	fmt.Fprintf(l.outObj, data)
	// l.outObj.Close()
}

// Debug init
func (l *FileLogger) Debug(msg string) {
	if l.enable(DEBUG) {
		l.fileLog(DEBUG, msg)
	}
}

// Info Info
func (l *FileLogger) Info(msg string) {
	if l.enable(INFO) {
		l.fileLog(INFO, msg)
	}
}

// Warning Warning
func (l *FileLogger) Warning(msg string) {
	if l.enable(WARNING) {
		l.fileLog(WARNING, msg)
	}
}

// Error Error
func (l *FileLogger) Error(msg string) {
	if l.enable(ERROR) {
		l.fileLog(ERROR, msg)
	}
}
