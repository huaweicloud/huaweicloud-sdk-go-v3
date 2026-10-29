package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type KeywordsAlarmConditions struct {

	// **参数解释：** 触发条件，即满足该条件时，会触发告警。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Condition string `json:"condition"`

	// **参数解释：** 匹配条数。当关键词搜索结果的日志条数达到设定的条数时，会触发告警。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Number int32 `json:"number"`

	// **参数解释：** 告警级别。 **约束限制：** 不涉及。 **取值范围：** -   Info：提示 -   Minor：次要 -   Major：重要 -   Critical：紧急 **默认取值：** 不涉及。
	AlarmLevel string `json:"alarm_level"`
}

func (o KeywordsAlarmConditions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KeywordsAlarmConditions struct{}"
	}

	return strings.Join([]string{"KeywordsAlarmConditions", string(data)}, " ")
}
