package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// StartSmartLiveReq 数字人直播任务请求。
type StartSmartLiveReq struct {
	VideoConfig *VideoConfig `json:"video_config,omitempty"`

	PlayPolicy *PlayPolicy `json:"play_policy,omitempty"`

	// RTMP视频推流第三方直播平台地址。
	OutputUrls *[]string `json:"output_urls,omitempty"`

	// RTMP视频推流第三方直播平台流秘钥，与推流地址对应。
	StreamKeys *[]string `json:"stream_keys,omitempty"`

	// 互动回调URL，含鉴权信息。
	InteractionCallbackUrl *string `json:"interaction_callback_url,omitempty"`

	LiveEventCallbackConfig *LiveEventCallBackConfig `json:"live_event_callback_config,omitempty"`

	// 横竖屏类型。默认值为：VERTICAL。 * LANDSCAPE：横屏。 * VERTICAL： 竖屏。
	ViewMode *StartSmartLiveReqViewMode `json:"view_mode,omitempty"`
}

func (o StartSmartLiveReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartSmartLiveReq struct{}"
	}

	return strings.Join([]string{"StartSmartLiveReq", string(data)}, " ")
}

type StartSmartLiveReqViewMode struct {
	value string
}

type StartSmartLiveReqViewModeEnum struct {
	LANDSCAPE StartSmartLiveReqViewMode
	VERTICAL  StartSmartLiveReqViewMode
}

func GetStartSmartLiveReqViewModeEnum() StartSmartLiveReqViewModeEnum {
	return StartSmartLiveReqViewModeEnum{
		LANDSCAPE: StartSmartLiveReqViewMode{
			value: "LANDSCAPE",
		},
		VERTICAL: StartSmartLiveReqViewMode{
			value: "VERTICAL",
		},
	}
}

func (c StartSmartLiveReqViewMode) Value() string {
	return c.value
}

func (c StartSmartLiveReqViewMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StartSmartLiveReqViewMode) UnmarshalJSON(b []byte) error {
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
