package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// **参数说明**：Kafka配置。
type KafkaConfigRequestDto struct {

	// **参数说明**：Topic前缀，不携带时以user_topics中具体Topic为准，携带时前缀将拼接在user_topics中的topic前方，例如：topic_prefixv2x-v1-tracks，topic_prefixv2x-v1-bsm。
	TopicPrefix *string `json:"topic_prefix,omitempty"`

	// **参数说明**：kafka的主题列表，根据需要转发的消息类型填写。  **取值范围**：  - v2x-v1-tracks：edge上报的车辆轨迹数据  - v2x-v1-bsm：车载T-BOX，mqtt协议接入rsu， websocket协议接入rsu上报的BSM消息数据  - v2x-v1-rsi：mqtt协议接入rsu，edge上报的RSI消息数据  - v2x-v1-rsm： mqtt协议接入rsu，edge上报的RSM消息数据  - v2x-v1-spat：mqtt协议接入rsu， websocket协议接入rsu上报的SPAT消息数据  - v2x-v1-edge-flow：edge上报的车流量统计信息数据
	UserTopics *[]string `json:"user_topics,omitempty"`

	// **参数说明**：Kafka broker列表。
	Brokers []string `json:"brokers"`

	// **参数说明**：kafka用户名。  **取值范围**：只允许字母、数字、下划线（_）、连接符（-）的组合。
	Username *string `json:"username,omitempty"`

	// **参数说明**：kafka密码。
	Password *string `json:"password,omitempty"`
}

func (o KafkaConfigRequestDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KafkaConfigRequestDto struct{}"
	}

	return strings.Join([]string{"KafkaConfigRequestDto", string(data)}, " ")
}
