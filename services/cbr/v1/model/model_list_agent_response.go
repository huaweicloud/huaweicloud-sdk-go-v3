package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAgentResponse Response Object
type ListAgentResponse struct {

	// 客户端实例列表
	Agents *[]Agent `json:"agents,omitempty"`

	// 客户端个数
	Count *int32 `json:"count,omitempty"`

	// 每页显示的条目数量
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，表示从此偏移量开始查询
	Offset         *int32 `json:"offset,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAgentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAgentResponse struct{}"
	}

	return strings.Join([]string{"ListAgentResponse", string(data)}, " ")
}
