package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowClusterRespCluster **参数解释**： 集群基本信息。
type ShowClusterRespCluster struct {

	// **参数解释**： 控制器ID。 **取值范围**： 不涉及。
	Controller *string `json:"controller,omitempty"`

	// **参数解释**： 节点列表。
	Brokers *[]ShowClusterRespClusterBrokers `json:"brokers,omitempty"`

	// **参数解释**： 主题数量。 **取值范围**： 不涉及。
	TopicsCount *int32 `json:"topics_count,omitempty"`

	// **参数解释**： 分区数量。 **取值范围**： 不涉及。
	PartitionsCount *int32 `json:"partitions_count,omitempty"`

	// **参数解释**： 在线分区数量。 **取值范围**： 不涉及。
	OnlinePartitionsCount *int32 `json:"online_partitions_count,omitempty"`

	// **参数解释**： 副本数量。 **取值范围**： 不涉及。
	ReplicasCount *int32 `json:"replicas_count,omitempty"`

	// **参数解释**： ISR（In-Sync Replicas） 副本总数。 **取值范围**： 不涉及。
	IsrReplicasCount *int32 `json:"isr_replicas_count,omitempty"`

	// **参数解释**： 消费组数量。 **取值范围**： 不涉及。
	ConsumersCount *int32 `json:"consumers_count,omitempty"`
}

func (o ShowClusterRespCluster) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClusterRespCluster struct{}"
	}

	return strings.Join([]string{"ShowClusterRespCluster", string(data)}, " ")
}
