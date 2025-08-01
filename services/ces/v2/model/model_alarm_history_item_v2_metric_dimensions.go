package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlarmHistoryItemV2MetricDimensions **参数解释**： 指标维度。 **取值范围**： 不涉及。
type AlarmHistoryItemV2MetricDimensions struct {

	// **参数解释**： 资源维度，如：弹性云服务器，则维度为instance_id；目前最大支持4个维度，各服务资源的指标维度名称可查看：“[服务维度名称](ces_03_0059.xml)”。 **取值范围**： 字符串长度在 1 到 32 之间。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 资源维度值，为资源的实例ID，如：4270ff17-aba3-4138-89fa-820594c39755。 **取值范围**： 字符串长度在 1 到 256 之间。
	Value *string `json:"value,omitempty"`
}

func (o AlarmHistoryItemV2MetricDimensions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmHistoryItemV2MetricDimensions struct{}"
	}

	return strings.Join([]string{"AlarmHistoryItemV2MetricDimensions", string(data)}, " ")
}
