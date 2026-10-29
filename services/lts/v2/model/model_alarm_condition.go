package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AlarmCondition struct {

	// **参数解释：** 触发告警的条件表达式。 **约束限制：** - 不支持纯数字。 - 不支持以数字开头的。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	ConditionExpression string `json:"condition_expression"`

	// **参数解释：** 触发的告警级别。 **约束限制：** 不支持纯数字，不支持以数字开头的。 **取值范围：** - Info：提示 - Minor：次要 - Major：重要 - Critical：紧急 **默认取值：** 不涉及。
	AlarmLevel string `json:"alarm_level"`
}

func (o AlarmCondition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmCondition struct{}"
	}

	return strings.Join([]string{"AlarmCondition", string(data)}, " ")
}
