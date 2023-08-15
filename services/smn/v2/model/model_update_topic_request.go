package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateTopicRequest Request Object
type UpdateTopicRequest struct {

	// Topic的唯一的资源标识。可以通过[查看主题列表](smn_api_51004.xml)获取该标识。
	TopicUrn string `json:"topic_urn"`

	Body *UpdateTopicRequestBody `json:"body,omitempty"`
}

func (o UpdateTopicRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTopicRequest struct{}"
	}

	return strings.Join([]string{"UpdateTopicRequest", string(data)}, " ")
}
