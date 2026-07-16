package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PoolMonitorDataPoints struct {

	// **参数解释**：指标数据时间戳，以毫秒为单位。 **取值范围**：不涉及。
	Timestamp *int32 `json:"timestamp,omitempty"`

	// **参数解释**：指标数据单位。 **取值范围**：可选值如下： - Percent：百分比。 - Megabytes：兆字节。
	Unit *string `json:"unit,omitempty"`

	// **参数解释**：指标数据值。
	Statistics *[]PoolMonitorStatistics `json:"statistics,omitempty"`
}

func (o PoolMonitorDataPoints) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoolMonitorDataPoints struct{}"
	}

	return strings.Join([]string{"PoolMonitorDataPoints", string(data)}, " ")
}
