package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Subscription4PublicationInfo 发布下的订阅信息。
type Subscription4PublicationInfo struct {

	// 订阅服务器名称。
	SubscriptionInstanceName *string `json:"subscription_instance_name,omitempty"`

	// 订阅服务器ip。
	SubscriptionInstanceIp *string `json:"subscription_instance_ip,omitempty"`

	// 订阅实例id。
	SubscriptionInstanceId *string `json:"subscription_instance_id,omitempty"`
}

func (o Subscription4PublicationInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Subscription4PublicationInfo struct{}"
	}

	return strings.Join([]string{"Subscription4PublicationInfo", string(data)}, " ")
}
