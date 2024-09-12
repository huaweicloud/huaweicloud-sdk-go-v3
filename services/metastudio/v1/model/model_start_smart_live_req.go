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

	// **参数解释**： RTMP视频推流第三方直播平台地址。 > 直播过程中刷新地址，需要调用COMMAND命令REFRESH_OUTPUT_URL。  **约束限制**： 不涉及 **取值范围**： 当前仅支持一条RTMP出流地址。 **默认取值**： 不涉及。
	OutputUrls *[]string `json:"output_urls,omitempty"`

	// **参数解释**： RTMP视频推流第三方直播平台流密钥，与推流地址对应。 > 直播过程中刷新地址，需要调用COMMAND命令REFRESH_OUTPUT_URL。  **约束限制**： 不涉及 **取值范围**： 当前仅支持一条RTMP出流地址。 **默认取值**： 不涉及。
	StreamKeys *[]string `json:"stream_keys,omitempty"`

	// **参数解释**： 互动回调URL，含鉴权信息。 互动规则trigger.reply_mode配置为CALLBACK时填写 **约束限制**： 不涉及 **取值范围**： 字符长度0-2048位 **默认取值**： 不涉及。
	InteractionCallbackUrl *string `json:"interaction_callback_url,omitempty"`

	LiveEventCallbackConfig *LiveEventCallBackConfig `json:"live_event_callback_config,omitempty"`

	RtcCallbackConfig *RtcLiveEventCallBackConfig `json:"rtc_callback_config,omitempty"`

	// **参数解释**： 横竖屏类型。 **约束限制**： 用户无需填写，通过video_config中分辨率判断 **取值范围**： * LANDSCAPE：横屏。 * VERTICAL： 竖屏。
	ViewMode *StartSmartLiveReqViewMode `json:"view_mode,omitempty"`

	CoStreamerConfig *CoStreamerConfig `json:"co_streamer_config,omitempty"`
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
