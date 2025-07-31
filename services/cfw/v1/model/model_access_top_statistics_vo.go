package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AccessTopStatisticsVo struct {

	// **参数解释**： 聚合时间 **取值范围**： 不涉及
	AggTime *int64 `json:"agg_time,omitempty"`

	// **参数解释**： 阻断数量 **取值范围**： 不涉及
	DenyAccessTopCounts *int64 `json:"deny_access_top_counts,omitempty"`

	// **参数解释**： 放行数量 **取值范围**： 不涉及
	PermitAccessTopCounts *int64 `json:"permit_access_top_counts,omitempty"`

	// **参数解释**： 命中次数 **取值范围**： 不涉及
	TotalAccessTopCounts *int64 `json:"total_access_top_counts,omitempty"`
}

func (o AccessTopStatisticsVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessTopStatisticsVo struct{}"
	}

	return strings.Join([]string{"AccessTopStatisticsVo", string(data)}, " ")
}
