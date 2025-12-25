package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AgentRunReq 知识型Agent执行请求体
type AgentRunReq struct {

	// 请求问题
	Query string `json:"query"`

	Inputs map[string]interface{} `json:"inputs,omitempty"`

	UserProfile *UserProfile `json:"user_profile,omitempty"`

	// 插件是否开启
	ToolSwitchDict map[string]bool `json:"tool_switch_dict,omitempty"`

	// 模型ID
	ModelDeploymentId *string `json:"model_deployment_id,omitempty"`

	// 是否记录会话历史
	EnableHistory *bool `json:"enable_history,omitempty"`

	// 传入的会话历史
	Histories *[]Message `json:"histories,omitempty"`

	// 上传文件列表
	Files *[]string `json:"files,omitempty"`
}

func (o AgentRunReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgentRunReq struct{}"
	}

	return strings.Join([]string{"AgentRunReq", string(data)}, " ")
}
