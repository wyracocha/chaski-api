package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.Info("starting web server")
	r := gin.Default()
	//var payload domain.Payload
	r.GET("/", func(c *gin.Context) {
		payloadString := c.GetHeader("Payload")
		c.JSON(http.StatusOK, gin.H{"body": payloadString})
		return
	})
	r.Run(fmt.Sprintf("0.0.0.0:%s", os.Getenv("PORT")))
}
