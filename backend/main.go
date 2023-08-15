package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Email    string
	Password string
}

// 
func login(c *gin.Context) {
	var user User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code": http.StatusBadRequest,
			"data": err.Error(),
		})
		return
	}

	if user.Email == "admin@gmail.com" && user.Password == "12345" {
		c.JSON(http.StatusOK, map[string]interface{}{
			"code": http.StatusOK,
			"data": "sukses login",
		})
	} else {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code": http.StatusBadRequest,
			"data": "email atau password salah",
		})
	}
}

func main() {
	r := gin.Default()

	r.POST("/login", login)

	if err := r.Run(":9000"); err != nil {
		panic(err)
	}
}
