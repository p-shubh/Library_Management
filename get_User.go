package main

import "fmt"

func getUserByEmail(email string) USER {

	reqBody := USER{}

	sqlStatement := "SELECT id, firstname, lastname, email,user_type FROM signup_detail where email = $1"

	row := DB.QueryRow(sqlStatement, email)

	fmt.Println("cvhgvh", reqBody)

	row.Scan(&reqBody.Id, &reqBody.First_name, &reqBody.Last_name, &reqBody.Email, &reqBody.User_type)

	fmt.Println("cvhgvh", reqBody)

	return reqBody
}

func getUserByGmail(id int) USER {

	reqBody := USER{}

	sqlStatement := "SELECT id, firstname, lastname, email,user_type FROM signup_detail where id = $1"

	row := DB.QueryRow(sqlStatement, id)

	row.Scan(&reqBody.Id, &reqBody.First_name, &reqBody.Last_name, &reqBody.Email, &reqBody.User_type)

	return reqBody
}
