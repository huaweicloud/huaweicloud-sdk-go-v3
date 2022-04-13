package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
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
