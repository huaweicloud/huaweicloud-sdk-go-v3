package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GetInstancesOpsResourceUsageResponse Response Object
type GetInstancesOpsResourceUsageResponse struct {
	Cpu *ResourceUsage `json:"cpu,omitempty"`

	Mem *ResourceUsage `json:"mem,omitempty"`

	Disk *ResourceUsage `json:"disk,omitempty"`

	Io             *ResourceUsage `json:"io,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o GetInstancesOpsResourceUsageResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GetInstancesOpsResourceUsageResponse struct{}"
	}

	return strings.Join([]string{"GetInstancesOpsResourceUsageResponse", string(data)}, " ")
}
