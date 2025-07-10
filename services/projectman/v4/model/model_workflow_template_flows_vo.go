package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkflowTemplateFlowsVo 状态流中流转线信息
type WorkflowTemplateFlowsVo struct {

	// 流转线编码
	Code *string `json:"code,omitempty"`

	// 流转线名称
	Name *string `json:"name,omitempty"`

	// 流转前校验配置
	BeforeRuleValidator *[]WorkflowTemplateConfigsVo `json:"before_rule_validator,omitempty"`

	// 流转中界面配置
	BeforeRuleConfigs *[]WorkflowTemplateConfigsVo `json:"before_rule_configs,omitempty"`

	// 流转线的入口状态
	FromCode *string `json:"from_code,omitempty"`

	// 流转线的出口状态
	ToCode *string `json:"to_code,omitempty"`
}

func (o WorkflowTemplateFlowsVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkflowTemplateFlowsVo struct{}"
	}

	return strings.Join([]string{"WorkflowTemplateFlowsVo", string(data)}, " ")
}
