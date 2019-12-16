package router

import (
	"github.com/gin-gonic/gin"
	"github.com/onyas/geekNews/handler"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// 添加 Get 请求路由
	router.GET("/cronJobs", handler.CronJobs)
	router.GET("/listJobInfos", handler.ListJobInfos)

	return router
}
