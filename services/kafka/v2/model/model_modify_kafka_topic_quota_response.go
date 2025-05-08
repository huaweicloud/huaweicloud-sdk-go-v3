package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyKafkaTopicQuotaResponse Response Object
type ModifyKafkaTopicQuotaResponse struct {

	// 流控配置的任务ID
	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ModifyKafkaTopicQuotaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyKafkaTopicQuotaResponse struct{}"
	}

	return strings.Join([]string{"ModifyKafkaTopicQuotaResponse", string(data)}, " ")
}
