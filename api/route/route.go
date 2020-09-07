package route

import (
	"collect_int_module_sn/api/controller"

	"github.com/gin-gonic/gin"
)

func InitRoute() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	moduleV1 := v1.Group("/module")
	moduleV1.POST("/sendtask", controller.SendTask)
	moduleV1.GET("/querytasklist", controller.QueryTaskList)
	moduleV1.POST("/getresult", controller.GetResult)

	return r
}
