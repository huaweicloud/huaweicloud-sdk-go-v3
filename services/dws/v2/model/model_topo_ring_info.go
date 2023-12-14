package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TopoRingInfo 集群拓扑实例环信息
type TopoRingInfo struct {

	// 集群实例列表信息
	InstanceInfoLists *[]TopoInstanceInfo `json:"instance_info_lists,omitempty"`
}

func (o TopoRingInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TopoRingInfo struct{}"
	}

	return strings.Join([]string{"TopoRingInfo", string(data)}, " ")
}
