package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// StartSmartChatJobResponse Response Object
type StartSmartChatJobResponse struct {

	// 智能交互对话任务ID。
	JobId *string `json:"job_id,omitempty"`

	RtcRoomInfo *RtcRoomInfoList `json:"rtc_room_info,omitempty"`

	ChatSubtitleConfig *ChatSubtitleConfig `json:"chat_subtitle_config,omitempty"`

	VideoConfig *ChatVideoConfigRsp `json:"video_config,omitempty"`

	// 智能交互对话端配置。 * COMPUTER: 电脑端 * MOBILE: 手机端 * HUB: 大屏
	ChatVideoType *StartSmartChatJobResponseChatVideoType `json:"chat_video_type,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o StartSmartChatJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartSmartChatJobResponse struct{}"
	}

	return strings.Join([]string{"StartSmartChatJobResponse", string(data)}, " ")
}

type StartSmartChatJobResponseChatVideoType struct {
	value string
}

type StartSmartChatJobResponseChatVideoTypeEnum struct {
	COMPUTER StartSmartChatJobResponseChatVideoType
	MOBILE   StartSmartChatJobResponseChatVideoType
	HUB      StartSmartChatJobResponseChatVideoType
}

func GetStartSmartChatJobResponseChatVideoTypeEnum() StartSmartChatJobResponseChatVideoTypeEnum {
	return StartSmartChatJobResponseChatVideoTypeEnum{
		COMPUTER: StartSmartChatJobResponseChatVideoType{
			value: "COMPUTER",
		},
		MOBILE: StartSmartChatJobResponseChatVideoType{
			value: "MOBILE",
		},
		HUB: StartSmartChatJobResponseChatVideoType{
			value: "HUB",
		},
	}
}

func (c StartSmartChatJobResponseChatVideoType) Value() string {
	return c.value
}

func (c StartSmartChatJobResponseChatVideoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StartSmartChatJobResponseChatVideoType) UnmarshalJSON(b []byte) error {
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
