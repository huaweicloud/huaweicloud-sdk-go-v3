package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkItemFlowNodeConfigVo 工作流节点配置
type WorkItemFlowNodeConfigVo struct {

	// 节点编码
	Code *string `json:"code,omitempty"`

	// 节点名称
	Name *string `json:"name,omitempty"`

	// 节点描述
	Description *string `json:"description,omitempty"`

	// 是否为结束节点
	End *bool `json:"end,omitempty"`

	// 是否为最末节点
	Last *bool `json:"last,omitempty"`

	// 是否为开始节点
	Start *bool `json:"start,omitempty"`

	// 是否允许挂起
	EnableSuspend *bool `json:"enable_suspend,omitempty"`

	// 节点扩展配置
	ExtraConfig map[string]interface{} `json:"extra_config,omitempty"`

	// 静态规则列表
	StaticRules *[]map[string]interface{} `json:"static_rules,omitempty"`

	// 静态动作配置
	StaticActions map[string]interface{} `json:"static_actions,omitempty"`

	// 是否任意状态可流转
	AnyStatus *bool `json:"any_status,omitempty"`

	// 提交时是否可操作
	SubmitCanOperate *bool `json:"submit_can_operate,omitempty"`
}

func (o WorkItemFlowNodeConfigVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkItemFlowNodeConfigVo struct{}"
	}

	return strings.Join([]string{"WorkItemFlowNodeConfigVo", string(data)}, " ")
}
