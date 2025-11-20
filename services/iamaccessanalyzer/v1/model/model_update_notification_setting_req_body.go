package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateNotificationSettingReqBody struct {

	// 是否开启消息中心通知开关。
	McSwitch bool `json:"mc_switch"`

	// 消息通知配置的SMN主题URN列表。
	SmnTopicUrns []string `json:"smn_topic_urns"`
}

func (o UpdateNotificationSettingReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNotificationSettingReqBody struct{}"
	}

	return strings.Join([]string{"UpdateNotificationSettingReqBody", string(data)}, " ")
}
