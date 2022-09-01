package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type NodePoolSpecUpdate struct {
	NodeTemplate *NodeSpecUpdate `json:"nodeTemplate" xml:"nodeTemplate"`

	// 节点池初始化节点个数。查询时为节点池目标节点数量。
	InitialNodeCount int32 `json:"initialNodeCount" xml:"initialNodeCount"`

	Autoscaling *NodePoolNodeAutoscaling `json:"autoscaling" xml:"autoscaling"`
}

func (o NodePoolSpecUpdate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodePoolSpecUpdate struct{}"
	}

	return strings.Join([]string{"NodePoolSpecUpdate", string(data)}, " ")
}
