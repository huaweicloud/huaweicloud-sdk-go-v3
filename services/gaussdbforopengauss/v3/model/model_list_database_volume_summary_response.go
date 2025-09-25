package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDatabaseVolumeSummaryResponse Response Object
type ListDatabaseVolumeSummaryResponse struct {

	// **参数解释**: 数据盘总量。 **取值范围**: 不涉及。
	DataDiskCapacity *string `json:"data_disk_capacity,omitempty"`

	// **参数解释**: 数据盘使用量。 **取值范围**: 不涉及。
	DataDiskUsage *string `json:"data_disk_usage,omitempty"`

	// **参数解释**: 空间使用日均增长量。 **取值范围**: 不涉及。
	SpaceUsageGrowthPerDay *string `json:"space_usage_growth_per_day,omitempty"`

	// **参数解释**: 预计可用天数。 **取值范围**: 不涉及。
	EstimatedRemainingDays *string `json:"estimated_remaining_days,omitempty"`

	// **参数解释**: CN节点信息。
	CnComponents *[]CnComponentInfoResult `json:"cn_components,omitempty"`

	// **参数解释**: DN节点信息。
	DnComponents   *[]DnComponentInfoResult `json:"dn_components,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ListDatabaseVolumeSummaryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDatabaseVolumeSummaryResponse struct{}"
	}

	return strings.Join([]string{"ListDatabaseVolumeSummaryResponse", string(data)}, " ")
}
