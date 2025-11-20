package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NotificationSettingSummary 消息通知配置。
type NotificationSettingSummary struct {

	// 消息通知配置的唯一标识符。
	Id string `json:"id"`

	// 消息通知配置的唯一资源标识符。
	Urn string `json:"urn"`

	// 分析器的唯一标识符。
	AnalyzerId string `json:"analyzer_id"`

	// 分析器的名称。
	AnalyzerName string `json:"analyzer_name"`

	AnalyzerType *AnalyzerType `json:"analyzer_type"`

	// 是否开启消息中心通知开关。
	McSwitch bool `json:"mc_switch"`

	// 消息通知配置的SMN主题URN列表。
	SmnTopicUrns []string `json:"smn_topic_urns"`

	// 创建消息通知配置的时间。
	CreatedAt *sdktime.SdkTime `json:"created_at"`

	// 上次更新消息通知配置的时间。
	UpdatedAt *sdktime.SdkTime `json:"updated_at"`
}

func (o NotificationSettingSummary) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NotificationSettingSummary struct{}"
	}

	return strings.Join([]string{"NotificationSettingSummary", string(data)}, " ")
}
