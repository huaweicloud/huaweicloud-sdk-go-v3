package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateImmutableRuleBody struct {

	// 优先级，默认值为0
	Priority *int32 `json:"priority,omitempty"`

	// 不可变策略是否生效，默认值为false
	Disabled *bool `json:"disabled,omitempty"`

	// 预留字段，支持填immutable
	Action *string `json:"action,omitempty"`

	// 预留字段，支持填immutable_template
	Template *string `json:"template,omitempty"`

	// 制品版本选择器，目前只支持长度为1
	TagSelectors []RuleSelector `json:"tag_selectors"`

	// 制品仓库选择器，目前只支持repository且长度为1
	ScopeSelectors map[string][]RuleSelector `json:"scope_selectors"`
}

func (o UpdateImmutableRuleBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateImmutableRuleBody struct{}"
	}

	return strings.Join([]string{"UpdateImmutableRuleBody", string(data)}, " ")
}
