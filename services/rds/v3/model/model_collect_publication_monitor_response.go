package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CollectPublicationMonitorResponse Response Object
type CollectPublicationMonitorResponse struct {

	// 发布关联的快照代理的运行状态。取值如下:  started:已启动。 succeeded:成功。 in_progress:正在运行。 idle:空闲。 retrying:重试中。 failed:失败。
	Status *string `json:"status,omitempty"`

	// 数据更改的最长延迟时间（以秒为单位）。
	WorstLatency *int32 `json:"worst_latency,omitempty"`

	// 数据更改的最短延迟时间（以秒为单位）。
	BestLatency *int32 `json:"best_latency,omitempty"`

	// 数据更改的平均延迟时间（以秒为单位）。
	AverageLatency *int32 `json:"average_latency,omitempty"`

	// 上一次分发代理运行时间。格式为“yyyy-mm-ddThh:mm:ssZ”。 其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	LastDistSync *string `json:"last_dist_sync,omitempty"`

	// 等待传送到分发数据库的事务数。
	ReplicatedTransactions *int32 `json:"replicated_transactions,omitempty"`

	// 平均每秒传送到分发数据库的事务数。
	ReplicationRateTrans float32 `json:"replication_rate_trans,omitempty"`
	HttpStatusCode       int     `json:"-"`
}

func (o CollectPublicationMonitorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CollectPublicationMonitorResponse struct{}"
	}

	return strings.Join([]string{"CollectPublicationMonitorResponse", string(data)}, " ")
}
