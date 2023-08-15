package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkloadQueueReq 资源池
type WorkloadQueueReq struct {
	WorkloadQueue *WorkloadQueue `json:"workload_queue"`
}

func (o WorkloadQueueReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkloadQueueReq struct{}"
	}

	return strings.Join([]string{"WorkloadQueueReq", string(data)}, " ")
}
