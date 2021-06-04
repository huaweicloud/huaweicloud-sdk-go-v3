package model

import (
	"encoding/json"

	"strings"
)

// 物联网平台转发数据的通道配置参数。
type ChannelDetail struct {
	HttpForwarding *HttpForwarding `json:"http_forwarding,omitempty"`

	DisForwarding *DisForwarding `json:"dis_forwarding,omitempty"`

	ObsForwarding *ObsForwarding `json:"obs_forwarding,omitempty"`

	AmqpForwarding *AmqpForwarding `json:"amqp_forwarding,omitempty"`

	DmsKafkaForwarding *DmsKafkaForwarding `json:"dms_kafka_forwarding,omitempty"`

	RomaForwarding *RomaForwarding `json:"roma_forwarding,omitempty"`

	IotaForwarding *IoTaForwarding `json:"iota_forwarding,omitempty"`

	MqsForwarding *MqsForwarding `json:"mqs_forwarding,omitempty"`

	MysqlForwarding *MysqlForwarding `json:"mysql_forwarding,omitempty"`

	MqttForwarding *MqttForwarding `json:"mqtt_forwarding,omitempty"`

	LtsForwarding *LtsForwarding `json:"lts_forwarding,omitempty"`

	InfluxdbForwarding *InfluxDbForwarding `json:"influxdb_forwarding,omitempty"`

	FunctiongraphForwarding *FunctionGraphForwarding `json:"functiongraph_forwarding,omitempty"`

	MrsKafkaForwarding *MrsKafkaForwarding `json:"mrs_kafka_forwarding,omitempty"`

	PulsarForwarding *PulsarForwarding `json:"pulsar_forwarding,omitempty"`
}

func (o ChannelDetail) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ChannelDetail struct{}"
	}

	return strings.Join([]string{"ChannelDetail", string(data)}, " ")
}
