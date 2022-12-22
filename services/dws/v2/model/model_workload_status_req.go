package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 资源管理状态
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
