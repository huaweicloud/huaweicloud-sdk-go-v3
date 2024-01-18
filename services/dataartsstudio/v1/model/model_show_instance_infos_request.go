package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceInfosRequest Request Object
type ShowInstanceInfosRequest struct {

	// 实例id
	Instance string `json:"instance"`

	// 空间id
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// 作业算子id
	TaskId string `json:"task_id"`

	// 作业算子名称
	JobName string `json:"job_name"`

	// 搜索参数时间范围的开始时间，例：1705248000000
	StartTime float32 `json:"start_time"`

	// 搜索参数时间范围的结束时间,例：1705311669796
	EndTime float32 `json:"end_time"`

	// 分页参数偏移量
	Offset int32 `json:"offset"`

	// 分页参每页数量
	Limit int32 `json:"limit"`
}

func (o ShowInstanceInfosRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceInfosRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceInfosRequest", string(data)}, " ")
}
