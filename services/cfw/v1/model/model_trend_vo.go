package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TrendVo struct {

	// **参数解释**： 聚合时间 **取值范围**： 不涉及
	AggTime *int64 `json:"agg_time,omitempty"`

	// **参数解释**： 带宽 **取值范围**： 不涉及
	Bps *float64 `json:"bps,omitempty"`

	// **参数解释**： 阻断数量 **取值范围**： 不涉及
	Deny *int64 `json:"deny,omitempty"`

	// **参数解释**： 入方向bps **取值范围**： 不涉及
	InBps *float64 `json:"in_bps,omitempty"`

	// **参数解释**： 出方向bps **取值范围**： 不涉及
	OutBps *float64 `json:"out_bps,omitempty"`

	// **参数解释**： 允许数量 **取值范围**： 不涉及
	Permit *int64 `json:"permit,omitempty"`
}

func (o TrendVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TrendVo struct{}"
	}

	return strings.Join([]string{"TrendVo", string(data)}, " ")
}
