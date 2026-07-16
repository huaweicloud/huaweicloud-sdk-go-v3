package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SloObjectives SLO目标配置
type SloObjectives struct {

	// **参数解释：** TTFT指标，单位毫秒。 **取值范围：** 0~10000。 **约束限制：** 不涉及。 **默认取值：** 100。
	MetricTtft *int32 `json:"metric_ttft,omitempty"`

	// **参数解释：** TPOT指标，单位毫秒。 **取值范围：** 0~1000。 **约束限制：** 不涉及。 **默认取值：** 50。
	MetricTpot *int32 `json:"metric_tpot,omitempty"`

	// **参数解释：** SLO满足百分比。 **取值范围：** 0~100。
	Percental *int32 `json:"percental,omitempty"`
}

func (o SloObjectives) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SloObjectives struct{}"
	}

	return strings.Join([]string{"SloObjectives", string(data)}, " ")
}
