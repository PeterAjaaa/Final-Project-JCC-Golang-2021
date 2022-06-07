package middleware

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv"
)

func SuperUsersAuth() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		os.Getenv("USERNAME"):  os.Getenv("PASSWORD"),
		os.Getenv("USERNAME2"): os.Getenv("PASSWORD2"),
		os.Getenv("USERNAME3"): os.Getenv("PASSWORD3"),
	})
}

func SetupLogOutput() {
	file, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(file)
}

func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] %s %s %d %s",
			params.ClientIP,
			params.TimeStamp.Format(time.RFC1123Z),
			params.Method,
			params.Path,
			params.StatusCode,
			params.Latency)
	})
}
