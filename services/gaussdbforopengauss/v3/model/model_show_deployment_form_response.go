package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDeploymentFormResponse Response Object
type ShowDeploymentFormResponse struct {

	// 初始节点数。
	InitialNodeNum *int32 `json:"initial_node_num,omitempty"`

	// 解决方案模板名称。
	Solution *string `json:"solution,omitempty"`

	// 分片数。
	ShardNum *int32 `json:"shard_num,omitempty"`

	// 副本数。
	ReplicaNum *int32 `json:"replica_num,omitempty"`

	// **参数解释**: 每次扩容的最小节点数。 **取值范围**: 不涉及。
	EachExpandNodes *int32 `json:"each_expand_nodes,omitempty"`

	// **参数解释**: 最大分片数。 **取值范围**: 不涉及。
	MaxShardCount *int32 `json:"max_shard_count,omitempty"`

	// **参数解释**: 每分片节点数。 **取值范围**: 不涉及。
	EachShardNum   *int32 `json:"each_shard_num,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowDeploymentFormResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDeploymentFormResponse struct{}"
	}

	return strings.Join([]string{"ShowDeploymentFormResponse", string(data)}, " ")
}
