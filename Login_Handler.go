package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func authentication(reqBody USER) bool {

	var count int
	result := false

	checkSql := "SELECT COUNT(*) FROM signup_detail WHERE email = $1 AND password = $2"

	row := DB.QueryRow(checkSql, reqBody.Email, reqBody.Password)

	err := row.Scan(&count)
	fmt.Println("count", count)

	if err != nil {
		log.Fatal(err)
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

	reqBody := USER{}

	err := c.Bind(&reqBody)

	if err != nil {
		res := gin.H{
			"error": err.Error(),
		}

		c.JSON(http.StatusBadRequest, res)
		return
	}

	// fmt.Println(true)

	result := authentication(reqBody)

	if result == true {

		user_data := getUserByEmail(reqBody.Email)

		c.SetCookie("id", strconv.Itoa(user_data.Id), time.Now().Hour(), "", "", true, true)
		c.Header("result", "5455")
		res := gin.H{
			"success": true,
			"message": "sucessfully login",
			"test":    user_data,
		}
		c.JSON(http.StatusOK, res)
		return
	}

}
