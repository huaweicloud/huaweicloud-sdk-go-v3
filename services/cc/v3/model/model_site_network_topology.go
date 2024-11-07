package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SiteNetworkTopology 拓扑结构。
type SiteNetworkTopology struct {
	Topology *SiteNetworkTopologyEnum `json:"topology"`
}

func (o SiteNetworkTopology) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SiteNetworkTopology struct{}"
	}

	return strings.Join([]string{"SiteNetworkTopology", string(data)}, " ")
}
