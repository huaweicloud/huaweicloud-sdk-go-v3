package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SloInfo SLO配置信息
type SloInfo struct {

	// **参数解释：** 仿真期望指标。 **取值范围：** 不涉及。
	SloObjectives []SloObjectives `json:"slo_objectives"`

	// **参数解释：** 预测时间窗口。 **约束限制：** 60~600。 **取值范围：** 不涉及。 **默认取值：** 60。
	PredictWindowSeconds *int32 `json:"predict_window_seconds,omitempty"`
}

func (o SloInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SloInfo struct{}"
	}

	return strings.Join([]string{"SloInfo", string(data)}, " ")
}
