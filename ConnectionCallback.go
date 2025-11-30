package http_con

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	RouteMethod_GET  = "GET"
	RouteMethod_POST = "POST"
)

type HttpPath struct {
	Name     string
	Callback func(*gin.Context)
	Method   string
}

type HttpGroupPath struct {
	Name  string
	Paths []HttpPath
}

// `...` is struct tag that some lib can read it
//
//	gin will get from url and map name to var
//	if not use struct tag, we cannot change var that use in app if we change params
type QueryUserParms struct {
	Name string `form:"name"`
	Id   string `form:"id"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func getUser(c *gin.Context) {
	var queryUser QueryUserParms
	err := c.ShouldBind(&queryUser)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":     "Mismatch params",
			"traceback": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": queryUser.Name,
	})

	fmt.Println(queryUser)
}

func login(c *gin.Context) {
	var loginReq LoginRequest
	err := c.ShouldBindJSON(&loginReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed",
		})
	}

	// authentication

	c.JSON(http.StatusOK, gin.H{})
}
