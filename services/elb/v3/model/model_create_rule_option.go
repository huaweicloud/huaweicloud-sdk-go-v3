package model

import (
	"encoding/json"

	"strings"
)

// 创建L7转发规则器请求
type CreateRuleOption struct {
	// 转发规则的管理状态；该字段为预留字段，暂未启用。默认为true。

	AdminStateUp *bool `json:"admin_state_up,omitempty"`
	// 转发规则的匹配方式。type为HOST_NAME时可以为EQUAL_TO。type为PATH时可以为REGEX， STARTS_WITH，EQUAL_TO。

	CompareType string `json:"compare_type"`
	// 匹配内容的键值。目前匹配内容为HOST_NAME和PATH时，该字段不生效。该字段能更新但不会生效。

	Key *string `json:"key,omitempty"`
	// 转发规则所在的项目ID。

	ProjectId *string `json:"project_id,omitempty"`
	// 一个l7policy下创建的l7rule的type不能重复。 匹配内容：可以为HOST_NAME，PATH

	Type string `json:"type"`
	// 匹配内容的值。不能包含空格。当type为HOST_NAME时，取值范围：String (100)，字符串只能包含英文字母、数字、“-”或“.”，且必须以字母或数字开头。当type为PATH时，取值范围：String (128)。当转发规则的compare_type为STARTS_WITH、EQUAL_TO时，字符串只能包含英文字母、数字、_~';@^-%#&$.*+?,=!:| /()[]{}，且必须以\"/\"开头。

	Value string `json:"value"`
	// 是否反向匹配；取值范围：true/false。默认值：false；该字段为预留字段，暂未启用。

	Invert *bool `json:"invert,omitempty"`
}

func (o CreateRuleOption) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateRuleOption struct{}"
	}

	return strings.Join([]string{"CreateRuleOption", string(data)}, " ")
}
