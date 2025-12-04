package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowKafkaTopicDetailResponse Response Object
type ShowKafkaTopicDetailResponse struct {

	// Topic详情。
	Body           *[]ShowKafkaTopicDetailResponseBody `json:"body,omitempty"`
	HttpStatusCode int                                 `json:"-"`
}

func (o ShowKafkaTopicDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKafkaTopicDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowKafkaTopicDetailResponse", string(data)}, " ")
}
