package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowNotificationSettingResponse Response Object
type ShowNotificationSettingResponse struct {

	// 消息通知配置的唯一标识符。
	Id *string `json:"id,omitempty"`

	// 消息通知配置的唯一资源标识符。
	Urn *string `json:"urn,omitempty"`

	// 分析器的唯一标识符。
	AnalyzerId *string `json:"analyzer_id,omitempty"`

	// 分析器的名称。
	AnalyzerName *string `json:"analyzer_name,omitempty"`

	AnalyzerType *AnalyzerType `json:"analyzer_type,omitempty"`

	// 是否开启消息中心通知开关。
	McSwitch *bool `json:"mc_switch,omitempty"`

	// 消息通知配置的SMN主题URN列表。
	SmnTopicUrns *[]string `json:"smn_topic_urns,omitempty"`

	// 创建消息通知配置的时间。
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 上次更新消息通知配置的时间。
	UpdatedAt      *sdktime.SdkTime `json:"updated_at,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowNotificationSettingResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNotificationSettingResponse struct{}"
	}

	return strings.Join([]string{"ShowNotificationSettingResponse", string(data)}, " ")
}
