// Package api
package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	Logger *logrus.Logger

	DBInsertSuccess string = "Created"

	Unauthorize          string = "Unauthorized"
	AccountNotExist      string = "Account not found"
	PermissionDenied     string = "Permission denied"
	RequestBodyError     string = "Request body error"
	ProcessError         string = "Internal process error"
	AccountAlreadyExists string = "Account already exists"
	DBQueryError         string = "DB query error"
	DBInsertError        string = "DB insert error"
	DBUpdateError        string = "DB update error"
)

// LogOut - Define log output
func LoggerOut() gin.HandlerFunc {
	return func(c *gin.Context) {
		// fmt.Println("I'm logger")
		startTime := time.Now()

		c.Next()

		// fmt.Println("I'm logger end")
		// fmt.Println(c.Writer.Header())

		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		reqMethod := c.Request.Method
		reqUri := c.Request.RequestURI
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()

		Logger.Infof("| %3d | %13v | %15s | %s | %s |",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
		)
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// fmt.Println("I'm cors.")
		// fmt.Println(c.Writer.Header())

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		// fmt.Println(c.Writer.Header())
		if c.Request.Method == "OPTIONS" {
			// c.AbortWithStatus(http.StatusNoContent)
			c.JSON(http.StatusNoContent, gin.H{
				"msg": "Request options",
			})
			return
		}

		c.Next()
		// fmt.Println("I'm cors end")
		// fmt.Println(c.Writer.Header())
	}
}

// BuildRouter - Build router with gin.Engine.group to build router tree
func BuildRouter() *gin.Engine {
	// apiv1PathElements := []string{viper.GetString("router.urlPath"), viper.GetString("router.v1")}
	// apiv1Path := "/" + strings.Join(apiv1PathElements[:], "/")

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	// Allow Cross Origin
	r.Use(CORSMiddleware())
	r.Use(LoggerOut())
	configAuthRouter("/auth", r)
	configUserRouter("/user", r)
	return r
}
