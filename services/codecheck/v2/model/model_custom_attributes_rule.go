package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CustomAttributesRule struct {

	// 规则ID
	RuleId *string `json:"rule_id,omitempty" xml:"rule_id"`

	// attribute的问题级别，0致命，1严重，2一般，3提示
	Value *string `json:"value,omitempty" xml:"value"`

	// 规则阈值详细
	RuleConfigList *[]RuleConfig `json:"rule_config_list,omitempty" xml:"rule_config_list"`
}

func (o CustomAttributesRule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomAttributesRule struct{}"
	}

	return strings.Join([]string{"CustomAttributesRule", string(data)}, " ")
}
