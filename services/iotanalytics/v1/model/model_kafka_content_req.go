package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Kafka数据源请求内容
type KafkaContentReq struct {

	// KAFKA连接方式
	ConnectionType string `json:"connection_type" xml:"connection_type"`

	// Kafka实例ID
	InstanceId *string `json:"instance_id,omitempty" xml:"instance_id"`

	// Kafka的VPCEP信息包括service_id,service_name,client_port
	VpcepInfos *[]VpcepInfo `json:"vpcep_infos,omitempty" xml:"vpcep_infos"`

	// Kafka的broker信息包括broker_ip, broker_port
	BrokerInfos *[]KafkaBrokerInfo `json:"broker_infos,omitempty" xml:"broker_infos"`

	AuthInfo *KafkaAuthInfo `json:"auth_info,omitempty" xml:"auth_info"`
}

func (o KafkaContentReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KafkaContentReq struct{}"
	}

	return strings.Join([]string{"KafkaContentReq", string(data)}, " ")
}
