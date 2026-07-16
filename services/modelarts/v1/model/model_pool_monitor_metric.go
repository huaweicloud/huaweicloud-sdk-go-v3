package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoolMonitorMetric **参数解释**：监控指标描述。
type PoolMonitorMetric struct {

	// **参数解释**：指标维度信息。
	Dimensions *[]PoolMonitorMetricDimensions `json:"dimensions,omitempty"`

	// **参数解释**：指标名称。 **取值范围**：可选值如下： - cpuUsage：CPU使用量。 - memUsedRate：内存利用率。 - gpuUtil：GPU显卡使用量。 - gpuMemUsage：GPU显存使用量。 - npuUtil：NPU显卡使用量。 - npuMemUsage：NPU显存使用量。 - diskCapacity：磁盘容量。 - diskAvailableCapacity：磁盘可用容量。 - diskUsedRate：磁盘利用率。
	MetricName *string `json:"metricName,omitempty"`

	// **参数解释**：指标命名空间。 **取值范围**：可选值如下： -  PAAS.CONTAINER：组件指标、实例指标、进程指标和容器指标的命名空间 - PAAS.NODE： 主机指标、网络指标、磁盘指标和文件系统指标的命名空间 -  PAAS.SLA：SLA指标的命名空间 - PAAS.AGGR：集群指标的命名空间 - CUSTOMMETRICS：默认的自定义指标的命名空间
	Namespace *string `json:"namespace,omitempty"`
}

func (o PoolMonitorMetric) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoolMonitorMetric struct{}"
	}

	return strings.Join([]string{"PoolMonitorMetric", string(data)}, " ")
}
