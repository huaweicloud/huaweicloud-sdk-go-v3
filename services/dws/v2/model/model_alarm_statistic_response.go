package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlarmStatisticResponse **参数解释**： 告警统计。 **取值范围**： 不涉及。
type AlarmStatisticResponse struct {

	// **参数解释**： 日期。 **取值范围**： 不涉及。
	Date *string `json:"date,omitempty"`

	// **参数解释**： 紧急。 **取值范围**： 不涉及。
	Urgent *string `json:"urgent,omitempty"`

	// **参数解释**： 重要。 **取值范围**： 不涉及。
	Important *string `json:"important,omitempty"`

	// **参数解释**： 次要。 **取值范围**： 不涉及。
	Minor *string `json:"minor,omitempty"`

	// **参数解释**： 提示。 **取值范围**： 不涉及。
	Prompt *string `json:"prompt,omitempty"`
}

func (o AlarmStatisticResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmStatisticResponse struct{}"
	}

	return strings.Join([]string{"AlarmStatisticResponse", string(data)}, " ")
}
