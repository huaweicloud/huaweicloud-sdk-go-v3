package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TrafficTrendVo struct {

	// **参数解释**： 聚合时间点 **取值范围**： 不涉及
	AggTime *int64 `json:"agg_time,omitempty"`

	// **参数解释**： bps **取值范围**： 不涉及
	Bps *float64 `json:"bps,omitempty"`

	// **参数解释**： 入方向bps **取值范围**： 不涉及
	InBps *float64 `json:"in_bps,omitempty"`

	// **参数解释**： 出方向bps **取值范围**： 不涉及
	OutBps *float64 `json:"out_bps,omitempty"`
}

func (o TrafficTrendVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TrafficTrendVo struct{}"
	}

	return strings.Join([]string{"TrafficTrendVo", string(data)}, " ")
}
