package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func SignUpPostHandler(c *gin.Context) {

	// calling the struct
	reqBody := USER{}

	err := c.Bind(&reqBody) //binding with the data being provided

	// creating response

	if err != nil {
		res := gin.H{
			"error": err.Error(),
			"req":   reqBody,
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	if reqBody.Email == "" {
		res := gin.H{
			"error": "email id must not be empty",
		}
		c.JSON(http.StatusBadRequest, res)
	}

	if reqBody.Password == "" {
		res := gin.H{
			"error": "Password must not be empty",
		}
		c.JSON(http.StatusBadRequest, res)
	}

	// if reqBody.Id   {			//concern
	// 	res := gin.H{
	// 		"error": "Id must not be empty",
	// 	}
	// 	c.JSON(http.StatusBadRequest, res)
	// }

	var user_type int

	if c.Request.URL.Path == "/signup/admin" {
		user_type = 1
	} else {
		user_type = 2
	}

	Result, err2 := InsertSignUpDB(reqBody, user_type)

	if err2 != "" {
		res := gin.H{
			"error":   err.Error(),
			"success": Result,
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := gin.H{
		"success": true,
		"result":  Result,
		"message": "Registered Successfull",
	}

	c.JSON(http.StatusOK, res)

}

func InsertSignUpDB(reqBody USER, user_type int) (bool, string) {

	var Result = true

	var err_respo = ""

	sqlStatement := `INSERT INTO signup_detail(id,lastname,firstname,email,user_type,password)values($1,$2,$3,$4,$5,$6)`

	_, err := DB.Exec(sqlStatement, reqBody.Id, reqBody.Last_name, reqBody.First_name, reqBody.Email, user_type, reqBody.Password)

	if err != nil {

		err2 := UniqueViolation(err)

		if err2 != nil {
			return false, err2.Detail
		}
		err_respo = "Something went wrong"

		Result = false

		return Result, err_respo

	}

	return Result, err_respo
}

func UniqueViolation(err error) *pq.Error {
	if pqerr, ok := err.(*pq.Error); ok &&
		pqerr.Code == "23505" {
		return pqerr
	}
	return nil
}
