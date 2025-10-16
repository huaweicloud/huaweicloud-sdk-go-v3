package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSubscriptionRequestBody 创建发布详情。
type CreateSubscriptionRequestBody struct {

	// 订阅详情。一次最多创建十个订阅。
	Subscriptions []CreateSubscriptionInfo `json:"subscriptions"`

	// 给发布创建订阅时的发布id。给发布创建订阅时必传，不传时则为创建本地订阅。
	CurrentPublicationId *string `json:"current_publication_id,omitempty"`
}

func (o CreateSubscriptionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSubscriptionRequestBody struct{}"
	}

	return strings.Join([]string{"CreateSubscriptionRequestBody", string(data)}, " ")
}
