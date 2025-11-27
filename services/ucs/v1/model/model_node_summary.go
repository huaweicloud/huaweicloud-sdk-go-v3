package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NodeSummary struct {

	// 集群中所有节点的个数。
	TotalNum *int32 `json:"totalNum,omitempty"`

	// 集群中已就绪节点的数量。
	ReadyNum *int32 `json:"readyNum,omitempty"`
}

func (o NodeSummary) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeSummary struct{}"
	}

	return strings.Join([]string{"NodeSummary", string(data)}, " ")
}
