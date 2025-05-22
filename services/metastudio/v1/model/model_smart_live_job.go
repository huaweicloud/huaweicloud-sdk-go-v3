package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SmartLiveJob 数字人直播任务信息。
type SmartLiveJob struct {

	// 数字人直播任务ID。
	JobId *string `json:"job_id,omitempty"`

	// 直播间ID
	RoomId *string `json:"room_id,omitempty"`

	// 直播间名称
	RoomName *string `json:"room_name,omitempty"`

	// 数字人直播任务的状态。 * WAITING: 等待 * PROCESSING: 处理中 * SUCCEED: 成功 * FAILED: 失败 * BLOCKED: 封禁
	State *SmartLiveJobState `json:"state,omitempty"`

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

	RelationLivePlatformInfo *PlatformLiveDetailInfo `json:"relation_live_platform_info,omitempty"`

	// 使用的资源类型。 * PERIOD：包周期资源 * ONDEMAND：按需资源 * ONE_TIME：一次性资源 * UNKNOW：未知资源类型。
	UsedResourceType *SmartLiveJobUsedResourceType `json:"used_resource_type,omitempty"`
}

func (o SmartLiveJob) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SmartLiveJob struct{}"
	}

	return strings.Join([]string{"SmartLiveJob", string(data)}, " ")
}

type SmartLiveJobState struct {
	value string
}

type SmartLiveJobStateEnum struct {
	WAITING    SmartLiveJobState
	PROCESSING SmartLiveJobState
	SUCCEED    SmartLiveJobState
	FAILED     SmartLiveJobState
	BLOCKED    SmartLiveJobState
}

func GetSmartLiveJobStateEnum() SmartLiveJobStateEnum {
	return SmartLiveJobStateEnum{
		WAITING: SmartLiveJobState{
			value: "WAITING",
		},
		PROCESSING: SmartLiveJobState{
			value: "PROCESSING",
		},
		SUCCEED: SmartLiveJobState{
			value: "SUCCEED",
		},
		FAILED: SmartLiveJobState{
			value: "FAILED",
		},
		BLOCKED: SmartLiveJobState{
			value: "BLOCKED",
		},
	}
}

func (c SmartLiveJobState) Value() string {
	return c.value
}

func (c SmartLiveJobState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SmartLiveJobState) UnmarshalJSON(b []byte) error {
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

type SmartLiveJobUsedResourceType struct {
	value string
}

type SmartLiveJobUsedResourceTypeEnum struct {
	PERIOD   SmartLiveJobUsedResourceType
	ONDEMAND SmartLiveJobUsedResourceType
	ONE_TIME SmartLiveJobUsedResourceType
	UNKNOW   SmartLiveJobUsedResourceType
}

func GetSmartLiveJobUsedResourceTypeEnum() SmartLiveJobUsedResourceTypeEnum {
	return SmartLiveJobUsedResourceTypeEnum{
		PERIOD: SmartLiveJobUsedResourceType{
			value: "PERIOD",
		},
		ONDEMAND: SmartLiveJobUsedResourceType{
			value: "ONDEMAND",
		},
		ONE_TIME: SmartLiveJobUsedResourceType{
			value: "ONE_TIME",
		},
		UNKNOW: SmartLiveJobUsedResourceType{
			value: "UNKNOW",
		},
	}
}

func (c SmartLiveJobUsedResourceType) Value() string {
	return c.value
}

func (c SmartLiveJobUsedResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SmartLiveJobUsedResourceType) UnmarshalJSON(b []byte) error {
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
