package routers

import (
	"github.com/Kozzen890/assignment3-016/helpers"
	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine{
	res := gin.Default()
	res.PUT("/apiUpdate", helpers.UpdateDataAPI)
	return res
}