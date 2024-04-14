package handler

import (
	"daydayup/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Greetings used to test query parameters not in request body
func Greetings(c *gin.Context) {
	name, ok := c.GetQuery("name")
	if !ok {
		c.JSON(http.StatusBadRequest, global.ErrorWithMsg(http.StatusBadRequest, "bad greetings!"))
	}
	c.JSON(http.StatusOK, global.SuccessWithData("Hello "+name+"! Welcome to Golang"))
	return
}
