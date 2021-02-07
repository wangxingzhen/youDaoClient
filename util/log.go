package util

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	_dateFormat = "2006-01-02"   // 时间格式
	_fileName   = "%s/%s-%s.log" // 文件名
)

var (
	logs *logger // 日志
	c    *Config // 配置
)

var (
	_stdout bool
	_dir    string
)

type Config struct {
	Stdout bool
	Dir    string
}

func init() {
	addFlag(flag.CommandLine)
}

func addFlag(fs *flag.FlagSet) {
	fs.BoolVar(&_stdout, "log.stdout", false, "log enable stdout or not")
	fs.StringVar(&_dir, "log.dir", "", "log file path")
}

// 日志格式
type formatter struct{}

// 日志对象
type logInstance struct {
	logger *logrus.Logger
	date   *time.Time
	lock   *sync.Mutex
}

// 日志
type logger struct {
	debugLogger *logInstance // debug
	infoLogger  *logInstance // info
	errorLogger *logInstance // error
	fatalLogger *logInstance // fatal
}

// 文件名
func fileNameFormat(logDir, level, time string) string {
	return fmt.Sprintf(_fileName, logDir, level, time)
}

// 格式化
func (s *formatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := time.Now().Local().Format("2006/01/02 15:04:05")
	msg := fmt.Sprintf("[%s] %s %s\n", strings.ToUpper(entry.Level.String()), timestamp, entry.Message)
	return []byte(msg), nil
}

// 写入
func (log *logInstance) write(format string, args ...interface{}) {
	log.lock.Lock()
	if log.checkDate() {
		log.newOutFile()
	}
	log.lock.Unlock()
	str := fmt.Sprintf(format, args...)
	log.logger.Log(log.logger.Level, str)
	if c.Stdout {
		fmt.Println(str)
	}
}

// 时间判断
func (log *logInstance) checkDate() bool {
	now := time.Now().Format(_dateFormat)
	t, _ := time.Parse(_dateFormat, now)
	return t.After(*log.date)
}

// 日志文件
func (log *logInstance) newOutFile() {
	if log.logger.Out != nil {
		_ = log.logger.Out.(*os.File).Close()
	}
	now := time.Now().Format(_dateFormat)
	t, _ := time.Parse(_dateFormat, now)
	log.date = &t
	f, err := os.OpenFile(fileNameFormat(c.Dir, strings.ToUpper(log.logger.Level.String()), now), os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err == nil {
		log.logger.SetOutput(f)
	}
}

// info
func Info(format string, args ...interface{}) {
	logs.infoLogger.write(format, args...)
}

// debug
func Debug(format string, args ...interface{}) {
	logs.debugLogger.write(format, args...)
}

// fatal
func Fatal(format string, args ...interface{}) {
	logs.fatalLogger.write(format, args...)
}

// error
func Error(format string, args ...interface{}) {
	logs.errorLogger.write(format, args...)
}

// 初始化
func NewLogger(config *Config) (err error) {
	c = config
	if c == nil {
		if _dir == "" {
			return errors.New("no log path")
		}
		c = &Config{
			Stdout: _stdout,
			Dir:    _dir,
		}
	}
	err = createDir(c.Dir)
	if err != nil {
		return
	}
	logs = &logger{
		debugLogger: newLog(logrus.DebugLevel),
		infoLogger:  newLog(logrus.InfoLevel),
		errorLogger: newLog(logrus.ErrorLevel),
		fatalLogger: newLog(logrus.FatalLevel),
	}
	return
}

// 创建日志目录
func createDir(dir string) (err error) {
	_, err = os.Stat(dir)
	b := err == nil || os.IsExist(err)
	if !b {
		if err = os.MkdirAll(dir, os.ModePerm); err != nil {
			if os.IsPermission(err) {
				return
			}
		}
	}
	return
}

// 实例单个日志对象
func newLog(level logrus.Level) *logInstance {
	l := &logInstance{
		logger: &logrus.Logger{
			Formatter: &formatter{},
			Level:     level,
		},
		lock: new(sync.Mutex),
	}
	l.newOutFile()
	return l
}

// 关闭
func Close() {
	_ = logs.infoLogger.logger.Out.(*os.File).Close()
	_ = logs.errorLogger.logger.Out.(*os.File).Close()
	_ = logs.fatalLogger.logger.Out.(*os.File).Close()
	_ = logs.infoLogger.logger.Out.(*os.File).Close()
}
