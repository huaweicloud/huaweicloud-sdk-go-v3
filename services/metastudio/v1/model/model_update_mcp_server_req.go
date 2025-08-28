package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateMcpServerReq 修改MCP服务端对接配置请求。
type UpdateMcpServerReq struct {

	// MCP服务端对接配置名称。
	Name *string `json:"name,omitempty"`

	// MCP服务端地址。
	McpServerUrl *string `json:"mcp_server_url,omitempty"`

	// 鉴权头域名称。
	AuthHeaderName *string `json:"auth_header_name,omitempty"`
}

func (o UpdateMcpServerReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMcpServerReq struct{}"
	}

	return strings.Join([]string{"UpdateMcpServerReq", string(data)}, " ")
}
