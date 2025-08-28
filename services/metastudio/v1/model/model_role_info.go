package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RoleInfo 角色信息。
type RoleInfo struct {

	// 角色ID。
	RoleId *string `json:"role_id,omitempty"`

	// 角色名称。
	Name *string `json:"name,omitempty"`

	// 角色描述。
	Description *string `json:"description,omitempty"`

	// 角色业务配置列表。
	RoleBusinessList *[]RoleBusinessInfo `json:"role_business_list,omitempty"`

	LlmSource *LlmSourceEnum `json:"llm_source,omitempty"`

	// 大语言模型配置ID。
	LlmConfigId *string `json:"llm_config_id,omitempty"`

	// 插件配置列表
	PluginConfigList *[]RolePluginConfigInfo `json:"plugin_config_list,omitempty"`

	// MCP服务端对接配置信息列表
	McpServerInfoList *[]McpServerBaseInfo `json:"mcp_server_info_list,omitempty"`

	// 指令集ID。
	InstructionLibraryId *string `json:"instruction_library_id,omitempty"`

	// 创建时间，格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间，格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o RoleInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RoleInfo struct{}"
	}

	return strings.Join([]string{"RoleInfo", string(data)}, " ")
}
