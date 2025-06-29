package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ImmutableRule struct {

	// 不可变规则ID
	Id *int32 `json:"id,omitempty"`

	NamespaceId *int32 `json:"namespace_id,omitempty"`

	NamespaceName *string `json:"namespace_name,omitempty"`

	Priority *int32 `json:"priority,omitempty"`

	// 不可变规则是否生效
	Disabled *bool `json:"disabled,omitempty"`

	// 预留字段，仅支持填immutable
	Action *string `json:"action,omitempty"`

	// 预留字段，仅支持填immutable_template
	Template *string `json:"template,omitempty"`

	// 制品版本选择器，目前只支持长度为1
	TagSelectors *[]RuleSelector `json:"tag_selectors,omitempty"`

	// 制品仓库选择器，目前只支持repository且长度为1
	ScopeSelectors map[string][]RuleSelector `json:"scope_selectors,omitempty"`
}

func (o ImmutableRule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImmutableRule struct{}"
	}

	return strings.Join([]string{"ImmutableRule", string(data)}, " ")
}
