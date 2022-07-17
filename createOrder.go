package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func OrderRequest(c *gin.Context) {

	//  ========================================================CHECK_ORDER==============================================================

	var scan_book_id string

	reqBody2 := ORDER{}
	c.Bind(&reqBody2)

	if reqBody2.Book_id == "" {
		res := gin.H{
			"err":    "book_id must not be empty",
			"result": reqBody2.Book_id,
		}
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	}

	if reqBody2.Issue_date == "" {
		res := gin.H{
			"err":    "Issue_date must not be empty",
			"result": reqBody2.Issue_date,
		}
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	}

	if reqBody2.Return_date == "" {
		res := gin.H{
			"err":    "Return_date must not be empty",
			"result": reqBody2.Return_date,
		}
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	}

	fmt.Println("create order request body", reqBody2)

	reqBody2.Id = live_user["presentuser"].Id

	fmt.Println("create order id", reqBody2.Id)

	fmt.Println("book_id", reqBody2.Book_id)

	var stock int

	sqlStatement2 := "SELECT book_title,book_author,book_cover_image,book_id,stock FROM books_detail where book_id= $1"

	row2 := DB.QueryRow(sqlStatement2, reqBody2.Book_id)

	err2 := row2.Scan(&reqBody2.Book_title, &reqBody2.Book_author, &reqBody2.Book_cover_image, &scan_book_id, &stock)

	fmt.Println("stock", stock)

	// fmt.Println(err2)

	// =====================================================book id match========================================================================

	if scan_book_id != reqBody2.Book_id {
		res := gin.H{
			"warning": "book id is incorrect",
		}
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	}

	// =====================================================book stock check===========================================================================

	if stock == 0 {
		res := gin.H{
			"book status": "book is out of stock",
		}

		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return

	}
	// ==============================================================================================================================================

	if err2 != nil {
		res := gin.H{
			"err":          err2.Error(),
			"order status": "sorry we don't have this book in the library",
			"result":       reqBody2,
		}

		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return

	} else {
		res := gin.H{
			"order status": "yes its available in book library",
			// "result":       reqBody2,
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

	// =====================================================INSERT_ORDER===========================================================================

	approve_grant := "pending"

	insertOrderSQL := `INSERT INTO students_order_detail(id,book_id,issue_date,return_date,approve_grant)values($1,$2,$3,$4,$5)`

	_, err := DB.Exec(insertOrderSQL, reqBody2.Id, reqBody2.Book_id, reqBody2.Issue_date, reqBody2.Return_date, approve_grant)

	if err != nil {

		res := gin.H{
			"message": "insertOrderSQL is not inserting",
			"result":  insertOrderSQL,
		}
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	} else {
		res := gin.H{
			// "status": reqBody2,
			"result": "successfully ordered",
		}
		c.JSON(http.StatusOK, res)
	}

	sqlStatement3 := `SELECT approve_grant,order_id FROM students_order_detail where id = $1`

	row3 := DB.QueryRow(sqlStatement3, reqBody2.Id)

	err3 := row3.Scan(&reqBody2.Approve_grant, &reqBody2.Order_ID)

	if err3 != nil {
		res := gin.H{
			"error":   err,
			"message": "sqlstatement3",
			"result":  sqlStatement3,
		}
		c.JSON(http.StatusBadRequest, res)
	} else {
		res := gin.H{
			"result": reqBody2,
		}
		c.JSON(http.StatusOK, res)
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

	fmt.Println("days", days)

	return int(days)

}

// ======================================completd========================================
