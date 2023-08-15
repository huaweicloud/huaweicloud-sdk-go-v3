package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KafkaContentRsp Kafka数据源请求内容
type KafkaContentRsp struct {

	// KAFKA连接类型
	ConnectionType *string `json:"connection_type,omitempty"`

	// Kafka实例ID
	InstanceId *string `json:"instance_id,omitempty"`

	// Kafka的VPCEP信息包括service_id,service_name,client_port
	VpcepInfos *[]VpcepInfoRsp `json:"vpcep_infos,omitempty"`

	// Kafka的broker信息包括broker_ip, broker_port
	BrokerInfos *[]KafkaBrokerInfo `json:"broker_infos,omitempty"`

	AuthInfo *KafkaAuthInfo `json:"auth_info,omitempty"`
}

func (o KafkaContentRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KafkaContentRsp struct{}"
	}

	return strings.Join([]string{"KafkaContentRsp", string(data)}, " ")
}
