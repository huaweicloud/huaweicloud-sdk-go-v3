package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 物联网平台转发数据的通道配置参数。
type ChannelDetail struct {
	HttpForwarding *HttpForwarding `json:"http_forwarding,omitempty" xml:"http_forwarding"`

	DisForwarding *DisForwarding `json:"dis_forwarding,omitempty" xml:"dis_forwarding"`

	ObsForwarding *ObsForwarding `json:"obs_forwarding,omitempty" xml:"obs_forwarding"`

	AmqpForwarding *AmqpForwarding `json:"amqp_forwarding,omitempty" xml:"amqp_forwarding"`

	DmsKafkaForwarding *DmsKafkaForwarding `json:"dms_kafka_forwarding,omitempty" xml:"dms_kafka_forwarding"`
}

func (o ChannelDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChannelDetail struct{}"
	}

	return strings.Join([]string{"ChannelDetail", string(data)}, " ")
}
