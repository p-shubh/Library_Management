package main

import (
	"log"
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

	sqlStatatement := `SELECT "id", "book_id", "issue_date", "return_date", "approve_grant" FROM "public"."students_order_detail"
	where "approve_grant" = 'pending';`

	row := DB.QueryRow(sqlStatatement)

	// fmt.Println("check getUserByEmail", reqBody)

	err := row.Scan(&reqBody.Id, &reqBody.Book_id, &reqBody.Issue_date, &reqBody.Return_date, &reqBody.Approve_grant)

	// fmt.Println("check getUserByEmail", reqBody)

	if err != nil {
		log.Fatal(err)
		res := gin.H{
			"error": "unable to view order list",
		}
		c.JSON(http.StatusBadRequest, res)
	} else {
		res := gin.H{
			"status": "success",
			"result": reqBody,
		}
		c.JSON(http.StatusOK, res)
	}

	return

}
