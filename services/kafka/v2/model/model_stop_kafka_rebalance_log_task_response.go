package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopKafkaRebalanceLogTaskResponse Response Object
type StopKafkaRebalanceLogTaskResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StopKafkaRebalanceLogTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopKafkaRebalanceLogTaskResponse struct{}"
	}

	return strings.Join([]string{"StopKafkaRebalanceLogTaskResponse", string(data)}, " ")
}
