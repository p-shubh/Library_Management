package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func studentsOrderReq(c *gin.Context) {

	// if 1 == 1 {
	// 	res := gin.H{
	// 		"result": "true",
	// 	}
	// 	c.JSON(http.StatusOK, res)
	// }

	reqBody := OrderRequested{}

	// reqBody.Approve_grant = "pending"

	sqlStatatement := `SELECT "id", "book_id", "issue_date", "approve_grant", "order_id" FROM "public"."students_order_detail"
	where "approve_grant" = 'pending';`

	row := DB.QueryRow(sqlStatatement)

	// fmt.Println("check getUserByEmail", reqBody)

	err := row.Scan(&reqBody.Id, &reqBody.Book_id, &reqBody.Issue_date, &reqBody.Approve_grant, &reqBody.Order_ID)

	// fmt.Println("check getUserByEmail", reqBody)

	if err != nil {
		//log.Fatal("ppt", err)
		res := gin.H{
			"error":  "unable to view order list",
			"status": "there is no any pending requests",
		}
		c.JSON(http.StatusBadRequest, res)
		// c.Abort()
		// return

	} else {

		res := gin.H{
			"status":           "success",
			"pending_order_id": reqBody.Order_ID,
		}
		c.JSON(http.StatusOK, res)

	}

	// return

	sqlStatatement2 := `SELECT "order_id" FROM "public"."students_order_detail" where "approve_grant" = 'approved';`

	row2 := DB.QueryRow(sqlStatatement2)

	// fmt.Println("check getUserByEmail", reqBody)

	err2 := row2.Scan(&reqBody.Order_ID)

	// fmt.Println("check getUserByEmail", reqBody)

	if err2 != nil {
		// log.Fatal(err2)
		res := gin.H{
			"still pending order id": reqBody.Order_ID,
			"error":                  "no any approved order list",
		}
		c.JSON(http.StatusOK, res)
	} else {
		res := gin.H{
			"status":                 "success",
			"approved order_id list": reqBody.Order_ID,
		}
		c.JSON(http.StatusOK, res)
	}

	// return

}
