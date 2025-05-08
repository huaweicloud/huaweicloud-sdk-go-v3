package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopKafkaRebalanceLogTaskRequest Request Object
type StopKafkaRebalanceLogTaskRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`
}

func (o StopKafkaRebalanceLogTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopKafkaRebalanceLogTaskRequest struct{}"
	}

	return strings.Join([]string{"StopKafkaRebalanceLogTaskRequest", string(data)}, " ")
}
