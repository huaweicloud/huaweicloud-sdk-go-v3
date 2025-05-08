package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowKafkaTopicQuotaRequest Request Object
type ShowKafkaTopicQuotaRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 查询类型，默认为topic。
	Type *string `json:"type,omitempty"`

	// 每一页显示的流控数量。
	Limit *string `json:"limit,omitempty"`

	// 页数。
	Offset *string `json:"offset,omitempty"`

	// 查询关键字。
	Keyword *string `json:"keyword,omitempty"`
}

func (o ShowKafkaTopicQuotaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKafkaTopicQuotaRequest struct{}"
	}

	return strings.Join([]string{"ShowKafkaTopicQuotaRequest", string(data)}, " ")
}
