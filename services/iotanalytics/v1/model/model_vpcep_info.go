package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VPCEP信息
type VpcepInfo struct {
	// Kafka的Broker ip

	KafkaBrokerIp string `json:"kafka_broker_ip"`
	// Kafka的VPCEP的service id

	VpcepServiceId string `json:"vpcep_service_id"`
	// Kafka的VPCEP的service name

	VpcepServiceName string `json:"vpcep_service_name"`
	// Kafka的VPCEP的client port

	VpcepClientPort int32 `json:"vpcep_client_port"`
}

func (o VpcepInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VpcepInfo struct{}"
	}

	return strings.Join([]string{"VpcepInfo", string(data)}, " ")
}
