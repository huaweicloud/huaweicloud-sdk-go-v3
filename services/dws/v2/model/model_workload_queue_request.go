package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkloadQueueRequest 更新资源池
type WorkloadQueueRequest struct {
	WorkloadQueue *WorkloadQueueInfo `json:"workload_queue"`
}

func (o WorkloadQueueRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkloadQueueRequest struct{}"
	}

	return strings.Join([]string{"WorkloadQueueRequest", string(data)}, " ")
}
