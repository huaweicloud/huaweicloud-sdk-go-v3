package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubscriptionTypeRes **参数解释**: 订阅的方式 **取值范围**: - smn_topic：SMN主题 - message_center：消息中心 - silent：无需发送到邮箱
type SubscriptionTypeRes struct {
}

func (o SubscriptionTypeRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubscriptionTypeRes struct{}"
	}

	return strings.Join([]string{"SubscriptionTypeRes", string(data)}, " ")
}
