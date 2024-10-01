package api

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"runtime"
	"sing-song/pkg/logging"
)

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE, OPTIONS, PUT, HEAD, TRACE, CONNECT")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Expose-Headers", "*")
		if r.Method == "OPTIONS" {
			w.Write([]byte("OPTIONS"))
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
func RecoverAllPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			logger := logging.GetLogger()
			if err := recover(); err != nil {
				pc, file, line, _ := runtime.Caller(4)
				funcName := runtime.FuncForPC(pc).Name()
				logger.WithFields(logrus.Fields{
					"panic": err,
					"file":  file,
					"line":  line,
					"func":  funcName,
				}).Error("Паника была обработана")
				http.Error(w, "Серверная ошибка", 500)
			}
			return
		}()
		next.ServeHTTP(w, r)
	})
}
