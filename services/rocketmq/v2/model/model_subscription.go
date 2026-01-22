package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Subscription struct {

	// **参数解释**： 订阅的Topic名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Topic *string `json:"topic,omitempty"`

	// **参数解释**： 订阅类型。 **约束限制**： 不涉及。 **取值范围**： - TAG：基于TAG进行订阅。 - SQL92：基于消息属性进行复杂条件过滤的订阅。 **默认取值**： 不涉及。
	Type *string `json:"type,omitempty"`

	// 订阅tag字符。
	Expression *string `json:"expression,omitempty"`
}

func (o Subscription) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Subscription struct{}"
	}

	return strings.Join([]string{"Subscription", string(data)}, " ")
}
