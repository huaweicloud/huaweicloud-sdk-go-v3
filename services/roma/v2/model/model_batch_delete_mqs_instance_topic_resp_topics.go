package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteMqsInstanceTopicRespTopics struct {
	// topic名称。

	Id *string `json:"id,omitempty"`
	// 是否删除成功。

	Success *bool `json:"success,omitempty"`
}

func (o BatchDeleteMqsInstanceTopicRespTopics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteMqsInstanceTopicRespTopics struct{}"
	}

	return strings.Join([]string{"BatchDeleteMqsInstanceTopicRespTopics", string(data)}, " ")
}
