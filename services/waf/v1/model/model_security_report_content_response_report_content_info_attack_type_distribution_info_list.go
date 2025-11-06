package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SecurityReportContentResponseReportContentInfoAttackTypeDistributionInfoList struct {

	// **参数解释：** 攻击类型标识（如custom_custom表示精准防护类攻击）。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Key *string `json:"key,omitempty"`

	// **参数解释：** 该攻击类型的总数量。 **约束限制：** 不涉及 **取值范围：** ≥0 **默认取值：** 0
	Num *int32 `json:"num,omitempty"`
}

func (o SecurityReportContentResponseReportContentInfoAttackTypeDistributionInfoList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityReportContentResponseReportContentInfoAttackTypeDistributionInfoList struct{}"
	}

	return strings.Join([]string{"SecurityReportContentResponseReportContentInfoAttackTypeDistributionInfoList", string(data)}, " ")
}
