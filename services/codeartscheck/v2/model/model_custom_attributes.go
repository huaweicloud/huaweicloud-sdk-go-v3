package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CustomAttributes struct {

	// **参数解释:** 配置项属性，severity：为规则告警等级。 **约束限制:** 不涉及。 **取值范围:** severity。 **默认取值:** 不涉及。
	Attribute *string `json:"attribute,omitempty"`

	// **参数解释:** 规则详细。
	Rules *[]CustomAttributesRule `json:"rules,omitempty"`
}

func (o CustomAttributes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomAttributes struct{}"
	}

	return strings.Join([]string{"CustomAttributes", string(data)}, " ")
}
