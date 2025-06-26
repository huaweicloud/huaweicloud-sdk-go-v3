package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClusterMetric **参数解释**： 指标详情。 **取值范围**： 不涉及。
type ClusterMetric struct {

	// **参数解释**： 指标名称。 **取值范围**： 不涉及。
	Scope *string `json:"scope,omitempty"`

	// **参数解释**： 指标表相关字段信息。 **取值范围**： 不涉及。
	Fields *[]SimpleFieldDto `json:"fields,omitempty"`

	// **参数解释**： 作用域。 **取值范围**： 不涉及。
	MetricName *string `json:"metric_name,omitempty"`

	// **参数解释**： 采集速率。 **取值范围**： 不涉及。
	CollectRate *int32 `json:"collect_rate,omitempty"`

	// **参数解释**： 采集时间范围。 **取值范围**： 不涉及。
	CollectRange *[]string `json:"collect_range,omitempty"`

	// **参数解释**： 创建时间。 **取值范围**： 不涉及。
	CreateTime *string `json:"create_time,omitempty"`
}

func (o ClusterMetric) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterMetric struct{}"
	}

	return strings.Join([]string{"ClusterMetric", string(data)}, " ")
}
