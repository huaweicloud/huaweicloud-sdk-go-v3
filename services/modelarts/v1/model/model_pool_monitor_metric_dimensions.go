package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PoolMonitorMetricDimensions struct {

	// **参数解释**：指标维度名称。 **取值范围**：指标名称。可选值如下： - clusterId：集群ID。
	Name *string `json:"name,omitempty"`

	// **参数解释**：指标维度取值。 **取值范围**：不涉及。
	Value *string `json:"value,omitempty"`
}

func (o PoolMonitorMetricDimensions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoolMonitorMetricDimensions struct{}"
	}

	return strings.Join([]string{"PoolMonitorMetricDimensions", string(data)}, " ")
}
