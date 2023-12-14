package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LogicalClusterRingInfo 集群实例环信息
type LogicalClusterRingInfo struct {

	// 集群主机环信息
	RingHosts *[]RingHost `json:"ring_hosts,omitempty"`
}

func (o LogicalClusterRingInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogicalClusterRingInfo struct{}"
	}

	return strings.Join([]string{"LogicalClusterRingInfo", string(data)}, " ")
}
