package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	r := gin.Default()

	//method GET
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Param("name")
		value := c.Query("value")
		dbValue, ok := db[user]

		if value != "" {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": dbValue})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": "no value"})
		}
	})

	//method POST
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar",
		"manu": "123",
	}))

	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})
	return r
}
func main() {
	r := setupRouter()
	r.Run(":8080")
}
