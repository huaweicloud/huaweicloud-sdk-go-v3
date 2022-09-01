package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchDeleteMqsInstanceTopicRespTopics struct {

	// topic名称。
	Id *string `json:"id,omitempty" xml:"id"`

	// 是否删除成功。
	Success *bool `json:"success,omitempty" xml:"success"`
}

func (o BatchDeleteMqsInstanceTopicRespTopics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteMqsInstanceTopicRespTopics struct{}"
	}

	return strings.Join([]string{"BatchDeleteMqsInstanceTopicRespTopics", string(data)}, " ")
}
