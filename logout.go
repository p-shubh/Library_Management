package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func logout(c *gin.Context) {

	c.SetCookie("id", "", -1, "", "", true, true)

	res := gin.H{
		"success": true,
		"status":  "logout succeessfully",
	}
	c.JSON(http.StatusOK, res)
}
