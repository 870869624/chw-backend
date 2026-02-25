
package example

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/example"
    exampleReq "github.com/flipped-aurora/gin-vue-admin/server/model/example/request"
)

type ChwActivityService struct {}
// CreateChwActivity 创建活动管理记录
// Author [yourname](https://github.com/yourname)
func (chwactService *ChwActivityService) CreateChwActivity(ctx context.Context, chwact *example.ChwActivity) (err error) {
	err = global.GVA_DB.Create(chwact).Error
	return err
}

// DeleteChwActivity 删除活动管理记录
// Author [yourname](https://github.com/yourname)
func (chwactService *ChwActivityService)DeleteChwActivity(ctx context.Context, id string) (err error) {
	err = global.GVA_DB.Delete(&example.ChwActivity{},"id = ?",id).Error
	return err
}

// DeleteChwActivityByIds 批量删除活动管理记录
// Author [yourname](https://github.com/yourname)
func (chwactService *ChwActivityService)DeleteChwActivityByIds(ctx context.Context, ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]example.ChwActivity{},"id in ?",ids).Error
	return err
}

// UpdateChwActivity 更新活动管理记录
// Author [yourname](https://github.com/yourname)
func (chwactService *ChwActivityService)UpdateChwActivity(ctx context.Context, chwact example.ChwActivity) (err error) {
	err = global.GVA_DB.Model(&example.ChwActivity{}).Where("id = ?",chwact.Id).Updates(&chwact).Error
	return err
}

// GetChwActivity 根据id获取活动管理记录
// Author [yourname](https://github.com/yourname)
func (chwactService *ChwActivityService)GetChwActivity(ctx context.Context, id string) (chwact example.ChwActivity, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&chwact).Error
	return
}
// GetChwActivityInfoList 分页获取活动管理记录
// Author [yourname](https://github.com/yourname)
func (chwactService *ChwActivityService)GetChwActivityInfoList(ctx context.Context, info exampleReq.ChwActivitySearch) (list []example.ChwActivity, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&example.ChwActivity{})
    var chwacts []example.ChwActivity
    // 如果有条件搜索 下方会自动创建搜索语句
    
    if info.Id != nil {
        db = db.Where("id = ?", *info.Id)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
        var OrderStr string
        orderMap := make(map[string]bool)
         	orderMap["start_time"] = true
         	orderMap["end_time"] = true
         	orderMap["status"] = true
         	orderMap["id"] = true
       if orderMap[info.Sort] {
          OrderStr = info.Sort
          if info.Order == "descending" {
             OrderStr = OrderStr + " desc"
          }
          db = db.Order(OrderStr)
       }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&chwacts).Error
	return  chwacts, total, err
}
func (chwactService *ChwActivityService)GetChwActivityPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
