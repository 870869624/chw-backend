package example

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ChwActivityRouter struct {}

// InitChwActivityRouter 初始化 活动管理 路由信息
func (s *ChwActivityRouter) InitChwActivityRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	chwactRouter := Router.Group("chwact").Use(middleware.OperationRecord())
	chwactRouterWithoutRecord := Router.Group("chwact")
	chwactRouterWithoutAuth := PublicRouter.Group("chwact")
	{
		chwactRouter.POST("createChwActivity", chwactApi.CreateChwActivity)   // 新建活动管理
		chwactRouter.DELETE("deleteChwActivity", chwactApi.DeleteChwActivity) // 删除活动管理
		chwactRouter.DELETE("deleteChwActivityByIds", chwactApi.DeleteChwActivityByIds) // 批量删除活动管理
		chwactRouter.PUT("updateChwActivity", chwactApi.UpdateChwActivity)    // 更新活动管理
	}
	{
		chwactRouterWithoutRecord.GET("findChwActivity", chwactApi.FindChwActivity)        // 根据ID获取活动管理
		chwactRouterWithoutRecord.GET("getChwActivityList", chwactApi.GetChwActivityList)  // 获取活动管理列表
	}
	{
	    chwactRouterWithoutAuth.GET("getChwActivityPublic", chwactApi.GetChwActivityPublic)  // 活动管理开放接口
	}
}
