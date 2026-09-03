package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetInstancesOpsMetricNamesRequest Request Object
type GetInstancesOpsMetricNamesRequest struct {

	// **参数解释**：  实例ID，此参数是实例的唯一标识。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，长度为36个字符。  **默认取值**：  不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**：  监控指标分组名称。  **约束限制**：  不涉及。  **取值范围**：  - realtimeMetric（实时指标） - highRequest（高请求指标） - slowSql（慢SQL指标） - lockWait（锁等待指标） - diskLimit（磁盘限制指标） - memoryLimit（内存限制指标） - importantMetric（重要指标）  **默认取值**：  不涉及。
	MetricGroup string `json:"metric_group"`

	// **参数解释**：  请求语言类型。  **约束限制**：  不涉及。  **取值范围**：  - en-us - zh-cn  **默认取值**：  en-us。
	XLanguage *string `json:"X-Language,omitempty"`
}

func (o GetInstancesOpsMetricNamesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetInstancesOpsMetricNamesRequest struct{}"
	}

	return strings.Join([]string{"GetInstancesOpsMetricNamesRequest", string(data)}, " ")
}
