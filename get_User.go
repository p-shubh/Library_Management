package main

import "fmt"

func getUserByEmail(email string) USER {

	reqBody := USER{}

	sqlStatement := "SELECT id, firstname, lastname, email,password,user_type FROM signup_detail where email = $1"

	row := DB.QueryRow(sqlStatement, email)

	fmt.Println("check getUserByEmail", reqBody)

	row.Scan(&reqBody.Id, &reqBody.First_name, &reqBody.Last_name, &reqBody.Email, &reqBody.Password, &reqBody.User_type)

	fmt.Println("check getUserByEmail", reqBody)

	return reqBody
}

func getUserByid(id int) USER {

	reqBody := USER{}

	sqlStatement := "SELECT id, firstname, lastname, email,password,user_type FROM signup_detail where id = $1"

	row := DB.QueryRow(sqlStatement, id)

	row.Scan(&reqBody.Id, &reqBody.First_name, &reqBody.Last_name, &reqBody.Email, &reqBody.Password, &reqBody.User_type)

	return reqBody
}

func get_students_order_detail_by_orders_id(id int) OrderRequested {

	reqBody := OrderRequested{}

	sqlStatement := "SELECT id,book_id,issue_date,return_date,student_return_date,approve_grant,order_id,return_grant,total_fine,return_request FROM students_order_detail where id = $1"

	row := DB.QueryRow(sqlStatement, id)

	row.Scan(&reqBody.Id, &reqBody.Book_id, &reqBody.Issue_date, &reqBody.Return_date, &reqBody.Student_Return_date, &reqBody.Approve_grant, &reqBody.Order_ID)

	return reqBody

}

func students_order_detail_by_order_id(order_id int) OrderRequested {

	reqBody := OrderRequested{}

	sqlStatement := "SELECT id,book_id,issue_date,return_date,student_return_date,approve_grant,order_id,return_grant,total_fine,return_request FROM students_order_detail where order_id = $1"

	row := DB.QueryRow(sqlStatement, order_id)

	row.Scan(&reqBody.Id, &reqBody.Book_id, &reqBody.Issue_date, &reqBody.Return_date, &reqBody.Student_Return_date, &reqBody.Approve_grant, &reqBody.Order_ID, &reqBody.Return_grant, &reqBody.Total_fine, &reqBody.Return_request)

	return reqBody

}
