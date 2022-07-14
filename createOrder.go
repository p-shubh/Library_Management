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

	var stock int

	sqlStatement2 := "SELECT book_title,book_author,book_cover_image,stock FROM books_detail where book_id= $1"

	row2 := DB.QueryRow(sqlStatement2, reqBody2.Book_id)

	err2 := row2.Scan(&reqBody2.Book_title, &reqBody2.Book_author, &reqBody2.Book_cover_image, &stock)

	fmt.Println("stock", stock)

	// fmt.Println(err2)

	// book stock check===========================================================================

	if stock == 0 {
		res := gin.H{
			"book status": "book is out of stock",
		}

		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return

	}
	// ===========================================================================================

	if err2 != nil {
		res := gin.H{
			"err":          err2.Error(),
			"order status": "sorry we don't have this book in the library",
		}

		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return

	} else {
		res := gin.H{
			"result":       reqBody2,
			"order status": "yes its available in book library",
		}

		c.JSON(http.StatusOK, res)

	}

	fmt.Println(reqBody2)

	fmt.Println("reqBody2.Issue_date :", reqBody2.Issue_date)
	fmt.Println("reqBody2.return_date :", reqBody2.Return_date)

	fmt.Println("calculate time =", calculateTime(reqBody2.Issue_date, reqBody2.Return_date))

	total_time := calculateTime(reqBody2.Issue_date, reqBody2.Return_date)

	if total_time <= 30 {
		res := gin.H{
			"status": "yes, its available for order",
		}
		c.JSON(http.StatusOK, res)
	} else if total_time > 30 {
		res := gin.H{
			"status": "its not available for more than 30 days",
		}
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return

	}

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
