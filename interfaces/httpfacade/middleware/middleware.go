package middleware

import (
	"bytes"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//PingHealth ping 服务是否正常
func PingHealth(c *gin.Context) {
	c.String(http.StatusOK, "success")
	return
}

//NotFindRouter 未找到相关路由
func NotFindRouter(c *gin.Context) {
	c.String(http.StatusNotFound, "The incorrect API route.")
	return
}

// GinLogWriter 自定义logger
type GinLogWriter struct {
	gin.ResponseWriter
	bodyBuf *bytes.Buffer
}

// Write implament gin.ResponseWriter
func (w *GinLogWriter) Write(b []byte) (int, error) {
	w.bodyBuf.Write(b)
	return w.ResponseWriter.Write(b)
}

// WriteString implament gin.ResponseWriter
func (w *GinLogWriter) WriteString(s string) (int, error) {
	w.bodyBuf.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

// GinQscLogger 自定义gin标准日志输出
func GinQscLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		body, _ := c.GetRawData()
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		newWriter := &GinLogWriter{
			bodyBuf:        bytes.NewBufferString(""),
			ResponseWriter: c.Writer,
		}
		c.Writer = newWriter

		log.Println("http request logger",
			zap.String("path", c.Request.URL.Path),
			zap.String("query", c.Request.URL.RawQuery),
			zap.String("client_ip", c.ClientIP()),
			zap.String("UserAgent", c.Request.UserAgent()),
			zap.ByteString("body", body),
		)

		c.Next()

		end := time.Now()
		cost := end.Sub(start)

		responseBody := "非标准返回，过滤返回内容"
		ct := c.Writer.Header().Get("content-type")
		if strings.Contains(strings.ToLower(ct), "application/json") {
			responseBody = newWriter.bodyBuf.String()
		}

		log.Println("http response logger",
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()),
			zap.String("body", responseBody),
			zap.String("start", start.Format(time.RFC3339Nano)),
			zap.String("end", end.Format(time.RFC3339Nano)),
			zap.Duration("cost", cost),
		)
	}
}

// GinQscRecovery 自定义gin recover输出
func GinQscRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					log.Println("http request broken pipe",
						zap.String("path", c.Request.URL.Path),
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				log.Println("http request panic",
					zap.String("path", c.Request.URL.Path),
					zap.Any("error", err),
					zap.String("request", string(httpRequest)),
					zap.String("stack", string(debug.Stack())),
				)

				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
