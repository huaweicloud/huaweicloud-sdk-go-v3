package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteMqsInstanceTopicResponse Response Object
type BatchDeleteMqsInstanceTopicResponse struct {

	// 待删除的topic列表。
	Topics         *[]BatchDeleteMqsInstanceTopicRespTopics `json:"topics,omitempty"`
	HttpStatusCode int                                      `json:"-"`
}

func (o BatchDeleteMqsInstanceTopicResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteMqsInstanceTopicResponse struct{}"
	}

	return strings.Join([]string{"BatchDeleteMqsInstanceTopicResponse", string(data)}, " ")
}
