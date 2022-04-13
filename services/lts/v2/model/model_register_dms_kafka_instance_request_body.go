package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RegisterDmsKafkaInstanceRequestBody struct {
	// kafka ID

	InstanceId string `json:"instance_id"`
	// kafka 名称

	KafkaName string `json:"kafka_name"`

	ConnectInfo *RegisterDmsKafkaInstanceRequestBodyConnectInfo `json:"connect_info"`
}

func (o RegisterDmsKafkaInstanceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterDmsKafkaInstanceRequestBody struct{}"
	}

	return strings.Join([]string{"RegisterDmsKafkaInstanceRequestBody", string(data)}, " ")
}
