package example

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/example"
    exampleReq "github.com/flipped-aurora/gin-vue-admin/server/model/example/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type ChwActivityApi struct {}



// CreateChwActivity 创建活动管理
// @Tags ChwActivity
// @Summary 创建活动管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body example.ChwActivity true "创建活动管理"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /chwact/createChwActivity [post]
func (chwactApi *ChwActivityApi) CreateChwActivity(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var chwact example.ChwActivity
	err := c.ShouldBindJSON(&chwact)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = chwactService.CreateChwActivity(ctx,&chwact)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteChwActivity 删除活动管理
// @Tags ChwActivity
// @Summary 删除活动管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body example.ChwActivity true "删除活动管理"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /chwact/deleteChwActivity [delete]
func (chwactApi *ChwActivityApi) DeleteChwActivity(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	id := c.Query("id")
	err := chwactService.DeleteChwActivity(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteChwActivityByIds 批量删除活动管理
// @Tags ChwActivity
// @Summary 批量删除活动管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /chwact/deleteChwActivityByIds [delete]
func (chwactApi *ChwActivityApi) DeleteChwActivityByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ids := c.QueryArray("ids[]")
	err := chwactService.DeleteChwActivityByIds(ctx,ids)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateChwActivity 更新活动管理
// @Tags ChwActivity
// @Summary 更新活动管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body example.ChwActivity true "更新活动管理"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /chwact/updateChwActivity [put]
func (chwactApi *ChwActivityApi) UpdateChwActivity(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var chwact example.ChwActivity
	err := c.ShouldBindJSON(&chwact)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = chwactService.UpdateChwActivity(ctx,chwact)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindChwActivity 用id查询活动管理
// @Tags ChwActivity
// @Summary 用id查询活动管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id query int true "用id查询活动管理"
// @Success 200 {object} response.Response{data=example.ChwActivity,msg=string} "查询成功"
// @Router /chwact/findChwActivity [get]
func (chwactApi *ChwActivityApi) FindChwActivity(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	id := c.Query("id")
	rechwact, err := chwactService.GetChwActivity(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(rechwact, c)
}
// GetChwActivityList 分页获取活动管理列表
// @Tags ChwActivity
// @Summary 分页获取活动管理列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query exampleReq.ChwActivitySearch true "分页获取活动管理列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /chwact/getChwActivityList [get]
func (chwactApi *ChwActivityApi) GetChwActivityList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo exampleReq.ChwActivitySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := chwactService.GetChwActivityInfoList(ctx,pageInfo)
	if err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败:" + err.Error(), c)
        return
    }
    response.OkWithDetailed(response.PageResult{
        List:     list,
        Total:    total,
        Page:     pageInfo.Page,
        PageSize: pageInfo.PageSize,
    }, "获取成功", c)
}

// GetChwActivityPublic 不需要鉴权的活动管理接口
// @Tags ChwActivity
// @Summary 不需要鉴权的活动管理接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /chwact/getChwActivityPublic [get]
func (chwactApi *ChwActivityApi) GetChwActivityPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    chwactService.GetChwActivityPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的活动管理接口信息",
    }, "获取成功", c)
}
