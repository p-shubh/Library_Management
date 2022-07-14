package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func isLogin() gin.HandlerFunc {

	return func(c *gin.Context) {

		ID_cookie, err := c.Cookie("id")

		if err != nil {
			res := gin.H{
				"message":   "cheaking code is running or not",
				"my-cookie": ID_cookie,
			}
			c.JSON(http.StatusBadRequest, res)
			c.Abort()
			return
		}
		fmt.Println("my-cookie", ID_cookie)

	}
}

// func isLogin() gin.HandlerFunc {

// 	return func(c *gin.Context) {
// 		if cookie, err := c.Request.Cookie("test_session"); err == nil {

// 		}

// 	}
// }
