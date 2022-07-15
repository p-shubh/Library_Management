package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func approveOrders(c *gin.Context) {

	reqBody := OrderApprove{}
	c.Bind(&reqBody)

	// count_rows()================================================================
	rows, err2 := DB.Query("SELECT COUNT(*) FROM students_order_detail where order_id=$1", reqBody.Order_ID)

	if err2 != nil {
		log.Fatal("count rows", err2)
	}
	defer rows.Close()

	var count int

	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			log.Fatal(err)
		}
	}

	if count == 0 {
		res := gin.H{
			"message":     "no any order in this order id",
			"order count": count,
		}
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	}

	fmt.Println("Number of rows are ", count)

	// sqlStatement :=================================================================
	approveOrderSQL := `update students_order_detail set approve_grant = 'approved' where order_id = $1`

	_, err := DB.Exec(approveOrderSQL, reqBody.Order_ID)

	if err != nil {

		res := gin.H{
			"message": "insertOrderSQL is not inserting",
			"result":  approveOrderSQL,
		}
		c.JSON(http.StatusBadRequest, res)
		c.Abort()
		return
	} else {
		res := gin.H{
			"status": reqBody.Order_ID,
			"result": "approved success",
		}
		c.JSON(http.StatusOK, res)
	}

}

// func check_order_id(order_id int){

// 	approveOrderSQL := `update students_order_detail set approve_grant = 'approved' where order_id = $1`

// }
// func count_rows() {

// rows, err := DB.Query("SELECT COUNT(*) FROM students_order_detail")
// if err != nil {
// 	log.Fatal(err)
// }
// defer rows.Close()

// var count int

// for rows.Next() {
// 	if err := rows.Scan(&count); err != nil {
// 		log.Fatal(err)
// 	}
// }

// fmt.Println("Number of rows are ", count)
// }
