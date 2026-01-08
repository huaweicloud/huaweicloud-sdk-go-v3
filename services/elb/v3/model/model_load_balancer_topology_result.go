package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LoadBalancerTopologyResult
type LoadBalancerTopologyResult struct {
	Nodes *TopologyNodes `json:"nodes"`

	// **参数解释**：拓扑边的信息
	Edges []TopologyEdge `json:"edges"`

	Labels *TopologyLabels `json:"labels"`
}

func (o LoadBalancerTopologyResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoadBalancerTopologyResult struct{}"
	}

	return strings.Join([]string{"LoadBalancerTopologyResult", string(data)}, " ")
}
