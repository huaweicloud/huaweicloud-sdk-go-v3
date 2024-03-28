package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEnhancedConnectionsResponse Response Object
type ListEnhancedConnectionsResponse struct {

	// 执行请求是否成功。“true”表示请求执行成功。
	IsSuccess *bool `json:"is_success,omitempty"`

	// 系统提示信息，执行成功时，信息可能为空。
	Message *string `json:"message,omitempty"`

	// 跨源连接信息列表。
	Connections *[]EnhancedConnection `json:"connections,omitempty"`

	// 返回的跨源连接个数。
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListEnhancedConnectionsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEnhancedConnectionsResponse struct{}"
	}

	return strings.Join([]string{"ListEnhancedConnectionsResponse", string(data)}, " ")
}
