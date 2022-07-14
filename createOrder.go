package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func createOrder(c *gin.Context) {

	reqBody2 := ORDER{}
	c.Bind(&reqBody2)

	fmt.Println("create order request body", reqBody2)

	reqBody2.Id = live_user["presentuser"].Id

	fmt.Println("create order id", reqBody2.Id)

	fmt.Println("book_id", reqBody2.Book_id)

	sqlStatement2 := "SELECT book_title,book_author,book_cover_image FROM books_detail where book_id= $1"

	row2 := DB.QueryRow(sqlStatement2, reqBody2.Book_id)

	err2 := row2.Scan(&reqBody2.Book_title, &reqBody2.Book_author, &reqBody2.Book_cover_image)

	// fmt.Println(err2)

	if err2 != nil {
		res := gin.H{
			"err": err2.Error(),
		}

		c.JSON(http.StatusBadRequest, res)
	} else {
		res := gin.H{
			"result":       reqBody2,
			"order status": "successfully orderd",
		}

		c.JSON(http.StatusOK, res)
	}

	fmt.Println(reqBody2)

	fmt.Println("reqBody2.Issue_date :", reqBody2.Issue_date)
	fmt.Println("reqBody2.return_date :", reqBody2.Return_date)

	fmt.Println("calculate time =", calculateTime(reqBody2.Issue_date, reqBody2.Return_date))

}

func calculateTime(issue_date string, return_date string) int {

	// dateString := "2021-11-22"
	issuedate, _ := time.Parse("2006-01-02", issue_date)

	returndate, _ := time.Parse("2006-01-02", return_date)

	total_duration := returndate.Sub(issuedate)

	fmt.Println(total_duration)

	d, _ := time.ParseDuration(total_duration.String())
	days := d.Hours() / 24 // 2 days

	fmt.Println(days)

	return int(days)

}
