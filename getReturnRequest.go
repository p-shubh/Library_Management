package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getReturnRequest(c *gin.Context) {

	// reqBody := OrderApprove{}

	// sqlStatatement := `select order_id from students_order_detail where return_request = 'pending'`

	// row := DB.QueryRow(sqlStatatement)

	// err := row.Scan(&reqBody.Order_ID)

	// if err != nil {
	// 	//log.Fatal("ppt", err)
	// 	res := gin.H{
	// 		"error":  "unable to view order",
	// 		"status": "check in get history from id",
	// 	}
	// 	c.JSON(http.StatusBadRequest, res)
	// 	// c.Abort()
	// 	// return

	// } else {

	// 	res := gin.H{
	// 		"status":           "success",
	// 		"student order id": reqBody.Order_ID,
	// 	}
	// 	c.JSON(http.StatusOK, res)

	// }

	// var return_request string

	return_request := "pending"

	total_rows := count_rows2(return_request)

	if total_rows > 0 {

		users := []OrderApprove{}

		// var approved_orders_id int

		sqlStatatement1 := (`select order_id from students_order_detail where return_request = 'pending'`)

		rows, err := DB.Query(sqlStatatement1)

		if err != nil {
			log.Println("Failed to execute query in get return request : ", err)
			return
		}

		defer rows.Close()
		user := OrderApprove{}
		for rows.Next() {
			rows.Scan(&user.Order_ID)
			users = append(users, user)
		}

		res := gin.H{
			"pending return request": users,
		}

		c.JSON(http.StatusOK, res)

		// return

		// users := []User{}
		// userSQL := "SELECT id, name, email,user_id, phone, city,  password FROM users"

	}

}

func count_rows2(grant_status string) int {
	var count int

	err := DB.QueryRow("SELECT COUNT(*) FROM students_order_detail where return_request = $1", grant_status).Scan(&count)
	switch {
	case err != nil:
		log.Fatal(err)
	default:
		fmt.Println("Number of rows are", count)
	}
	return count
}
