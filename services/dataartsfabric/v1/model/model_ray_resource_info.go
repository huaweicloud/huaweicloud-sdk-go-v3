package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RayResourceInfo Ray资源配置
type RayResourceInfo struct {
	HeadGroupSpec *HeadGroupSpec `json:"head_group_spec,omitempty"`

	// Ray worker group配置
	WorkerGroupSpec *[]WorkerGroupSpec `json:"worker_group_spec,omitempty"`
}

func (o RayResourceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RayResourceInfo struct{}"
	}

	return strings.Join([]string{"RayResourceInfo", string(data)}, " ")
}
