package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 任务关联的子任务信息
type SubTaskAssociatedWithTask struct {

	// 子任务id
	Id *int64 `json:"id,omitempty" xml:"id"`

	// 子任务名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 子任务的进度，取值为0-100之间的整数
	Progress *int32 `json:"progress,omitempty" xml:"progress"`

	// 子任务开始时间
	StartDate *int64 `json:"start_date,omitempty" xml:"start_date"`

	// 子任务结束时间（如果子任务还没有结束，则为空）
	EndDate *int64 `json:"end_date,omitempty" xml:"end_date"`

	// 迁移或同步时，具体的迁移详情
	ProcessTrace *string `json:"process_trace,omitempty" xml:"process_trace"`
}

func (o SubTaskAssociatedWithTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubTaskAssociatedWithTask struct{}"
	}

	return strings.Join([]string{"SubTaskAssociatedWithTask", string(data)}, " ")
}
