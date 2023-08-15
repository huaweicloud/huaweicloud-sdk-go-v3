package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTopicRequest Request Object
type DeleteTopicRequest struct {

	// Topic的唯一的资源标识，可通过[查询主题列表](smn_api_51004.xml)获取该标识。
	TopicUrn string `json:"topic_urn"`
}

func (o DeleteTopicRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTopicRequest struct{}"
	}

	return strings.Join([]string{"DeleteTopicRequest", string(data)}, " ")
}
