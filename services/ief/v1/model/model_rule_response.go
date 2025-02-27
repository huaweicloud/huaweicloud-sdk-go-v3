package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RuleResponse 规则配置
type RuleResponse struct {

	// 创建时间
	CreatedAt *string `json:"created_at,omitempty"`

	// 规则描述，最大长度255，不允许^~#$%&*<>()[]{}'\"\\
	Description *string `json:"description,omitempty"`

	// 转发失败的消息数
	FailMessages *int32 `json:"fail_messages,omitempty"`

	// 规则ID
	Id *string `json:"id,omitempty"`

	// [铂金版实例ID，如果为空则表示是专业版实例。](tag:hws,hws_hk)[铂金版实例ID](tag:hcs,hcs_sm)
	IefInstanceId *string `json:"ief_instance_id,omitempty"`

	// 是否启用规则，默认为true（启用）
	InUsing *bool `json:"in_using,omitempty"`

	// 规则名称，只允许中文字符、英文字符、数字、下划线、中划线，最大长度64 同一个账号中创建的规则名唯一
	Name string `json:"name"`

	// 项目ID
	ProjectId string `json:"project_id"`

	Source *EndpointObjResp `json:"source"`

	// 源端点资源。示例： - rest: {\"path\":\"\\<standard uri format\\>\"} - eventbus: {\"topic\":\"\\<project id\\>/nodes/\\<node id\\>/user/\\<租户自定义且满足eventbus topic要求的字符串\\>\",\"node_id\":\"\\<node id\\>\"}
	SourceResource map[string]string `json:"source_resource"`

	Target *EndpointObjResp `json:"target"`

	// 目的端点资源，示例： - dis: {\"channel\": \"dis channel name\"} - servicebus: {\"path\": \"/request path\"} - apigw: {\"resource\": \"http://ssss.com\"} - eventbus: {\"topic\": \"/xxxx\"}
	TargetResource map[string]string `json:"target_resource"`

	// 更新时间
	UpdatedAt string `json:"updated_at"`

	// 转发成功的消息数
	SuccessMessages int32 `json:"success_messages"`
}

func (o RuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleResponse struct{}"
	}

	return strings.Join([]string{"RuleResponse", string(data)}, " ")
}
