package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAssetEditTaskRequest Request Object
type ListAssetEditTaskRequest struct {

	// 编辑任务ID
	TaskId *string `json:"task_id,omitempty"`

	// 查询偏移量，默认0
	Offset *int32 `json:"offset,omitempty"`

	// 查询一页返回数，默认10，最大支持100
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListAssetEditTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAssetEditTaskRequest struct{}"
	}

	return strings.Join([]string{"ListAssetEditTaskRequest", string(data)}, " ")
}
