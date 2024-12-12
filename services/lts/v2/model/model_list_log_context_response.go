package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLogContextResponse Response Object
type ListLogContextResponse struct {

	// 上下文日志信息。
	Logs *[]LogContents `json:"logs,omitempty"`

	// 返回的总日志条数，包含请求参数中所指定的起始日志。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 向前查询到的日志条数。
	BackwardsCount *int32 `json:"backwards_count,omitempty"`

	// 向后查询到的日志条数。
	ForwardsCount *int32 `json:"forwards_count,omitempty"`

	// 是否查询完成。
	IsQueryComplete *bool `json:"isQueryComplete,omitempty"`
	HttpStatusCode  int   `json:"-"`
}

func (o ListLogContextResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLogContextResponse struct{}"
	}

	return strings.Join([]string{"ListLogContextResponse", string(data)}, " ")
}
