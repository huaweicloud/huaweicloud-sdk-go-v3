package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkloadQueueReq **参数解释**： 资源池信息。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
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
