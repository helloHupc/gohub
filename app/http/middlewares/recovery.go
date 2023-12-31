package middlewares

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gohub/pkg/logger"
	"gohub/pkg/response"
	"net"
	"net/http/httputil"
	"os"
	"strings"
	"time"
)

func Recovery() gin.HandlerFunc {

	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 获取用户的请求信息
				httpRequest, _ := httputil.DumpRequest(c.Request, true)

				// 链接中断
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						errStr := strings.ToLower(se.Error())
						if strings.Contains(errStr, "brokenPipe") || strings.Contains(errStr, "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				if brokenPipe {
					logger.Error(c.Request.URL.Path,
						zap.Time("time", time.Now()),
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					c.Error(err.(error))
					c.Abort()

					return
				}

				// 如果不是链接断开，就开始记录堆栈信息
				logger.Error("recovery from panic",
					zap.Time("time", time.Now()),
					zap.Any("error", err),
					zap.String("request", string(httpRequest)), // 请求信息
					zap.Stack("stacktrace"),
				)

				// 返回500状态码
				response.Abort500(c)
			}
		}()
		c.Next()
	}
}
