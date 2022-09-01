package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 修改任务进度
type SubTask struct {

	// 子任务名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 子任务的进度，取值为0-100之间的整数
	Progress int32 `json:"progress" xml:"progress"`

	// 子任务开始时间
	StartDate *int64 `json:"start_date,omitempty" xml:"start_date"`

	// 子任务结束时间（如果子任务还没有结束，则为空）
	EndDate *int64 `json:"end_date,omitempty" xml:"end_date"`

	// 迁移速率，Mbit/s
	MigrateSpeed *float64 `json:"migrate_speed,omitempty" xml:"migrate_speed"`

	// 触发子任务的用户操作名称
	UserOp *string `json:"user_op,omitempty" xml:"user_op"`

	// 迁移或同步时，具体的迁移详情
	ProcessTrace *string `json:"process_trace,omitempty" xml:"process_trace"`
}

func (o SubTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubTask struct{}"
	}

	return strings.Join([]string{"SubTask", string(data)}, " ")
}
