package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateWorkloadQueueReq **参数解释**： 资源池信息。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type UpdateWorkloadQueueReq struct {
	WorkloadQueue *WorkloadResourceQueue `json:"workload_queue"`
}

func (o UpdateWorkloadQueueReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWorkloadQueueReq struct{}"
	}

	return strings.Join([]string{"UpdateWorkloadQueueReq", string(data)}, " ")
}
