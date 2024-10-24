package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RayResource Ray资源配置
type RayResource struct {
	HeadGroupSpec *HeadGroupSpec `json:"head_group_spec,omitempty"`

	// Ray worker group配置
	WorkerGroupSpec *[]WorkerGroupSpec `json:"worker_group_spec,omitempty"`
}

func (o RayResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RayResource struct{}"
	}

	return strings.Join([]string{"RayResource", string(data)}, " ")
}
