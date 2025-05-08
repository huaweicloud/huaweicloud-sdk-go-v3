package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteKafkaTopicQuotaResponse Response Object
type DeleteKafkaTopicQuotaResponse struct {

	// 流控配置的任务ID
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteKafkaTopicQuotaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteKafkaTopicQuotaResponse struct{}"
	}

	return strings.Join([]string{"DeleteKafkaTopicQuotaResponse", string(data)}, " ")
}
