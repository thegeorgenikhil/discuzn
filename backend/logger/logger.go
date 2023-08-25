package logger

import (
	"path"
	"runtime"
	"strings"

	log "github.com/sirupsen/logrus"
)

func InitLogger() {
	log.SetLevel(log.DebugLevel)
	log.SetReportCaller(true)

	log.SetFormatter(&log.TextFormatter{
		ForceColors:     true,
		DisableColors:   false,
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			s := strings.Split(f.Function, ".")
			funcname := s[len(s)-1]
			_, filename := path.Split(f.File)
			return "[" + funcname + "()]", "[" + filename + "]"
		},
	})

	log.Println("Logger initialized")
}
