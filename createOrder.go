package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func createOrder(c *gin.Context) {

	type OrderReqBookID struct {
		Id      int    `json:"id"`
		Book_id string `json:"book_id"`
	}

	reqBody := OrderReqBookID{}
	fmt.Println("create order request body", reqBody)
	reqBody.Id = live_user["presentuser"].Id
	fmt.Println("create order id", reqBody.Id)

	sqlStatement := "SELECT book_id FROM books_detail where book_id= $1"

	row := DB.QueryRow(sqlStatement, reqBody.Book_id)

	row.Scan(&reqBody.Book_id)

	err := c.Bind(&reqBody) //binding with the data being provided

	if err != nil {
		res := gin.H{
			"unable to get book id": "err",
		}

		c.JSON(http.StatusBadRequest, res)
	}

	fmt.Println("book_id", reqBody.Book_id)

	reqBody2 := ORDER{}

	reqBody2.Id = live_user["presentuser"].Id

	reqBody2.Book_id = reqBody.Book_id

	sqlStatement2 := "SELECT book_title,book_author,book_cover_image FROM books_detail where book_id= $1"

	row2 := DB.QueryRow(sqlStatement2, reqBody.Book_id)

	// row2.Scan(&reqBody2.Book_title, reqBody2.Book_author, reqBody2.Book_cover_image,)

	row2.Scan(&reqBody2.Book_title, &reqBody2.Book_author, &reqBody2.Book_cover_image)

	err2 := c.Bind(&reqBody) //binding with the data being provided

	if err2 != nil {
		res := gin.H{
			"unable to insert the data": "err",
		}

		c.JSON(http.StatusBadRequest, res)
	}

	// if reqBody.Book_id != int {

	// }

}
