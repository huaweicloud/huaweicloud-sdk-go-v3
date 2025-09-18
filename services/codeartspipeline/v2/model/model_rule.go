package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Rule struct {

	// **参数解释**： 规则ID。 **取值范围**： 不涉及。
	Id string `json:"id"`

	// **参数解释**： 规则类型。 **取值范围**： 不涉及。
	Type string `json:"type"`

	// **参数解释**： 规则名称。 **取值范围**： 不涉及。
	Name string `json:"name"`

	// **参数解释**： 规则版本。 **取值范围**： 不涉及。
	Version string `json:"version"`

	// **参数解释**： 最近操作人员。 **取值范围**： 不涉及。
	Operator string `json:"operator"`

	// **参数解释**： 最近操作时间。 **取值范围**： 不涉及。
	OperateTime int64 `json:"operate_time"`
}

func (o Rule) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Rule struct{}"
	}

	return strings.Join([]string{"Rule", string(data)}, " ")
}
