package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateNotificationSettingReqBody struct {

	// 分析器的唯一标识符。
	AnalyzerId string `json:"analyzer_id"`

	// 是否开启消息中心通知开关。
	McSwitch bool `json:"mc_switch"`

	// 消息通知配置的SMN主题URN列表。
	SmnTopicUrns []string `json:"smn_topic_urns"`
}

func (o CreateNotificationSettingReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateNotificationSettingReqBody struct{}"
	}

	return strings.Join([]string{"CreateNotificationSettingReqBody", string(data)}, " ")
}
