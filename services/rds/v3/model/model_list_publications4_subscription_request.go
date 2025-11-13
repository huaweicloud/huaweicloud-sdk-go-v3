package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPublications4SubscriptionRequest Request Object
type ListPublications4SubscriptionRequest struct {

	// 语言。默认en-us。
	XLanguage *string `json:"X-Language,omitempty"`

	// 实例ID
	InstanceId string `json:"instance_id"`

	Body *ListPublications4SubscriptionRequestBody `json:"body,omitempty"`
}

func (o ListPublications4SubscriptionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPublications4SubscriptionRequest struct{}"
	}

	return strings.Join([]string{"ListPublications4SubscriptionRequest", string(data)}, " ")
}
