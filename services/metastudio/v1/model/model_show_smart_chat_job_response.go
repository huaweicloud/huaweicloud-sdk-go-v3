package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowSmartChatJobResponse Response Object
type ShowSmartChatJobResponse struct {

	// 数字人智能交互对话任务ID。
	JobId *string `json:"job_id,omitempty"`

	// 数字人智能交互对话任务的状态。 * WAITING: 等待 * PROCESSING: 处理中 * SUCCEED: 成功 * FAILED: 失败 * CANCELED: 取消 * HEARTBEAT: 心跳
	State *ShowSmartChatJobResponseState `json:"state,omitempty"`

	// 数字人智能交互对话时长，单位秒。
	Duration *float32 `json:"duration,omitempty"`

	// 数字人智能交互对话任务开始时间。格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	StartTime *string `json:"start_time,omitempty"`

	// 数字人智能交互对话任务结束时间。格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	EndTime *string `json:"end_time,omitempty"`

	ErrorInfo *ErrorResponse `json:"error_info,omitempty"`

	// 数字人智能交互对话任务创建时间。格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	CreateTime *string `json:"create_time,omitempty"`

	// 数字人智能交互对话任务最后更新时间。格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	LastupdateTime *string `json:"lastupdate_time,omitempty"`

	RtcRoomInfo *RtcRoomInfoList `json:"rtc_room_info,omitempty"`

	ChatSubtitleConfig *SmartChatSubtitleConfig `json:"chat_subtitle_config,omitempty"`

	VideoConfig *SmartChatVideoConfig `json:"video_config,omitempty"`

	// 数字人智能交互对话的状态。 0: 等待建链 1: 等待关闭链路 2: 建链成功 3: 进入休眠 4: 等待休眠
	ChatState *int32 `json:"chat_state,omitempty"`

	Language *LanguageEnum `json:"language,omitempty"`

	// 智能交互对话端配置。 * COMPUTER: 电脑端 * MOBILE: 手机端 * HUB: 大屏
	ChatVideoType *ShowSmartChatJobResponseChatVideoType `json:"chat_video_type,omitempty"`

	// 智能交互接入地址。
	ChatAccessAddress *string `json:"chat_access_address,omitempty"`

	// 是否透明背景
	IsTransparent *bool `json:"is_transparent,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowSmartChatJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSmartChatJobResponse struct{}"
	}

	return strings.Join([]string{"ShowSmartChatJobResponse", string(data)}, " ")
}

type ShowSmartChatJobResponseState struct {
	value string
}

type ShowSmartChatJobResponseStateEnum struct {
	WAITING    ShowSmartChatJobResponseState
	PROCESSING ShowSmartChatJobResponseState
	SUCCEED    ShowSmartChatJobResponseState
	FAILED     ShowSmartChatJobResponseState
	CANCELED   ShowSmartChatJobResponseState
	HEARTBEAT  ShowSmartChatJobResponseState
}

func GetShowSmartChatJobResponseStateEnum() ShowSmartChatJobResponseStateEnum {
	return ShowSmartChatJobResponseStateEnum{
		WAITING: ShowSmartChatJobResponseState{
			value: "WAITING",
		},
		PROCESSING: ShowSmartChatJobResponseState{
			value: "PROCESSING",
		},
		SUCCEED: ShowSmartChatJobResponseState{
			value: "SUCCEED",
		},
		FAILED: ShowSmartChatJobResponseState{
			value: "FAILED",
		},
		CANCELED: ShowSmartChatJobResponseState{
			value: "CANCELED",
		},
		HEARTBEAT: ShowSmartChatJobResponseState{
			value: "HEARTBEAT",
		},
	}
}

func (c ShowSmartChatJobResponseState) Value() string {
	return c.value
}

func (c ShowSmartChatJobResponseState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSmartChatJobResponseState) UnmarshalJSON(b []byte) error {
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

type ShowSmartChatJobResponseChatVideoType struct {
	value string
}

type ShowSmartChatJobResponseChatVideoTypeEnum struct {
	COMPUTER ShowSmartChatJobResponseChatVideoType
	MOBILE   ShowSmartChatJobResponseChatVideoType
	HUB      ShowSmartChatJobResponseChatVideoType
}

func GetShowSmartChatJobResponseChatVideoTypeEnum() ShowSmartChatJobResponseChatVideoTypeEnum {
	return ShowSmartChatJobResponseChatVideoTypeEnum{
		COMPUTER: ShowSmartChatJobResponseChatVideoType{
			value: "COMPUTER",
		},
		MOBILE: ShowSmartChatJobResponseChatVideoType{
			value: "MOBILE",
		},
		HUB: ShowSmartChatJobResponseChatVideoType{
			value: "HUB",
		},
	}
}

func (c ShowSmartChatJobResponseChatVideoType) Value() string {
	return c.value
}

func (c ShowSmartChatJobResponseChatVideoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSmartChatJobResponseChatVideoType) UnmarshalJSON(b []byte) error {
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
