package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateMrsKafkaConfigRequestDto **参数说明**：kafka更新配置，支持topicPrefix、userTopics、brokers，需要把该字段新的对应值全量写入。
type UpdateMrsKafkaConfigRequestDto struct {

	// **参数说明**：用户订阅kafka的主题列表。  **取值范围**：  - v2x-v1-tracks：edge上报的车辆轨迹数据  - v2x-v1-bsm：车载T-BOX，mqtt协议接入rsu， websocket协议接入rsu上报的BSM消息数据  - v2x-v1-rsi：mqtt协议接入rsu，edge上报的RSI消息数据  - v2x-v1-rsm： mqtt协议接入rsu，edge上报的RSM消息数据  - v2x-v1-spat：mqtt协议接入rsu， websocket协议接入rsu上报的SPAT消息数据  - v2x-v1-edge-flow：edge上报的车流量统计信息数据
	UserTopics *[]string `json:"user_topics,omitempty"`

	// **参数说明**：用户mrskafka的brokers列表。
	Brokers *[]string `json:"brokers,omitempty"`

	// **参数说明**：mrskafka用户名，若开启安全认证该参数必填。  **取值范围**：只允许字母、数字、下划线（_）、连接符（-）的组合。
	Username *string `json:"username,omitempty"`

	// **参数说明**：是否开启kerberos安全认证的开关。若开启安全认证则需要先上传kerberos安全认证的凭证。
	Authentication *bool `json:"authentication,omitempty"`

	// **参数说明**：若开启安全认证则需要先上传kerberos安全认证的krb5_file的凭证内容。
	Krb5Content *string `json:"krb5_content,omitempty"`

	// **参数说明**：若开启安全认证则需要先上传kerberos安全认证的keytab_file的凭证内容。
	KeytabContent *string `json:"keytab_content,omitempty"`
}

func (o UpdateMrsKafkaConfigRequestDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMrsKafkaConfigRequestDto struct{}"
	}

	return strings.Join([]string{"UpdateMrsKafkaConfigRequestDto", string(data)}, " ")
}
