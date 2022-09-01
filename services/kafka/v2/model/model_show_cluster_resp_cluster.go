package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 集群基本信息。
type ShowClusterRespCluster struct {

	// 控制器ID。
	Controller *string `json:"controller,omitempty" xml:"controller"`

	// 节点列表。
	Brokers *[]ShowClusterRespClusterBrokers `json:"brokers,omitempty" xml:"brokers"`

	// 主题数量。
	TopicsCount *int32 `json:"topics_count,omitempty" xml:"topics_count"`

	// 分区数量。
	PartitionsCount *int32 `json:"partitions_count,omitempty" xml:"partitions_count"`

	// 在线分区数量。
	OnlinePartitionsCount *int32 `json:"online_partitions_count,omitempty" xml:"online_partitions_count"`

	// 副本数量。
	ReplicasCount *int32 `json:"replicas_count,omitempty" xml:"replicas_count"`

	// ISR（In-Sync Replicas） 副本总数。
	IsrReplicasCount *int32 `json:"isr_replicas_count,omitempty" xml:"isr_replicas_count"`

	// 消费组数量。
	ConsumersCount *int32 `json:"consumers_count,omitempty" xml:"consumers_count"`
}

func (o ShowClusterRespCluster) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClusterRespCluster struct{}"
	}

	return strings.Join([]string{"ShowClusterRespCluster", string(data)}, " ")
}
