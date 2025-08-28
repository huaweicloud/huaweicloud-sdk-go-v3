package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMcpServerResponse Response Object
type ShowMcpServerResponse struct {

	// MCP服务端对接配置ID。
	McpServerId *string `json:"mcp_server_id,omitempty"`

	// MCP服务端对接配置名称。
	Name *string `json:"name,omitempty"`

	// MCP服务端地址。
	McpServerUrl *string `json:"mcp_server_url,omitempty"`

	// 鉴权头域名称。
	AuthHeaderName *string `json:"auth_header_name,omitempty"`

	// 创建时间，格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间，格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	UpdateTime *string `json:"update_time,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowMcpServerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMcpServerResponse struct{}"
	}

	return strings.Join([]string{"ShowMcpServerResponse", string(data)}, " ")
}
