package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMcpServerResponse Response Object
type CreateMcpServerResponse struct {

	// MCP服务端对接配置ID。
	McpServerId *string `json:"mcp_server_id,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateMcpServerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMcpServerResponse struct{}"
	}

	return strings.Join([]string{"CreateMcpServerResponse", string(data)}, " ")
}
