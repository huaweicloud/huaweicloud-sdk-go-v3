package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HardwareMetric **参数解释**： 监控指标信息 **约束限制**： 不涉及
type HardwareMetric struct {

	// **参数解释**： 监控指标名 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	MetricName *string `json:"metric_name,omitempty"`

	// **参数解释**： 监控指标对象 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	MetricDevices *[]string `json:"metric_devices,omitempty"`
}

func (o HardwareMetric) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HardwareMetric struct{}"
	}

	return strings.Join([]string{"HardwareMetric", string(data)}, " ")
}
