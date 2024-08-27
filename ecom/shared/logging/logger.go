package logging

import (
	"ecom/shared/utils/config_utils"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"path"
	"runtime"
	"sort"
	"strings"
	"time"
)

var Log *logrus.Logger

var logFormatter = &logrus.TextFormatter{
	PadLevelText:     true,
	TimestampFormat:  time.RFC1123Z,
	FullTimestamp:    true,
	CallerPrettyfier: caller(),
	FieldMap: logrus.FieldMap{
		logrus.FieldKeyTime:  "DateTime",
		logrus.FieldKeyLevel: "Level",
		logrus.FieldKeyFile:  "File",
		logrus.FieldKeyFunc:  "Line",
		logrus.FieldKeyMsg:   "Message",
	},
	SortingFunc: func(fields []string) {
		fmt.Println()
		sort.Slice(fields, func(i, j int) bool {
			// fmt.Println(fields[i])
			return fields[i] < fields[j]
		})
	},
}

// caller gets the function, file name and time w.r.t the instance of logger that has been called.
func caller() func(*runtime.Frame) (function string, file string) {
	// using runtime.Frame to get the function, file name, format time.
	return func(f *runtime.Frame) (function string, file string) {
		p, _ := os.Getwd()
		functionName := f.Function
		fileName := fmt.Sprintf("%s:%d", strings.TrimPrefix(f.File, p), f.Line)
		// dt := time.Now()
		return fileName + " ", functionName
	}
}

func Initializelogger() {
	wd := config_utils.LoadEnv()
	Log = logrus.New()
	Log.SetReportCaller(true)
	Log.SetFormatter(logFormatter)
	file, err := os.OpenFile(path.Join(wd+os.Getenv("LOG_FILE_LOCATION")), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		Log.Warn("Failed to log to the file..", err)
	} else {
		writers := []io.Writer{}
		writers = append(writers, os.Stdout, file)
		multiWriter := io.MultiWriter(writers...)
		Log.SetOutput(multiWriter)
	}
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		startTime := time.Now()
		next.ServeHTTP(res, req)
		Log.WithFields(logrus.Fields{
			"method":     req.Method,
			"path":       req.URL.Path,
			"remote":     req.RemoteAddr,
			"user-agent": req.UserAgent(),
			"status":     res.Header().Get("Status"),
			"duration":   time.Since(startTime),
		}).Info("Handled request")
	})
}
