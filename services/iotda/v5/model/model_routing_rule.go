package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建或修改规则条件的响应消息
type RoutingRule struct {

	// 规则触发条件ID，用于唯一标识一个规则触发条件，在创建规则条件时由物联网平台分配获得。
	RuleId *string `json:"rule_id,omitempty" xml:"rule_id"`

	// 用户自定义的规则名称。
	RuleName *string `json:"rule_name,omitempty" xml:"rule_name"`

	// 用户自定义的规则描述。
	Description *string `json:"description,omitempty" xml:"description"`

	Subject *RoutingRuleSubject `json:"subject,omitempty" xml:"subject"`

	// 租户规则的生效范围，取值如下： - GLOBAL：生效范围为租户级 - APP：生效范围为资源空间级。
	AppType *string `json:"app_type,omitempty" xml:"app_type"`

	// 资源空间ID
	AppId *string `json:"app_id,omitempty" xml:"app_id"`

	// 用户自定义sql select语句，最大长度500，该参数仅供标准版和企业版用户使用。
	Select *string `json:"select,omitempty" xml:"select"`

	// 用户自定义sql where语句，最大长度500，该参数仅供标准版和企业版用户使用。
	Where *string `json:"where,omitempty" xml:"where"`

	// 规则条件的状态是否为激活。
	Active *bool `json:"active,omitempty" xml:"active"`
}

func (o RoutingRule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RoutingRule struct{}"
	}

	return strings.Join([]string{"RoutingRule", string(data)}, " ")
}
