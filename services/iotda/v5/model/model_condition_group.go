package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 规则条件组
type ConditionGroup struct {

	// **参数说明**：规则的条件列表，单个规则最多支持设置10个条件。
	Conditions *[]RuleCondition `json:"conditions,omitempty" xml:"conditions"`

	// **参数说明**：规则条件列表中多个条件之间的逻辑关系，默认值：and。 **取值范围**： - and：逻辑且。 - or：逻辑或。
	Logic *string `json:"logic,omitempty" xml:"logic"`

	TimeRange *TimeRange `json:"time_range,omitempty" xml:"time_range"`
}

func (o ConditionGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConditionGroup struct{}"
	}

	return strings.Join([]string{"ConditionGroup", string(data)}, " ")
}
