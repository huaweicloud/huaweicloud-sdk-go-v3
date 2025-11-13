package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UsageStatistics 小药物设计作业使用统计结果
type UsageStatistics struct {

	// **参数解释**： 用户名。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	UserName *string `json:"user_name,omitempty"`

	// **参数解释**： 用户id。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	UserId *string `json:"user_id,omitempty"`

	// **参数解释**： 作业类型。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	JobType *string `json:"job_type,omitempty"`

	// **参数解释**： 作业使用量。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	ChargeSum *float64 `json:"charge_sum,omitempty"`
}

func (o UsageStatistics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UsageStatistics struct{}"
	}

	return strings.Join([]string{"UsageStatistics", string(data)}, " ")
}
