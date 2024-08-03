package main

import (
	"github.com/gin-gonic/gin"
	"github.com/savabush/restApiTest/internal/config"
	"log"
	"net/http"
)

func main() {
	ginSettings := config.Settings.GIN
	r := gin.Default()

	expectedHost := ginSettings.Host

	r.Use(func(c *gin.Context) {
		if c.Request.Host != expectedHost {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid host header"})
			return
		}
		c.Header("X-Frame-Options", "DENY")
		c.Header("Content-Security-Policy", "default-src 'self'; connect-src *; font-src *; script-src-elem * 'unsafe-inline'; img-src * data:; style-src * 'unsafe-inline';")
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
		c.Header("Referrer-Policy", "strict-origin")
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("Permissions-Policy", "geolocation=(),midi=(),sync-xhr=(),microphone=(),camera=(),magnetometer=(),gyroscope=(),fullscreen=(self),payment=()")
		c.Next()
	})

	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "hello"})
		})
		v1.GET("/someDataFromReader", func(c *gin.Context) {
			response, err := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")
			if err != nil || response.StatusCode != http.StatusOK {
				c.Status(http.StatusServiceUnavailable)
				return
			}

			reader := response.Body
			contentLength := response.ContentLength
			contentType := response.Header.Get("Content-Type")

			extraHeaders := map[string]string{
				"Content-Disposition": `attachment; filename="gopher.png"`,
			}

			c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
		})

	}

	err := r.Run(ginSettings.Host)
	if err != nil {
		log.Fatal(err)
	}
}
