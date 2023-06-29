package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Action 工作流节点定义
type Action struct {

	// 节点名称
	ActionName string `json:"action_name"`

	// 节点使用的委托
	ActionAgency string `json:"action_agency"`

	// 节点相关联的函数URN
	Function string `json:"function"`

	// 算子模板使用的URM
	FunctionTemplate string `json:"function_template"`

	// 节点使用的算子名称
	ActionTemplateName *string `json:"action_template_name,omitempty"`

	// 节点使用的模板类别
	ActionTemplateCategory *string `json:"action_template_category,omitempty"`

	// 节点使用的模板提供方
	ActionTemplateProviderName *string `json:"action_template_provider_name,omitempty"`

	// 触发模式
	InvocationMode string `json:"invocation_mode"`

	// 超时时间
	Timeout int32 `json:"timeout"`

	// 动态参数与inputs参数相关联使用的filter。默认为\"$\"
	PayloadFilter *string `json:"payload_filter,omitempty"`

	// 节点使用的动态参数
	DynamicSource map[string]interface{} `json:"dynamic_source"`

	// action错误处理
	Results *[]ActionResult `json:"results,omitempty"`
}

func (o Action) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Action struct{}"
	}

	return strings.Join([]string{"Action", string(data)}, " ")
}
