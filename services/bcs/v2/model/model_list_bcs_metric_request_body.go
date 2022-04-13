package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BCS服务监控数据查询请求结构
type ListBcsMetricRequestBody struct {
	// 指标列表 取值范围 cpuUsage：CPU使用率 diskUsedRate：磁盘使用率 memUsedRate：物理内存使用率 sendBytesRate：上行Bps recvBytesRate：下行Bps cpuCoreLimit：CPU内核总量 cpuCoreUsed：CPU内核占用 totalMem：物理内存总量 freeMem：可用物理内存 diskCapacity：磁盘空间容量 diskAvailableCapacity：可用磁盘空间 默认值：前5项

	MetricNames *[]string `json:"metric_names,omitempty"`
}

func (o ListBcsMetricRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBcsMetricRequestBody struct{}"
	}

	return strings.Join([]string{"ListBcsMetricRequestBody", string(data)}, " ")
}
