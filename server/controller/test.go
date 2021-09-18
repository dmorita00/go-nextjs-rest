package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	_ "server/model"
	"server/service"
	_ "strconv"
)

func Test(c *gin.Context) {
	bookService := service.TestService{}
	BookLists := bookService.GetTestList()
	c.JSONP(http.StatusOK, gin.H{
		"message": "ok",
		"data":    BookLists,
	})
}
