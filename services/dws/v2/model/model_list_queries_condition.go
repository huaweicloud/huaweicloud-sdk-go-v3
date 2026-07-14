package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListQueriesCondition struct {

	// **参数解释**： 字段名称。 **取值范围**： systemQuery：是否隐藏系统查询。 userName：用户名称。 applicationName：应用名称。 dbName：数据库名称。 resourcePool：资源池。 queryStatus：查询状态。 enqueue：排队状态。 lane：通道。 instName：实例名称。 pid：进程ID。 blockTime：阻塞时间。 duration：持续时间。 minCpuTime：最小CPU时间。 maxCpuTime：最大CPU时间。 totalCpuTime：总CPU时间。 cpuSkewPercent：CPU倾斜率。 spillInfo：下盘信息。 minSpillSize：最小下盘大小。 maxSpillSize：最大下盘大小。 averageSpillSize：平均下盘大小。 spillSkewPercent：下盘倾斜率。 queryBand：查询标签。 jobName：作业名称。 jobInst：作业实例。 clientHostname：客户端主机名。 clientPort：客户端端口。 waiting：等待状态。 estimateTotalTime：预估总时间。 estimateLeftTime：预估剩余时间。 controlGroup：控制组。 minPeakMemory：最小峰值内存。 maxPeakMemory：最大峰值内存。 averagePeakMemory：平均峰值内存。 memorySkewPercent：内存倾斜率。 estimateMemory：预估内存。 minDnTime：最小DN时间。 maxDnTime：最大DN时间。 averageDnTime：平均DN时间。 dntimeSkewPercent：DN时间倾斜率。 warning：告警。 averagePeakIops：平均峰值IOPS。 iopsSkewPercent：IOPS倾斜率。 wlmStatus：WLM状态。 wlmAttrib：WLM属性。
	Field string `json:"field"`

	// **参数解释**： 字段值。 **取值范围**： 不涉及。
	Value string `json:"value"`

	// **参数解释**： 比较方式。 **取值范围**： String类型参数：=、!=、like、not like、<> int类型参数：=、!=、>、<、>=、<=、<> boolean类型参数：=、!=
	Operator string `json:"operator"`
}

func (o ListQueriesCondition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQueriesCondition struct{}"
	}

	return strings.Join([]string{"ListQueriesCondition", string(data)}, " ")
}
