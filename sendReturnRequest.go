package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func returnRequest(c *gin.Context) {

	reqBody := student_Return_Detail{}

	err2 := c.Bind(&reqBody)

	var count int

	err3 := DB.QueryRow(`select count(*) from students_order_detail where order_id = $1;`, reqBody.Order_ID).Scan(&count)

	// checkOrderidSQL := (`select count(*) from students_order_detail where order_id = $1;`).Scan(&count)

	switch {
	case err3 != nil:
		log.Fatal(err3)
	default:
		fmt.Println("Number of rows are", count)
	}

	// check order id

	if err2 != nil {
		res := gin.H{
			"warning": "return date or order id can't be empty",
			"status":  reqBody,
		}

		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	}

	if count == 0 {
		res := gin.H{
			"status": "order_id is incorrect",
		}
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return

	} else {
		res := gin.H{
			"status": "order id is correct",
		}
		c.JSON(http.StatusOK, res)

	}

	return_request := "pending"

	updateOrderSQL := `UPDATE students_order_detail SET student_return_date=$1, return_request=$2 WHERE order_id = $3;`

	_, err := DB.Exec(updateOrderSQL, reqBody.Student_Return_date, return_request, reqBody.Order_ID)

	if err != nil {

		res := gin.H{
			"message": "update student return date gone wrong",
			"result":  updateOrderSQL,
		}
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	} else {
		res := gin.H{
			"message": "sent return requestt sucessfully",
			// "result":  updateOrderSQL,
		}
		c.JSON(http.StatusOK, res)
		// c.Abort()
		// return
	}
}
