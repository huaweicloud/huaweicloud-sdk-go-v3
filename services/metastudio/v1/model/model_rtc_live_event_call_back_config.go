package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RtcLiveEventCallBackConfig RTC回调事件配置。兼容处理，未携带配置则默认订阅LIVE_PROGRESS
type RtcLiveEventCallBackConfig struct {

	// RTC回调的直播事件类型列表。  当前仅支持如下取值： * LIVE_PROGRESS：直播剧本进度通知。  * REPLY_COMMAND_FINISH：回复播放完成通知。  回调事件结构体定义： * message_type：消息类型。 * data：消息描述。   - LIVE_PROGRESS事件回调定义如下：     ```json     {         \"message_type\": \"live_progress_notify\",         \"data\": {             \"script_name\": \"场景一\",             \"shoot_script_sequence_no\": 2,             \"shoot_script_title\": \"引导语\",             \"offset\": \"247\",             \"reply_id\": \"e87104f76d7546ce8a46ac6b04c49c3c\"         }     }     ```   - REPLY_COMMAND_FINISH回调定义如下：     ```json     {       \"message_type\": \"reply_command_finish_notify\",       \"data\":\"{         \"reply_id\":\"e87104f76d7546ce8a46ac6b04c49c3c\"       }\"     }     ```
	RtcCallbackEventType *[]string `json:"rtc_callback_event_type,omitempty"`
}

func (o RtcLiveEventCallBackConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RtcLiveEventCallBackConfig struct{}"
	}

	return strings.Join([]string{"RtcLiveEventCallBackConfig", string(data)}, " ")
}
