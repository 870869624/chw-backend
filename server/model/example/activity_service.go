
// 自动生成模板ChwActivity
package example
import (
	"time"
	"gorm.io/datatypes"
)

// 活动管理 结构体  ChwActivity
type ChwActivity struct {
  Name  *string `json:"name" form:"name" gorm:"comment:活动名称;column:name;" binding:"required"`  //活动名称
  Type  string `json:"type" form:"type" gorm:"comment:活动类型;column:type;type:enum('线上活动','线下活动','混合活动');" binding:"required"`  //活动类型
  Content  *string `json:"content" form:"content" gorm:"comment:活动内容;column:content;type:text;" binding:"required"`  //活动内容
  StartTime  *time.Time `json:"startTime" form:"startTime" gorm:"comment:活动开始时间;column:start_time;" binding:"required"`  //开始时间
  EndTime  *time.Time `json:"endTime" form:"endTime" gorm:"comment:活动结束时间;column:end_time;" binding:"required"`  //结束时间
  Status  string `json:"status" form:"status" gorm:"comment:活动状态;column:status;type:enum('未开始','进行中','已结束','已取消');" binding:"required"`  //活动状态
  Participants  *int64 `json:"participants" form:"participants" gorm:"comment:参与人数;column:participants;"`  //参与人数
  Images  datatypes.JSON `json:"images" form:"images" gorm:"comment:活动图片;column:images;" swaggertype:"array,object"`  //活动图片
  Contact  *string `json:"contact" form:"contact" gorm:"comment:联系方式;column:contact;"`  //联系方式
  Location  *string `json:"location" form:"location" gorm:"comment:活动地点;column:location;"`  //活动地点
  Budget  *float64 `json:"budget" form:"budget" gorm:"comment:活动预算;column:budget;"`  //活动预算
  Organizer  *string `json:"organizer" form:"organizer" gorm:"comment:活动组织者;column:organizer;" binding:"required"`  //组织者
  Id  *int64 `json:"id" form:"id" gorm:"primarykey;column:id;" binding:"required"`  //活动ID
}


// TableName 活动管理 ChwActivity自定义表名 chw_activity
func (ChwActivity) TableName() string {
    return "chw_activity"
}





