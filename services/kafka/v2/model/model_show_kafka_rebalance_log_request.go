package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowKafkaRebalanceLogRequest Request Object
type ShowKafkaRebalanceLogRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`
}

func (o ShowKafkaRebalanceLogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKafkaRebalanceLogRequest struct{}"
	}

	return strings.Join([]string{"ShowKafkaRebalanceLogRequest", string(data)}, " ")
}
