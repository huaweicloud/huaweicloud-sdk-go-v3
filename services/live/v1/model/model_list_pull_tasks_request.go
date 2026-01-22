package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPullTasksRequest Request Object
type ListPullTasksRequest struct {

	// 任务所在区域
	Region *string `json:"region,omitempty"`

	// 偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 每页记录数
	Limit *int32 `json:"limit,omitempty"`

	// 任务id
	TaskId *string `json:"task_id,omitempty"`
}

func (o ListPullTasksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPullTasksRequest struct{}"
	}

	return strings.Join([]string{"ListPullTasksRequest", string(data)}, " ")
}
