package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTopicAttributesRequest Request Object
type ListTopicAttributesRequest struct {

	// Topic的唯一的资源标识，可通过[查询主题列表](smn_api_51004.xml)获取该标识。
	TopicUrn string `json:"topic_urn"`

	// 主题策略名称。  只支持特定的策略名称，请参见[Topic属性表](smn_api_a1000.xml)。
	Name *string `json:"name,omitempty"`
}

func (o ListTopicAttributesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTopicAttributesRequest struct{}"
	}

	return strings.Join([]string{"ListTopicAttributesRequest", string(data)}, " ")
}
