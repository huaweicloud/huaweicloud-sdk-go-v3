package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// McpServerInfo MCP服务端对接配置基本信息。
type McpServerInfo struct {

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
}

func (o McpServerInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "McpServerInfo struct{}"
	}

	return strings.Join([]string{"McpServerInfo", string(data)}, " ")
}
