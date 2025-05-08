package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateKafkaTopicQuotaResponse Response Object
type CreateKafkaTopicQuotaResponse struct {

	// 流控配置的任务ID
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateKafkaTopicQuotaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateKafkaTopicQuotaResponse struct{}"
	}

	return strings.Join([]string{"CreateKafkaTopicQuotaResponse", string(data)}, " ")
}
