package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func studentsHistory(c *gin.Context) {

	id, _ := c.Params.Get("id") // get method

	reqBody2 := student_order_history{}

	sqlStatatement := `SELECT id,book_id,issue_date,return_date,student_return_date,approve_grant,order_id,total_fine FROM students_order_detail where id = $1`

	row := DB.QueryRow(sqlStatatement, id)

	err := row.Scan(&reqBody2.Id, &reqBody2.Book_id, &reqBody2.Issue_date, &reqBody2.Return_date, &reqBody2.Student_Return_date, &reqBody2.Approve_grant, &reqBody2.Order_ID, &reqBody2.Fine)

	fmt.Println(err)

	if err != nil {
		//log.Fatal("ppt", err)
		res := gin.H{
			"error":  "unable to view order",
			"status": "check in get history from id",
		}
		c.JSON(http.StatusBadRequest, res)
		// c.Abort()
		// return

	} else {

		res := gin.H{
			"status":                 "success",
			"student history detail": reqBody2,
		}
		c.JSON(http.StatusOK, res)

	}

}
