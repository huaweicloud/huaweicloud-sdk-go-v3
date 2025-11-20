package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LiveNotifyConfig 直播通知配置。通过短信，邮件，站内信通知用户直播中断。用户可在华为云消息中心配置接受者信息
type LiveNotifyConfig struct {

	// **参数解释**： 启用通知事件列表。 **约束限制**： 不涉及 **默认取值**： 无。
	NotifyEvents *[]NotifyEventEnum `json:"notify_events,omitempty"`
}

func (o LiveNotifyConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LiveNotifyConfig struct{}"
	}

	return strings.Join([]string{"LiveNotifyConfig", string(data)}, " ")
}
