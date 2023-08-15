package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KafkaContentReq Kafka数据源请求内容
type KafkaContentReq struct {

	// KAFKA连接方式
	ConnectionType string `json:"connection_type"`

	// Kafka实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// Kafka的VPCEP信息包括service_id,service_name,client_port
	VpcepInfos *[]VpcepInfo `json:"vpcep_infos,omitempty"`

	// Kafka的broker信息包括broker_ip, broker_port
	BrokerInfos *[]KafkaBrokerInfo `json:"broker_infos,omitempty"`

	AuthInfo *KafkaAuthInfo `json:"auth_info,omitempty"`
}

func (o KafkaContentReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KafkaContentReq struct{}"
	}

	return strings.Join([]string{"KafkaContentReq", string(data)}, " ")
}
