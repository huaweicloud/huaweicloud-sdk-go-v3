package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClusterRing 集群主机信息
type ClusterRing struct {

	// 集群主机信息
	RingHosts []RingHost `json:"ring_hosts"`
}

func (o ClusterRing) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterRing struct{}"
	}

	return strings.Join([]string{"ClusterRing", string(data)}, " ")
}
