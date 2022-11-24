package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VPCEP信息
type VpcepInfoRsp struct {

	// Kafka的VPCEP的service id
	VpcepServiceId string `json:"vpcep_service_id"`

	// Kafka的VPCEP的service name
	VpcepServiceName string `json:"vpcep_service_name"`

	// Kafka的VPCEP的client ip
	VpcepClientIp *string `json:"vpcep_client_ip,omitempty"`

	// Kafka的VPCEP的client port
	VpcepClientPort int32 `json:"vpcep_client_port"`

	// Kafka的Broker ip
	KafkaBrokerIp string `json:"kafka_broker_ip"`
}

func (o VpcepInfoRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VpcepInfoRsp struct{}"
	}

	return strings.Join([]string{"VpcepInfoRsp", string(data)}, " ")
}
