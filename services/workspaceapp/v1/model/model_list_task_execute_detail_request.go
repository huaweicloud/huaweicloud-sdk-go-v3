package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTaskExecuteDetailRequest Request Object
type ListTaskExecuteDetailRequest struct {

	// 定时任务唯一标识。
	ExecuteHistoryId string `json:"execute_history_id"`

	// 查询的偏移量。
	Offset *int32 `json:"offset,omitempty"`

	// 查询的数量，值区间[1-100]。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListTaskExecuteDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTaskExecuteDetailRequest struct{}"
	}

	return strings.Join([]string{"ListTaskExecuteDetailRequest", string(data)}, " ")
}
