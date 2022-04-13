package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 集群的节点信息
type ClusterInstance struct {
	// 节点的虚拟机ID。

	Id *string `json:"id,omitempty"`
	// 节点的虚拟机名称。

	Name *string `json:"name,omitempty"`
	// 节点类型，只支持一种类型“cdm”。

	Type *string `json:"type,omitempty"`
	// 分片ID

	ShardId *string `json:"shard_id,omitempty"`
}

func (o ClusterInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterInstance struct{}"
	}

	return strings.Join([]string{"ClusterInstance", string(data)}, " ")
}
