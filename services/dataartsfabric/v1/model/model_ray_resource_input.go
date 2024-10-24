package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RayResourceInput 创建ray集群的资源配置
type RayResourceInput struct {
	HeadNodeResource *HeadNodeResourceDemand `json:"head_node_resource"`

	// 工作节点资源配置
	WorkNodeResources []WorkNodeResourceDemand `json:"work_node_resources"`
}

func (o RayResourceInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RayResourceInput struct{}"
	}

	return strings.Join([]string{"RayResourceInput", string(data)}, " ")
}
