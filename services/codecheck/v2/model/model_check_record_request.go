package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CheckRecordRequest struct {
	// 项目ID

	ProjectId string `json:"project_id"`
	// 任务ID

	TaskId string `json:"task_id"`
	// 分页索引，偏移量

	Offset *int32 `json:"offset,omitempty"`
	// 每页显示的数量,每页最多显示1000条

	Limit *int32 `json:"limit,omitempty"`
	// 过滤开始时间,根据任务检查开始时间过滤

	StartTime *string `json:"start_time,omitempty"`
	// 过滤结束时间,根据任务检查开始时间过滤

	EndTime *string `json:"end_time,omitempty"`
}

func (o CheckRecordRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckRecordRequest struct{}"
	}

	return strings.Join([]string{"CheckRecordRequest", string(data)}, " ")
}
