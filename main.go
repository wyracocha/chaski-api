package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	r.Use(gin.Recovery())
	//var payload domain.Payload
	r.Any("/", func(c *gin.Context) {
		payloadString := c.GetHeader("Payload")
		c.Header("Powered-By", "Chaski")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET,POST")
		c.Header("Access-Control-Allow-Headers", "Content-Type,Payload")
		c.JSON(http.StatusOK, gin.H{"body": payloadString})
		return
	})
	r.Run(fmt.Sprintf("0.0.0.0:%s", os.Getenv("PORT")))
}
