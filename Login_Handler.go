package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func authentication(reqBodycheck USER) bool {

	var count int
	result := false

	checkSql := "SELECT COUNT(*) FROM signup_detail WHERE email = $1 AND password = $2"

	row := DB.QueryRow(checkSql, reqBodycheck.Email, reqBodycheck.Password)

	err := row.Scan(&count)

	if err != nil {
		// log.Fatal(err)
	}

	if count == 1 {
		result = true
	} else {
		result = false
	}
	return result
}

func LoginPostHandler(c *gin.Context) {

	// w := gin.New()

	reqBodycheck := USER{}

	err := c.Bind(&reqBodycheck)

	if err == nil {
		res := gin.H{
			"error": err.Error(),
		}
		//c.Writer.Header().Set("Content-Type", "application/json")

		c.JSON(http.StatusBadRequest, res)
		return
	}

	// fmt.Println(true)

	result := authentication(reqBodycheck)

	if result == true {
		res := gin.H{
			"success": true,
			"message": "sucessfully login",
		}
		c.JSON(http.StatusOK, res)
		return
	} else {
		res := gin.H{
			"success": false,
			"message": "Invalid Credential",
		}
		c.JSON(http.StatusOK, res)
		return
	}

}
