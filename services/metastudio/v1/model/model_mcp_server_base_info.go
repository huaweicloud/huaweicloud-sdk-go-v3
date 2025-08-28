package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// McpServerBaseInfo MCP服务端对接配置基础信息。
type McpServerBaseInfo struct {

	// MCP服务端对接配置ID。
	McpServerId string `json:"mcp_server_id"`

	// MCP服务端对接配置名称。
	Name *string `json:"name,omitempty"`
}

func (o McpServerBaseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "McpServerBaseInfo struct{}"
	}

	return strings.Join([]string{"McpServerBaseInfo", string(data)}, " ")
}
