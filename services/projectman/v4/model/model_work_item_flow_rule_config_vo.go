package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkItemFlowRuleConfigVo 工作项流转规则配置
type WorkItemFlowRuleConfigVo struct {

	// 规则编码
	Code *string `json:"code,omitempty"`

	// 规则开关
	Open *bool `json:"open,omitempty"`

	// 字段配置值列表
	ConfigValue *[]WorkItemFlowFieldConfigVo `json:"config_value,omitempty"`
}

func (o WorkItemFlowRuleConfigVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkItemFlowRuleConfigVo struct{}"
	}

	return strings.Join([]string{"WorkItemFlowRuleConfigVo", string(data)}, " ")
}
