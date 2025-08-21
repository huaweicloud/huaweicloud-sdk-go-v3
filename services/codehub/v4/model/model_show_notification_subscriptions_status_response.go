package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowNotificationSubscriptionsStatusResponse Response Object
type ShowNotificationSubscriptionsStatusResponse struct {
	InternalMessage *RepoNotificationSubscriptionStateDto `json:"internal_message,omitempty"`

	Email *RepoNotificationSubscriptionStateDto `json:"email,omitempty"`

	Qyweixin *RepoNotificationSubscriptionStateDto `json:"qyweixin,omitempty"`

	Feishu *RepoNotificationSubscriptionStateDto `json:"feishu,omitempty"`

	Dingding       *RepoNotificationSubscriptionStateDto `json:"dingding,omitempty"`
	HttpStatusCode int                                   `json:"-"`
}

func (o ShowNotificationSubscriptionsStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNotificationSubscriptionsStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowNotificationSubscriptionsStatusResponse", string(data)}, " ")
}
