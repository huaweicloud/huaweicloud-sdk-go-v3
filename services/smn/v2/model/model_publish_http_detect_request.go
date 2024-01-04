package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PublishHttpDetectRequest Request Object
type PublishHttpDetectRequest struct {

	// Topic的唯一的资源标识，可通过[查询主题列表](smn_api_51004.xml)获取该标识。
	TopicUrn string `json:"topic_urn"`

	Body *HttpDetectRequestBody `json:"body,omitempty"`
}

func (o PublishHttpDetectRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PublishHttpDetectRequest struct{}"
	}

	return strings.Join([]string{"PublishHttpDetectRequest", string(data)}, " ")
}
