package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateKafkaRebalanceLogTaskRequest Request Object
type CreateKafkaRebalanceLogTaskRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`
}

func (o CreateKafkaRebalanceLogTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateKafkaRebalanceLogTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateKafkaRebalanceLogTaskRequest", string(data)}, " ")
}
