package main

import (
	// "encoding/csv"
	"database/sql"

	"github.com/gin-gonic/gin"
)

var DB *sql.DB

type USER struct {
	Id         int    `json:"id"`
	First_name string `json:"first_name" binding:"required"`
	Last_name  string `json:"last_name"`
	Email      string `json:"email" binding:"required"`
	Password   string `json:"password" binding:"required,alphanum,min=12" `
	User_type  string `json:"user_type"`
}

type ORDER struct {
	Book_id          string `json:"book_id"`
	Id               int    `json:"id"`
	Book_title       string `json:"book_title"`
	Book_author      string `json:"book_author"`
	Book_cover_image string `json:"book_cover_image"`
	Issue_date       string `json:"issue_date"`
	Return_date      string `json:"return_date"`
	Fine             string `json:"fine"`
}

var (
	Data map[string]USER
)

func main() {

	// fmt.Println(books_csv)

	connection_with_db()
	defer DB.Close()
	// books_csv := readCsvFile("./books.csv")
	// importcsv(books_csv)

	router := gin.Default()
	setupRoutes(router)
	router.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func setupRoutes(g *gin.Engine) {

	g.POST("/signup", SignUpPostHandler)
	g.POST("/signup/admin", SignUpPostHandler)

	// g.POST("/signupstudent", StudentSignUpPostHandler)
	g.POST("/login", LoginPostHandler)
	g.POST("/logout", logout)

	// g.POST("/loginadmin", AdminLoginPOSTHandler)
	// g.PUT("/updtaeadmin", AdminUpdtaePUTHandler)
	// g.PUT("/updtaeadmin", AdminUpdtaePUTHandler)
	// g.DELETE("/deleteadmin", AdminDELETEHandler)
	// g.DELETE("/deletestudent", StudentDELETEHandler)

}
