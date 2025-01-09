package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTaskExecuteHistoryRequest Request Object
type ListTaskExecuteHistoryRequest struct {

	// 定时任务唯一标识。
	TaskId string `json:"task_id"`

	// 查询的偏移量。
	Offset *int32 `json:"offset,omitempty"`

	// 查询的数量，值区间[1-100]。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListTaskExecuteHistoryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTaskExecuteHistoryRequest struct{}"
	}

	return strings.Join([]string{"ListTaskExecuteHistoryRequest", string(data)}, " ")
}
