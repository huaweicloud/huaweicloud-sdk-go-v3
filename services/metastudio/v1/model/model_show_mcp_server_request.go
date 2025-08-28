package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMcpServerRequest Request Object
type ShowMcpServerRequest struct {

	// 使用AK/SK方式认证时必选，携带的鉴权信息。
	Authorization *string `json:"Authorization,omitempty"`

	// 使用AK/SK方式认证时必选，请求的发生时间。
	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	// 使用AK/SK方式认证时必选，携带项目ID信息。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 第三方用户ID。不允许输入中文。
	XAppUserId *string `json:"X-App-UserId,omitempty"`

	// MCP服务端对接配置ID。
	McpServerId string `json:"mcp_server_id"`
}

func (o ShowMcpServerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMcpServerRequest struct{}"
	}

	return strings.Join([]string{"ShowMcpServerRequest", string(data)}, " ")
}
