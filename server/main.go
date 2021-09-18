package main

//import "github.com/gin-gonic/gin"
import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"server/controller"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// ミドルウェア
	//r.Use(middleware.RecordUaAndTime)
	// CRUD 書籍
	testEngine := r.Group("/test")
	{
		v1 := testEngine.Group("/v1")
		{
			v1.GET("/list", controller.Test)
		}
	}
	r.Run() // listen and serve on 0.0.0.0:8080
}
