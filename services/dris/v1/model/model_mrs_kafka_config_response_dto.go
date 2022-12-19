package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// **参数说明**：mrsKafka配置信息。
type MrsKafkaConfigResponseDto struct {

	// **参数说明**：每一套Kafka配置的唯一ID。  **取值范围**：只允许字母、数字、下划线（_）、连接符（-）的组合。
	KafkaConfigId *string `json:"kafka_config_id,omitempty"`

	// **参数说明**：kafka的主题列表。  **取值范围**：  - v2x-v1-tracks：edge上报的车辆轨迹数据  - v2x-v1-bsm：车载T-BOX，mqtt协议接入rsu， websocket协议接入rsu上报的BSM消息数据  - v2x-v1-rsi：mqtt协议接入rsu，edge上报的RSI消息数据  - v2x-v1-rsm： mqtt协议接入rsu，edge上报的RSM消息数据  - v2x-v1-spat：mqtt协议接入rsu， websocket协议接入rsu上报的SPAT消息数据  - v2x-v1-edge-flow：edge上报的车流量统计信息数据
	KafkaTopics *[]string `json:"kafka_topics,omitempty"`

	// **参数说明**：Kafka broker列表。
	Brokers *[]string `json:"brokers,omitempty"`

	// **参数说明**：mrskafka用户名，若开启安全认证该参数必填。  **取值范围**：只允许字母、数字、下划线（_）、连接符（-）的组合。
	Username *string `json:"username,omitempty"`

	// **参数说明**：一套kafka的连接状态。  **取值范围**：  - OFFLINE：离线  - ONLINE：在线
	Status *MrsKafkaConfigResponseDtoStatus `json:"status,omitempty"`

	// **参数说明**：是否开启kerberos安全认证的开关。若开启安全认证则需要先上传kerberos安全认证的凭证。
	Authentication *bool `json:"authentication,omitempty"`

	// **参数说明**：创建时间。 格式为yyyy-MM-dd'T'HH:mm:ss'Z' 例如：2015-12-12T12:12:12Z
	CreatedTime *sdktime.SdkTime `json:"created_time,omitempty"`

	// **参数说明**：修改时间。 格式为yyyy-MM-dd'T'HH:mm:ss'Z' 例如：2015-12-12T12:12:12Z
	LastModifiedTime *sdktime.SdkTime `json:"last_modified_time,omitempty"`
}

func (o MrsKafkaConfigResponseDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MrsKafkaConfigResponseDto struct{}"
	}

	return strings.Join([]string{"MrsKafkaConfigResponseDto", string(data)}, " ")
}

type MrsKafkaConfigResponseDtoStatus struct {
	value string
}

type MrsKafkaConfigResponseDtoStatusEnum struct {
	OFFLINE MrsKafkaConfigResponseDtoStatus
	ONLINE  MrsKafkaConfigResponseDtoStatus
}

func GetMrsKafkaConfigResponseDtoStatusEnum() MrsKafkaConfigResponseDtoStatusEnum {
	return MrsKafkaConfigResponseDtoStatusEnum{
		OFFLINE: MrsKafkaConfigResponseDtoStatus{
			value: "OFFLINE",
		},
		ONLINE: MrsKafkaConfigResponseDtoStatus{
			value: "ONLINE",
		},
	}
}

func (c MrsKafkaConfigResponseDtoStatus) Value() string {
	return c.value
}

func (c MrsKafkaConfigResponseDtoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MrsKafkaConfigResponseDtoStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
