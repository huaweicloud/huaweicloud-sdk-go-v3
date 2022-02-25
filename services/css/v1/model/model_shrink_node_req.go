package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShrinkNodeReq struct {
	// 下线节点个数。 没有Master节点的集群，缩容后剩余的数据节点个数(包含冷数据节点和其他类型节点)须大于之前的一半。 有Master节点的集群，缩容后Master节点的总数须为大于等于3的奇数。 跨AZ的集群，缩容后须确保剩余的节点个数大等于AZ个数。

	ReducedNodeNum int32 `json:"reducedNodeNum"`
	// 下线节点类型。（ess、ess-master、ess-client、ess-cold、lgs）

	Type string `json:"type"`
}

func (o ShrinkNodeReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShrinkNodeReq struct{}"
	}

	return strings.Join([]string{"ShrinkNodeReq", string(data)}, " ")
}
