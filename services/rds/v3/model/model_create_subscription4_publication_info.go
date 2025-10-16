package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateSubscription4PublicationInfo 给发布创建订阅时需要的参数。
type CreateSubscription4PublicationInfo struct {

	// 订阅服务器来源为云上实例时的订阅实例id。
	SubscriptionInstanceId string `json:"subscription_instance_id"`

	UserInfo *ReplicationUserInfo `json:"user_info,omitempty"`
}

func (o CreateSubscription4PublicationInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSubscription4PublicationInfo struct{}"
	}

	return strings.Join([]string{"CreateSubscription4PublicationInfo", string(data)}, " ")
}
