package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SubscriptionTypeReq **参数解释**: 订阅的方式 **约束限制**: 不涉及 **取值范围**: - smn_topic：SMN主题 - message_center：消息中心 - silent：无需发送到邮箱  **默认取值**: 不涉及
type SubscriptionTypeReq struct {
}

func (o SubscriptionTypeReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubscriptionTypeReq struct{}"
	}

	return strings.Join([]string{"SubscriptionTypeReq", string(data)}, " ")
}
