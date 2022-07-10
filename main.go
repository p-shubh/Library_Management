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
	// g.POST("/loginadmin", AdminLoginPOSTHandler)
	// g.PUT("/updtaeadmin", AdminUpdtaePUTHandler)
	// g.PUT("/updtaeadmin", AdminUpdtaePUTHandler)
	// g.DELETE("/deleteadmin", AdminDELETEHandler)
	// g.DELETE("/deletestudent", StudentDELETEHandler)

}
