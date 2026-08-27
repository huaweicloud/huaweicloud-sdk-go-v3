package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkloadInfo 资源可用容量，不考虑资源已分配量，即资源总容量减去故障资源和热备节点的资源
type WorkloadInfo struct {
	Allocated *WorkloadStatistics `json:"allocated,omitempty"`

	Queue *WorkloadStatistics `json:"queue,omitempty"`

	// UTC时间，格式yyyy-MM-dd'T'HH:mm:ss'Z'。
	Timestamp *string `json:"timestamp,omitempty"`

	// 统计间隔，1s表示1秒，1m表示一分钟，1h为一小时。
	Window *string `json:"window,omitempty"`
}

func (o WorkloadInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkloadInfo struct{}"
	}

	return strings.Join([]string{"WorkloadInfo", string(data)}, " ")
}
