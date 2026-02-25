import service from '@/utils/request'
// @Tags ChwActivity
// @Summary 创建活动管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ChwActivity true "创建活动管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /chwact/createChwActivity [post]
export const createChwActivity = (data) => {
  return service({
    url: '/chwact/createChwActivity',
    method: 'post',
    data
  })
}

// @Tags ChwActivity
// @Summary 删除活动管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ChwActivity true "删除活动管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /chwact/deleteChwActivity [delete]
export const deleteChwActivity = (params) => {
  return service({
    url: '/chwact/deleteChwActivity',
    method: 'delete',
    params
  })
}

// @Tags ChwActivity
// @Summary 批量删除活动管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除活动管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /chwact/deleteChwActivity [delete]
export const deleteChwActivityByIds = (params) => {
  return service({
    url: '/chwact/deleteChwActivityByIds',
    method: 'delete',
    params
  })
}

// @Tags ChwActivity
// @Summary 更新活动管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ChwActivity true "更新活动管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /chwact/updateChwActivity [put]
export const updateChwActivity = (data) => {
  return service({
    url: '/chwact/updateChwActivity',
    method: 'put',
    data
  })
}

// @Tags ChwActivity
// @Summary 用id查询活动管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.ChwActivity true "用id查询活动管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /chwact/findChwActivity [get]
export const findChwActivity = (params) => {
  return service({
    url: '/chwact/findChwActivity',
    method: 'get',
    params
  })
}

// @Tags ChwActivity
// @Summary 分页获取活动管理列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取活动管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /chwact/getChwActivityList [get]
export const getChwActivityList = (params) => {
  return service({
    url: '/chwact/getChwActivityList',
    method: 'get',
    params
  })
}

// @Tags ChwActivity
// @Summary 不需要鉴权的活动管理接口
// @Accept application/json
// @Produce application/json
// @Param data query exampleReq.ChwActivitySearch true "分页获取活动管理列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /chwact/getChwActivityPublic [get]
export const getChwActivityPublic = () => {
  return service({
    url: '/chwact/getChwActivityPublic',
    method: 'get',
  })
}
