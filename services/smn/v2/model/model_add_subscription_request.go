package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddSubscriptionRequest Request Object
type AddSubscriptionRequest struct {

	// Topic的唯一的资源标识，可通过[查询主题列表](smn_api_51004.xml)获取该标识。
	TopicUrn string `json:"topic_urn"`

	Body *AddSubscriptionRequestBody `json:"body,omitempty"`
}

func (o AddSubscriptionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddSubscriptionRequest struct{}"
	}

	return strings.Join([]string{"AddSubscriptionRequest", string(data)}, " ")
}
