package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMcpServerResponse Response Object
type ListMcpServerResponse struct {

	// 与第一条数据的偏移量
	Offset *int32 `json:"offset,omitempty"`

	// 页面大小
	Limit *int32 `json:"limit,omitempty"`

	// 总数量
	Count *int32 `json:"count,omitempty"`

	// MCP服务端对接配置信息
	Data *[]McpServerInfo `json:"data,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListMcpServerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMcpServerResponse struct{}"
	}

	return strings.Join([]string{"ListMcpServerResponse", string(data)}, " ")
}
