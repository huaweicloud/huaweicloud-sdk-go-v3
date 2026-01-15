package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SummaryInfo struct {

	// **参数解释**： 集群风险概要。 **取值范围**： 不涉及
	HighRisk *int32 `json:"highRisk,omitempty"`

	// **参数解释**： 集群风险检测项判定为中风险的数量。 **取值范围**： 0-100
	MediumRisk *int32 `json:"mediumRisk,omitempty"`

	// **参数解释： 集群风险检测项判定为正常项的数量。 **取值范围**： 0-100
	Normal *int32 `json:"normal,omitempty"`

	// **参数解释**： 集群风险检测项判定为提示项的数量。 **取值范围**： 0-100
	Suggestion *int32 `json:"suggestion,omitempty"`
}

func (o SummaryInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SummaryInfo struct{}"
	}

	return strings.Join([]string{"SummaryInfo", string(data)}, " ")
}
