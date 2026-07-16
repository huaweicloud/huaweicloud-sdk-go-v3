package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PoolMonitorStatistics struct {

	// **参数解释**：资源池监控信息在指定时间粒度下的统计方式。 **取值范围**：可选值如下： - maximum：最大值统计，默认值。 - minimum：最小值统计。 - sum：求和统计。 - average：平均值统计。 - sampleCount：采样统计。
	Statistic *string `json:"statistic,omitempty"`

	// **参数解释**：指标数据的值，值为-1时表示无该指标数据。 **取值范围**：不涉及。
	Value *float32 `json:"value,omitempty"`
}

func (o PoolMonitorStatistics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoolMonitorStatistics struct{}"
	}

	return strings.Join([]string{"PoolMonitorStatistics", string(data)}, " ")
}
