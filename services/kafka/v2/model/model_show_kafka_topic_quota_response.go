package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowKafkaTopicQuotaResponse Response Object
type ShowKafkaTopicQuotaResponse struct {

	// Topic流控配置
	Quotas *[]KafkaTopicQuota `json:"quotas,omitempty"`

	// Topic流控数量
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowKafkaTopicQuotaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKafkaTopicQuotaResponse struct{}"
	}

	return strings.Join([]string{"ShowKafkaTopicQuotaResponse", string(data)}, " ")
}
