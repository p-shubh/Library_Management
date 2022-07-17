package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Return_approve_with_fine(c *gin.Context) {

	reqBody := OrderApprove{}
	c.Bind(&reqBody)

	if reqBody.Order_ID == 0 {
		res := gin.H{
			"status": "order_id ",
			"result": reqBody,
		}
		c.JSON(http.StatusBadRequest, res)
		// return
		// c.Abort()
	}
	// ================================================== order_id check =======================================

	// count, err := DB.Query("SELECT COUNT(*) FROM main_table where order_id=$1",reqBody.Order_ID)

	var count int

	err := DB.QueryRow("SELECT COUNT(*) FROM students_order_detail where order_id = $1", reqBody.Order_ID).Scan(&count)
	switch {
	case err != nil:
		log.Fatal(err)
	default:
		fmt.Println("Number of rows are", count)
	}

	// ================================================== order_id check =======================================

	if count == 0 {
		res := gin.H{
			"status": "order_id is incorrect",
		}
		c.JSON(http.StatusBadRequest, res)
		c.Abort()

	} else {
		res := gin.H{
			"status": "order id is correct",
		}
		c.JSON(http.StatusOK, res)

	}

	// ID_cookie := `select id from `

	// id, _ := strconv.Atoi(ID_cookie)

	Data := get_students_order_detail_by_orders_id(reqBody.Order_ID)

	// res := gin.H{
	// 	"result": Data,
	// }

	// c.JSON(http.StatusOK, res)

	// Total_days := calculateTime(Data.Issue_date, reqBody.Student_Return_date)

	// ================================================= fine calculation ===================================

	total_time := calculateTime(Data.Issue_date, Data.Student_Return_date)

	fmt.Print("Total_days = ", total_time)

	// if total_time > 30 {

	// fine_time := total_time - 30
	fine_amount := 10

	total_fine := total_time * fine_amount

	res := gin.H{
		// 		"message": "update student return date gone wrong",
		// 		"result":  updateOrderSQL,
		// 	}
		// 	c.JSON(http.StatusBadRequest, res)
		// 	c.Abort()
		// 	return
		// } else {
		// res := gin.H{
		// "status":     "update Student_Return_date, total_fine",
		"total_fine": total_fine,
		"result":     Data,
	}
	c.JSON(http.StatusOK, res)

	var return_grant string
	var return_request string

	if total_time > 30 {
		return_grant = "pending"
		return_request = "approve"
	} else {
		return_grant = "approve"
		return_request = "approve"

	}

	// ============================================= update pending ===========================================

	updateOrderSQL := `UPDATE students_order_detail SET total_fine= $1,return_grant=$2, return_request=$3 WHERE order_id = $4;` //work is pemding******

	_, err = DB.Exec(updateOrderSQL, total_fine, return_grant, return_request, reqBody.Order_ID)

	if err != nil {

		res := gin.H{
			"message": "update student return date gone wrong",
			"result":  updateOrderSQL,
		}
		c.JSON(http.StatusBadRequest, res)
		// c.Abort()
		return
	} else {

		res := gin.H{
			"message": "update succesfully",
			"result":  updateOrderSQL,
		}
		c.JSON(http.StatusOK, res)
		c.Abort()
		return

	}

	// res := gin.H{
	// 	"total_fine": total_fine,
	// }
	// c.JSON(http.StatusOK, res)

	// ============================================= inseting fine ===========================================

	// updateOrderSQL := `UPDATE students_order_detail SET student_return_date= $1, total_fine= $2 WHERE order_id = $3;`

	// _, err = DB.Exec(updateOrderSQL, reqBody.Student_Return_date, total_fine, reqBody.Order_ID)

	// if err != nil {

	// 	res := gin.H{
	// 		"message": "update student return date gone wrong",
	// 		"result":  updateOrderSQL,
	// 	}
	// 	c.JSON(http.StatusBadRequest, res)
	// 	c.Abort()
	// 	return
	// } else {
	// 	res := gin.H{
	// 		"status":     "update Student_Return_date, total_fine",
	// 		"total_fine": total_fine,
	// 		"result":     Data,
	// 	}
	// 	c.JSON(http.StatusOK, res)
	// }

	// } else {

	// 	total_fine := 0

	// 	updateOrderSQL := `UPDATE students_order_detail SET student_return_date= $1, total_fine= $2 WHERE order_id = $3;`

	// 	_, err := DB.Exec(updateOrderSQL, reqBody.Student_Return_date, total_fine, reqBody.Order_ID)

	// 	if err != nil {

	// 		res := gin.H{
	// 			"message": "update student return date gone wrong",
	// 			"result":  updateOrderSQL,
	// 		}
	// 		c.JSON(http.StatusBadRequest, res)
	// 		c.Abort()
	// 		return
	// 	} else {
	// 		res := gin.H{
	// 			"status": "update Student_Return_date, total_fine",
	// 			"result": "success",
	// 			"message": "no fine get charged",

	// 		}
	// 		c.JSON(http.StatusOK, res)
	// 	}

	// }

}

// func total_fine()
