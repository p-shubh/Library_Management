package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func logout(c *gin.Context) {

	ID_cookie, err := c.Cookie("id")

	c.SetCookie("id", "", -1, "", "", true, true)

	// ID_cookie, err := c.Cookie("id")

	id, err := strconv.Atoi(ID_cookie)
	if err != nil {
		res := gin.H{
			"status":  "access denied",
			"warning": "couldn't able to fetch the id",
		}
		c.JSON(http.StatusBadRequest, res)
		c.Abort() // change 1
		return

	}

	Data := getUserByid(id)

	if Data.User_type == 1 {
		res := gin.H{
			"success": true,
			"status":  "admin logout succeessfully",
		}
		c.JSON(http.StatusOK, res)
		c.Abort()
	} else if Data.User_type == 2 {
		res := gin.H{
			"success": true,
			"status":  "student logout succeessfully",
		}
		c.JSON(http.StatusOK, res)
		c.Abort()
	}
}
