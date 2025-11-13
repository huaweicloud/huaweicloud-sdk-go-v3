package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDrugJobStatisticsResponse Response Object
type ShowDrugJobStatisticsResponse struct {

	// **参数解释**： 药物设计作业数量统计结果。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	QuantityStatistics *[]QuantityStatistics `json:"quantity_statistics,omitempty"`

	// **参数解释**： 药物设计作业使用量统计结果。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	UsageStatistics *[]UsageStatistics `json:"usage_statistics,omitempty"`
	HttpStatusCode  int                `json:"-"`
}

func (o ShowDrugJobStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDrugJobStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ShowDrugJobStatisticsResponse", string(data)}, " ")
}
