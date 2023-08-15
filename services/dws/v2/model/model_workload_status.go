package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkloadStatus 资源管理信息
type WorkloadStatus struct {

	// 开关。
	WorkloadSwitch string `json:"workload_switch"`

	// 最大并发数。
	MaxConcurrencyNum *string `json:"max_concurrency_num,omitempty"`
}

func (o WorkloadStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkloadStatus struct{}"
	}

	return strings.Join([]string{"WorkloadStatus", string(data)}, " ")
}
