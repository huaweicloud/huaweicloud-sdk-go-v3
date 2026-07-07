package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowDeadLockTopologyGraphRespEdges struct {

	// 起始节点唯一标识
	Source string `json:"source"`

	// 终点节点唯一标识
	Target string `json:"target"`

	// 关系类型
	Type string `json:"type"`
}

func (o ShowDeadLockTopologyGraphRespEdges) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDeadLockTopologyGraphRespEdges struct{}"
	}

	return strings.Join([]string{"ShowDeadLockTopologyGraphRespEdges", string(data)}, " ")
}
