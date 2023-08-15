package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubTaskAssociatedWithTask 任务关联的子任务信息
type SubTaskAssociatedWithTask struct {

	// 子任务ID
	Id *int64 `json:"id,omitempty"`

	// 子任务名称
	Name *string `json:"name,omitempty"`

	// 子任务的进度，取值为0-100之间的整数
	Progress *int32 `json:"progress,omitempty"`

	// 子任务开始时间
	StartDate *int64 `json:"start_date,omitempty"`

	// 子任务结束时间（如果子任务还没有结束，则为空）
	EndDate *int64 `json:"end_date,omitempty"`

	// 迁移或同步时，具体的迁移详情
	ProcessTrace *string `json:"process_trace,omitempty"`
}

func (o SubTaskAssociatedWithTask) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubTaskAssociatedWithTask struct{}"
	}

	return strings.Join([]string{"SubTaskAssociatedWithTask", string(data)}, " ")
}
