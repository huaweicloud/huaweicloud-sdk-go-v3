package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateKafkaRebalanceLogTaskResponse Response Object
type CreateKafkaRebalanceLogTaskResponse struct {

	// 任务ID。
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateKafkaRebalanceLogTaskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateKafkaRebalanceLogTaskResponse struct{}"
	}

	return strings.Join([]string{"CreateKafkaRebalanceLogTaskResponse", string(data)}, " ")
}
