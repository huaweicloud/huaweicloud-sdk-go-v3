package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeleteMqsInstanceTopicRespTopics struct {

	// Topic名称。
	Id *string `json:"id,omitempty" xml:"id"`

	// 是否删除成功。
	Success *bool `json:"success,omitempty" xml:"success"`
}

func (o DeleteMqsInstanceTopicRespTopics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteMqsInstanceTopicRespTopics struct{}"
	}

	return strings.Join([]string{"DeleteMqsInstanceTopicRespTopics", string(data)}, " ")
}
