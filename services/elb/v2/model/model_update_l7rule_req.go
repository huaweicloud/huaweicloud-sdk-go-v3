package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// L7转发策略请求实体
type UpdateL7ruleReq struct {
	// 转发匹配方式： type为HOST_NAME时，取值范围：EQUAL_TO：精确匹配；t ype为PATH时，取值范围：REGEX：正则匹配；STARTS_WITH：前缀匹配；EQUAL_TO：精确匹配。

	CompareType *string `json:"compare_type,omitempty"`
	// 转发规则的管理状态；取值范围： true/false。该字段为预留字段，暂未启用。默认为true。

	AdminStateUp *bool `json:"admin_state_up,omitempty"`
	// 是否反向匹配；取值范围：true/false。默认值：false；该字段为预留字段，暂未启用。

	Invert *bool `json:"invert,omitempty"`
	// 匹配内容的键值。默认为null。该字段为预留字段，暂未启用。

	Key *string `json:"key,omitempty"`
	// 匹配内容的值。不能包含空格。 当type为HOST_NAME时，取值范围：String (100)，字符串只能包含英文字母、数字、“-”或“.”，且必须以字母或数字开头。 当type为PATH时，取值范围：String (128)。当转发规则的compare_type为STARTS_WITH、EQUAL_TO时，字符串只能包含英文字母、数字、_~';@^-%#&$.*+?,=!:| /()[]{}，且必须以\"/\"开头。

	Value *string `json:"value,omitempty"`
}

func (o UpdateL7ruleReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateL7ruleReq struct{}"
	}

	return strings.Join([]string{"UpdateL7ruleReq", string(data)}, " ")
}
