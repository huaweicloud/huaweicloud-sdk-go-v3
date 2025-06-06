package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkloadStatusReq **参数解释**： 资源管理状态请求。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type WorkloadStatusReq struct {
	WorkloadStatus *WorkloadStatus `json:"workload_status,omitempty"`
}

func (o WorkloadStatusReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkloadStatusReq struct{}"
	}

	return strings.Join([]string{"WorkloadStatusReq", string(data)}, " ")
}
