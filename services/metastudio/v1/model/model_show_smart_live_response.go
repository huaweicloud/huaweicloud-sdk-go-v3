package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowSmartLiveResponse Response Object
type ShowSmartLiveResponse struct {

	// 数字人直播任务ID。
	JobId *string `json:"job_id,omitempty"`

	// 直播间ID
	RoomId *string `json:"room_id,omitempty"`

	// 直播间名称
	RoomName *string `json:"room_name,omitempty"`

	// 数字人直播任务的状态。 * WAITING: 等待 * PROCESSING: 处理中 * SUCCEED: 成功 * FAILED: 失败 * BLOCKED: 封禁
	State *ShowSmartLiveResponseState `json:"state,omitempty"`

	// **参数解释**： 数字人直播时长，单位秒。
	Duration *float32 `json:"duration,omitempty"`

	// 数字人直播任务开始时间。格式遵循：RFC 3339 如“2021-01-10T08:43:17Z”。
	StartTime *string `json:"start_time,omitempty"`

	// 数字人直播任务结束时间。格式遵循：RFC 3339 如“2021-01-10T08:43:17Z”。
	EndTime *string `json:"end_time,omitempty"`

	ErrorInfo *ErrorResponse `json:"error_info,omitempty"`

	// 数字人直播任务创建时间。格式遵循：RFC 3339 如“2021-01-10T08:43:17Z”。
	CreateTime *string `json:"create_time,omitempty"`

	// 数字人直播任务最后更新时间。格式遵循：RFC 3339 如“2021-01-10T08:43:17Z”。
	LastupdateTime *string `json:"lastupdate_time,omitempty"`

	RtcRoomInfo *RtcRoomInfoList `json:"rtc_room_info,omitempty"`

	// 直播事件上报地址。用户将自行获取的直播间事件上报到此地址，用于触发智能互动，自动回复话术。
	LiveEventReportUrl *string `json:"live_event_report_url,omitempty"`

	LiveEventCallbackConfig *LiveEventCallBackConfig `json:"live_event_callback_config,omitempty"`

	RtcCallbackConfig *RtcLiveEventCallBackConfig `json:"rtc_callback_config,omitempty"`

	// **参数解释**： 数字人直播推流时长，单位秒。
	StreamDuration *float32 `json:"stream_duration,omitempty"`

	// 封禁信息
	BlockReason *string `json:"block_reason,omitempty"`

	// 直播间封面图URL
	CoverUrl *string `json:"cover_url,omitempty"`

	CoStreamerConfig *CoStreamerConfig `json:"co_streamer_config,omitempty"`

	LiveJobLog *LiveJobLog `json:"live_job_log,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowSmartLiveResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSmartLiveResponse struct{}"
	}

	return strings.Join([]string{"ShowSmartLiveResponse", string(data)}, " ")
}

type ShowSmartLiveResponseState struct {
	value string
}

type ShowSmartLiveResponseStateEnum struct {
	WAITING    ShowSmartLiveResponseState
	PROCESSING ShowSmartLiveResponseState
	SUCCEED    ShowSmartLiveResponseState
	FAILED     ShowSmartLiveResponseState
	BLOCKED    ShowSmartLiveResponseState
}

func GetShowSmartLiveResponseStateEnum() ShowSmartLiveResponseStateEnum {
	return ShowSmartLiveResponseStateEnum{
		WAITING: ShowSmartLiveResponseState{
			value: "WAITING",
		},
		PROCESSING: ShowSmartLiveResponseState{
			value: "PROCESSING",
		},
		SUCCEED: ShowSmartLiveResponseState{
			value: "SUCCEED",
		},
		FAILED: ShowSmartLiveResponseState{
			value: "FAILED",
		},
		BLOCKED: ShowSmartLiveResponseState{
			value: "BLOCKED",
		},
	}
}

func (c ShowSmartLiveResponseState) Value() string {
	return c.value
}

func (c ShowSmartLiveResponseState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSmartLiveResponseState) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
