package main

import (
	// "encoding/csv"
	"database/sql"

	"github.com/gin-gonic/gin"
)

var DB *sql.DB

// type date interface {
// }

type Login struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required,alphanum,min=12" `
}

type USER struct {
	Id         int    `json:"id"`
	First_name string `json:"first_name" binding:"required"`
	Last_name  string `json:"last_name"`
	Email      string `json:"email" binding:"required"`
	Password   string `json:"password" binding:"required,alphanum,min=12" `
	User_type  int    `json:"user_type"`
}

type ORDER struct {
	Id               int    `json:"id"`
	Book_id          string `json:"book_id"`
	Book_title       string `json:"book_title"`
	Book_author      string `json:"book_author"`
	Book_cover_image string `json:"book_cover_image"`
	Issue_date       string `json:"issue_date"`
	Return_date      string `json:"return_date"`
	Fine             string `json:"fine"`
	Approve_grant    string `json:"approve_grant"`
}

type OrderRequested struct {
	Id                  int    `json:"id"`
	Book_id             string `json:"book_id"`
	Issue_date          string `json:"issue_date"`
	Return_date         string `json:"return_date"`
	Approve_grant       string `json:"approve_grant"`
	Order_ID            int    `json:"order_id"`
	Student_Return_date string `json:"student_return_date"`
}

type OrderApprove struct {
	Order_ID int `json:"order_id"`
}

type Return_Fine struct {
	Student_Return_date string `json:"student_return_date"`
	// Fine        string `json:"fine"`
	Order_ID int `json:"order_id"`
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

	// c := isLogin()

	// fmt.Println(c)
}

func setupRoutes(g *gin.Engine) {

	g.POST("/signup", SignUpPostHandler)                             //for students
	g.POST("/signup/admin", SignUpPostHandler)                       //for admin
	g.POST("/login", LoginPostHandler)                               //for students and admin
	g.POST("/logout", logout)                                        //for students and admin
	g.POST("/requestorder", isStudentLogin(), OrderRequest)          //for students
	g.GET("/studentsrequestorder", isAdminLogin(), studentsOrderReq) //for admin
	g.POST("/approveorders", isAdminLogin(), approveOrders)          //for admin
	g.POST("/return", isStudentLogin(), Return_with_fine)            //for students

}
