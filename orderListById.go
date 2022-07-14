package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func try(c *gin.Context) {
	fmt.Println("try fucntion is successfully running after islogin function")
}
