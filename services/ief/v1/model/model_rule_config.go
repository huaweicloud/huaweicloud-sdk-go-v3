package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 规则配置
type RuleConfig struct {

	// 规则描述，最大长度255，不允许^~#$%&*<>()[]{}'\"\\
	Description *string `json:"description,omitempty" xml:"description"`

	// 铂金版实例ID，如果为空则表示是专业版实例。
	IefInstanceId *string `json:"ief_instance_id,omitempty" xml:"ief_instance_id"`

	// 是否启用规则，默认为true（启用）
	InUsing *bool `json:"in_using,omitempty" xml:"in_using"`

	// 规则名称，只允许中文字符、英文字符、数字、下划线、中划线，最大长度64 同一个帐号中创建的规则名唯一
	Name string `json:"name" xml:"name"`

	// 源端点ID
	Source string `json:"source" xml:"source"`

	// 源端点资源。示例： - rest: {\"path\":\"<standard uri format>\"} - eventbus: {\"topic\":\"<project id>/nodes/<node id>/user/<租户自定义且满足eventbus topic要求的字符串>\",\"node_id\":\"<node id>\"}
	SourceResource map[string]string `json:"source_resource" xml:"source_resource"`

	// 目的端点ID
	Target string `json:"target" xml:"target"`

	// 目的端点资源。示例： - dis: {\"channel\": \"dis channel name\"} - servicebus: {\"path\": \"/request path\"} - apigw: {\"resource\": \"http://ssss.com\"} - eventbus: {\"topic\": \"/xxxx\"}
	TargetResource map[string]string `json:"target_resource" xml:"target_resource"`
}

func (o RuleConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleConfig struct{}"
	}

	return strings.Join([]string{"RuleConfig", string(data)}, " ")
}
