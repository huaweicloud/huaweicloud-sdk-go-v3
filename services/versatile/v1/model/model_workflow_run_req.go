package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WorkflowRunReq struct {
	Inputs map[string]interface{} `json:"inputs"`

	MemoryInputs map[string]interface{} `json:"memory_inputs,omitempty"`

	Globals map[string]interface{} `json:"globals,omitempty"`

	// user消息
	Messages *[]Message `json:"messages,omitempty"`

	UserProfile *UserProfile `json:"user_profile,omitempty"`

	// 插件参数数组
	PluginConfigs *[]PluginConfig `json:"plugin_configs,omitempty"`

	Version *int64 `json:"version,omitempty"`

	// 用户ID
	UserId *string `json:"userId,omitempty"`

	Conversation *Conversation `json:"conversation,omitempty"`

	// 会话历史写入
	EnableHistory *bool `json:"enable_history,omitempty"`
}

func (o WorkflowRunReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkflowRunReq struct{}"
	}

	return strings.Join([]string{"WorkflowRunReq", string(data)}, " ")
}
