package router

import (
	"github.com/gin-gonic/gin"
	"github.com/onyas/geekNews/handler"
	"net/http"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "jobs.html", nil)
	})

	// 添加 Get 请求路由
	router.GET("/cronJobs", handler.CronJobs)
	router.GET("/listJobInfos", handler.ListJobInfos)

	return router
}
