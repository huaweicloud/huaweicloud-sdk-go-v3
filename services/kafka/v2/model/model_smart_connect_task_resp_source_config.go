package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SmartConnectTaskRespSourceConfig struct {

	// **参数解释**： 当前Kafka实例别名。（仅源端类型为Kafka时会显示） **取值范围**： 不涉及。
	CurrentClusterName *string `json:"current_cluster_name,omitempty"`

	// **参数解释**： 对端Kafka实例别名。（仅源端类型为Kafka时会显示） **取值范围**： 不涉及。
	ClusterName *string `json:"cluster_name,omitempty"`

	// **参数解释**： 对端Kafka用户名。（仅源端类型为Kafka时会显示） **取值范围**： 不涉及。
	UserName *string `json:"user_name,omitempty"`

	// **参数解释**： 对端Kafka认证机制。（仅源端类型为Kafka时会显示） **取值范围**： - PLAIN - SCRAM-SHA-512
	SaslMechanism *string `json:"sasl_mechanism,omitempty"`

	// **参数解释**： 对端Kafka实例ID。（仅源端类型为Kafka时会显示） **取值范围**： 不涉及。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**： 对端Kafka实例地址。（仅源端类型为Kafka时会显示） **取值范围**： 不涉及。
	BootstrapServers *string `json:"bootstrap_servers,omitempty"`

	// **参数解释**： 对端Kafka认证方式。（仅源端类型为Kafka时会显示） **取值范围**： - PLAINTEXT：不开启SSL，明文传输。 - SASL_SSL：采用SASL方式进行认证，数据通过SSL证书进行加密传输，安全性更高。 - SASL_PLAINTEXT：采用SASL方式进行认证，数据通过明文传输，性能更好。
	SecurityProtocol *string `json:"security_protocol,omitempty"`

	// **参数解释**： 同步方向。（仅源端类型为Kafka时会显示） **取值范围**： - pull：拉取。 - push：推送。 - two-way：双向。
	Direction *string `json:"direction,omitempty"`

	// **参数解释**： 是否同步消费进度。（仅源端类型为Kafka时会显示） **取值范围**： - true：同步消费进度 - false：不同步消费进度
	SyncConsumerOffsetsEnabled *bool `json:"sync_consumer_offsets_enabled,omitempty"`

	// **参数解释**： 副本数。（仅源端类型为Kafka时会显示） **取值范围**： 不涉及。
	ReplicationFactor *int32 `json:"replication_factor,omitempty"`

	// **参数解释**： 任务数。（仅源端类型为Kafka时会显示） **取值范围**： 不涉及。
	TaskNum *int32 `json:"task_num,omitempty"`

	// **参数解释**： 是否重命名Topic。（仅源端类型为Kafka时会显示） **取值范围**： - true：重命名Topic。 - false：不进行重命名Topic。
	RenameTopicEnabled *bool `json:"rename_topic_enabled,omitempty"`

	// **参数解释**： 是否添加来源header。（仅源端类型为Kafka时会显示） **取值范围**： - true：添加来源header。 - false：不添加来源header。
	ProvenanceHeaderEnabled *bool `json:"provenance_header_enabled,omitempty"`

	// **参数解释**： 启动偏移量。（仅源端类型为Kafka时会显示） **取值范围**： - latest：获取最新的数据。 - earliest：获取最早的数据。
	ConsumerStrategy *string `json:"consumer_strategy,omitempty"`

	// **参数解释**： 压缩算法。（仅源端类型为Kafka时会显示） **取值范围**： - none - gzip - snappy - lz4 - zstd
	CompressionType *string `json:"compression_type,omitempty"`

	// **参数解释**： Topic映射。（仅源端类型为Kafka时会显示） **取值范围**： 不涉及。
	TopicsMapping *string `json:"topics_mapping,omitempty"`
}

func (o SmartConnectTaskRespSourceConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmartConnectTaskRespSourceConfig struct{}"
	}

	return strings.Join([]string{"SmartConnectTaskRespSourceConfig", string(data)}, " ")
}
