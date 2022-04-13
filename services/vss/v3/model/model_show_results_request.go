package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowResultsRequest struct {
	// 任务ID

	TaskId string `json:"task_id"`
	// 分页查询，偏移量，表示从此偏移量开始查询

	Offset *int32 `json:"offset,omitempty"`
	// 分页查询，每页显示的条目数量

	Limit *int32 `json:"limit,omitempty"`
}

func (o ShowResultsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResultsRequest struct{}"
	}

	return strings.Join([]string{"ShowResultsRequest", string(data)}, " ")
}
