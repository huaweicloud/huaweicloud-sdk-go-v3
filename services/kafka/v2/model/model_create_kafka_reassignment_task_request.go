package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateKafkaReassignmentTaskRequest Request Object
type CreateKafkaReassignmentTaskRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *PartitionReassignRequest `json:"body,omitempty"`
}

func (o CreateKafkaReassignmentTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateKafkaReassignmentTaskRequest struct{}"
	}

	return strings.Join([]string{"CreateKafkaReassignmentTaskRequest", string(data)}, " ")
}
