package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SchedulerBean 调度信息
type SchedulerBean struct {

	// 数据库ID
	DbIds *string `json:"db_ids,omitempty"`

	// 文件类型,默认ZIP
	Format string `json:"format"`

	// 周期
	Frequency string `json:"frequency"`

	// 模板ID
	Id string `json:"id"`

	// 调度方式
	Mode string `json:"mode"`

	// 是否通知  - OFF：不通知  - ON：通知
	Notice string `json:"notice"`

	// 开始时间
	StartTime string `json:"start_time"`

	// 模板状态
	Status string `json:"status"`

	// 主题URN
	TopicUrn string `json:"topic_urn"`
}

func (o SchedulerBean) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SchedulerBean struct{}"
	}

	return strings.Join([]string{"SchedulerBean", string(data)}, " ")
}
