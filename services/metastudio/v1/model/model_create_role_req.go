package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRoleReq 创建角色请求。
type CreateRoleReq struct {

	// 角色名称。
	Name string `json:"name"`

	// 角色描述。
	Description string `json:"description"`

	// 角色业务配置列表
	RoleBusinessList []RoleBusinessReq `json:"role_business_list"`

	LlmSource *LlmSourceEnum `json:"llm_source"`

	// 大语言模型配置ID。
	LlmConfigId *string `json:"llm_config_id,omitempty"`

	// 插件配置列表
	PluginConfigList *[]RolePluginConfigInfo `json:"plugin_config_list,omitempty"`

	// MCP服务端对接配置ID列表
	McpServerIdList *[]string `json:"mcp_server_id_list,omitempty"`

	// 指令集ID。
	InstructionLibraryId *string `json:"instruction_library_id,omitempty"`
}

func (o CreateRoleReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRoleReq struct{}"
	}

	return strings.Join([]string{"CreateRoleReq", string(data)}, " ")
}
