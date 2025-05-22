package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkloadQueueRequest **参数解释**： 更新资源池请求信息。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
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
