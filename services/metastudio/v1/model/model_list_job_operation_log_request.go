package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListJobOperationLogRequest Request Object
type ListJobOperationLogRequest struct {

	// 偏移量，表示从此偏移量开始查询。
	Offset *int32 `json:"offset,omitempty"`

	// 每页显示的条目数量。
	Limit *int32 `json:"limit,omitempty"`

	// 任务id
	JobId string `json:"job_id"`
}

func (o ListJobOperationLogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJobOperationLogRequest struct{}"
	}

	return strings.Join([]string{"ListJobOperationLogRequest", string(data)}, " ")
}
