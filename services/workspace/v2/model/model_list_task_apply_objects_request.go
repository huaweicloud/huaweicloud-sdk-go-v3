package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTaskApplyObjectsRequest Request Object
type ListTaskApplyObjectsRequest struct {

	// 任务ID（精确查询）
	TaskId string `json:"task_id"`

	// 应用对象名称（支持模糊查询）
	ObjectName *string `json:"object_name,omitempty"`

	// 偏移量，默认0
	Offset *int32 `json:"offset,omitempty"`

	// 每页数量，默认10，最大100
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListTaskApplyObjectsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTaskApplyObjectsRequest struct{}"
	}

	return strings.Join([]string{"ListTaskApplyObjectsRequest", string(data)}, " ")
}
