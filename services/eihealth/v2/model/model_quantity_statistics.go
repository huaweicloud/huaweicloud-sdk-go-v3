package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// QuantityStatistics 小药物设计作业数量统计结果
type QuantityStatistics struct {

	// **参数解释**： 用户名。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	UserName *string `json:"user_name,omitempty"`

	// **参数解释**： 用户id。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	UserId *string `json:"user_id,omitempty"`

	// **参数解释**： 作业类型。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	JobType *string `json:"job_type,omitempty"`

	// **参数解释**： 作业状态。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Status *string `json:"status,omitempty"`

	// **参数解释**： 作业数量。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Count *int32 `json:"count,omitempty"`
}

func (o QuantityStatistics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuantityStatistics struct{}"
	}

	return strings.Join([]string{"QuantityStatistics", string(data)}, " ")
}
