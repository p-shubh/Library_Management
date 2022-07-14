package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var live_user = make(map[string]USER) //present user

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

		id, err := strconv.Atoi(ID_cookie)
		if err != nil {
			res := gin.H{
				"status": "access denied",
				"":       "couldn't able to fetch the id",
			}
			c.JSON(http.StatusBadRequest, res)
		}

		data := getUserByid(id)
		fmt.Println(data)

		live_user["presentuser"] = data
		fmt.Println(live_user)
		fmt.Println("id =", live_user["presentuser"].Id)
		fmt.Println("firstname =", live_user["presentuser"].First_name)
		fmt.Println("lastname =", live_user["presentuser"].Last_name)
		fmt.Println("email =", live_user["presentuser"].Email)
		fmt.Println("passwrd =", live_user["presentuser"].Password)

	}
}
