package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMcpServerReq 创建MCP服务端对接配置请求。
type CreateMcpServerReq struct {

	// MCP服务端对接配置名称。
	Name string `json:"name"`

	// MCP服务端地址。
	McpServerUrl string `json:"mcp_server_url"`

	// 鉴权头域名称。
	AuthHeaderName *string `json:"auth_header_name,omitempty"`
}

func (o CreateMcpServerReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMcpServerReq struct{}"
	}

	return strings.Join([]string{"CreateMcpServerReq", string(data)}, " ")
}
