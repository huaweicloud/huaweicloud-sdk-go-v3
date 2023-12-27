package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartSmartLiveResponse Response Object
type StartSmartLiveResponse struct {

	// 直播任务ID。
	JobId *string `json:"job_id,omitempty"`

	RtcRoomInfo *RtcRoomInfoList `json:"rtc_room_info,omitempty"`

	// 直播事件上报地址。用户将自行获取的直播间事件上报到此地址，用于触发智能互动，自动回复话术。
	LiveEventReportUrl *string `json:"live_event_report_url,omitempty"`

	LiveEventCallbackConfig *LiveEventCallBackConfig `json:"live_event_callback_config,omitempty"`

	// 开播风险告警列表。
	LiveWarningInfo *[]LiveWarningItem `json:"live_warning_info,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StartSmartLiveResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartSmartLiveResponse struct{}"
	}

	return strings.Join([]string{"StartSmartLiveResponse", string(data)}, " ")
}
