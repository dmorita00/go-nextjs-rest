package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/model"
	_ "server/model"
	"server/service"
	"strconv"
	_ "strconv"
)

func Test(c *gin.Context) {
	testService := service.TestService{}
	TestLists := testService.GetTestList()
	c.JSONP(http.StatusOK, gin.H{
		"message": "ok",
		"data":    TestLists,
	})
}

func TestAdd(c *gin.Context) {
	test := model.Test{}
	err := c.Bind(&test)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	testService := service.TestService{}
	err = testService.SetTest(&test)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}

func TestList(c *gin.Context) {
	testService := service.TestService{}
	TestLists := testService.GetTestList()
	c.JSONP(http.StatusOK, gin.H{
		"message": "ok",
		"data":    TestLists,
	})
}

func TestUpdate(c *gin.Context) {
	test := model.Test{}
	err := c.Bind(&test)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	testService := service.TestService{}
	err = testService.UpdateTest(&test)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}

func TestDelete(c *gin.Context) {
	id := c.PostForm("id")
	intId, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}
	testService := service.TestService{}
	err = testService.DeleteTest(int(intId))
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}
