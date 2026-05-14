package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EditingTaskInfo struct {

	// 编辑任务ID
	TaskId *string `json:"task_id,omitempty"`

	// 输出的媒资ID
	AssetId *string `json:"asset_id,omitempty"`

	// 任务状态 - WAITING 等待中 - PROCESSING 处理中； - SUCCEED 成功； - FAILED 失败； - CANCEL 已取消；
	Status *string `json:"status,omitempty"`

	// 进度，取值0-100
	Progress *int32 `json:"progress,omitempty"`

	// 任务创建时间，格式按照RFC3339，UTC时间，如2020-09-01T18:50:20Z
	CreateTime *string `json:"create_time,omitempty"`

	// 任务结束时间，格式按照RFC3339，UTC时间，如2020-09-01T18:50:20Z，当任务结束时才有值
	EndTime *string `json:"end_time,omitempty"`
}

func (o EditingTaskInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EditingTaskInfo struct{}"
	}

	return strings.Join([]string{"EditingTaskInfo", string(data)}, " ")
}
