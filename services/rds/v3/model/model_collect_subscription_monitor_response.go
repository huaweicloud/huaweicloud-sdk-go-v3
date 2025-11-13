package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CollectSubscriptionMonitorResponse Response Object
type CollectSubscriptionMonitorResponse struct {

	// 订阅关联的快照代理的运行状态。取值如下:  started:已启动。 succeeded:成功。 in_progress:正在运行。 idle:空闲。 retrying:重试中。 failed:失败。
	Status *string `json:"status,omitempty"`

	// 数据更改的最长延迟时间（以秒为单位）。
	Latency *int32 `json:"latency,omitempty"`

	// 上一次分发代理运行时间。格式为“yyyy-mm-ddThh:mm:ssZ”。 其中，T指某个时间的开始；Z指时区偏移量，例如北京时间偏移显示为+0800。
	LastDistSync *string `json:"last_dist_sync,omitempty"`

	// 代理未运行的时长（以小时为单位）。
	AgentNotRunning *int32 `json:"agent_not_running,omitempty"`

	// 订阅未执行的命令数。
	PendingCmdCount *int32 `json:"pending_cmd_count,omitempty"`

	// 预计执行完未执行的命令数所需时间（以秒为单位）。
	EstimatedProcessTime *int32 `json:"estimated_process_time,omitempty"`
	HttpStatusCode       int    `json:"-"`
}

func (o CollectSubscriptionMonitorResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CollectSubscriptionMonitorResponse struct{}"
	}

	return strings.Join([]string{"CollectSubscriptionMonitorResponse", string(data)}, " ")
}
