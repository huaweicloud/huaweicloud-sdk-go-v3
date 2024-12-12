package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type LiveRoomEventCallBackConfig struct {

	// **参数解释**： 直播事件回调地址，为https地址。 **约束限制**： 不涉及。 **取值范围**： 字符长度0-2048位。 **默认取值**： 不涉及。
	LiveEventTypeCallbackUrl *string `json:"live_event_type_callback_url,omitempty"`

	// **参数解释**： 认证类型。 **约束限制**： 不涉及。 **取值范围**： * NONE：URL中自带认证。 * MSS_A：HMACSHA256签名模式，在URL中追加参数hwSecret、hwTime。   取值方式：hwSecret=hmac_sha256(Key, URI（live_event_callback_url）+ hwTime)&hwTime=hex(timestamp) * MSS_A_HEAD：HMACSHA256签名模式，参数hwSecret、hwTime放置在Head中。   取值方式：x-hw-mss-secret=hmac_sha256(Key, URI（live_event_callback_url）+ hwTime)     x-hw-mss-time=hex(timestamp) * MEITUAN_DEFAULT：仅用于美团平台调用回调使用。
	AuthType *LiveRoomEventCallBackConfigAuthType `json:"auth_type,omitempty"`

	// **参数解释**： 密钥Key。 **约束限制**： 不涉及。 **取值范围**： 字符长度0-32位。 **默认取值**： 不涉及。
	Key *string `json:"key,omitempty"`

	// **参数解释**： 回调的直播事件类型列表。 **约束限制**： 不涉及。 **取值范围**： 当前仅支持如下取值： * SHOOT_SCRIPT_SWITCH：剧本段落切换事件。  * RTMP_STREAM_STATE_CHANGE：RTMP链接发生变化回调事件。 * REPLY_COMMAND_FINISH：回复播放完成通知。  回调事件结构体定义： * event_type：事件类型。 * message：事件描述。   - SHOOT_SCRIPT_SWITCH事件回调定义如下：     ```json     {       \"event_type\":  \"SHOOT_SCRIPT_SWITCH\",       \"message\":\"{\\\"room_id\\\":\\\"26f065244f754b3aa853b649a21aaf66\\\",\\\"job_id\\\":\\\"e87104f76d7546ce8a46ac6b04c49c3c\\\",\\\"scene_script_name\\\":\\\"商品1\\\",\\\"shoot_script_sequence_no\\\":\\\"2\\\",\\\"shoot_script_title\\\":\\\"段落2\\\"}\"     }     ```   - RTMP_STREAM_STATE_CHANGE回调定义如下：     ```json     {       \"event_type\":  \"RTMP_STREAM_STATE_CHANGE\",       \"message\":\"{\\\"room_id\\\":\\\"26f065244f754b3aa853b649a21aaf66\\\",\\\"job_id\\\":\\\"e87104f76d7546ce8a46ac6b04c49c3c\\\",\\\"output_url\\\":\\\"rtmp://xxx/xx/xx\\\",\\\"stream_key\\\":\\\"xxxxx\\\",\\\"state\\\":\\\"CONNECTED\\\"}\"     }     ```     其中state取值：CONNECTING链路连接中；CONNECTED链路已连接；DISCONNECTED链路已断开，RECONNECTING链路重连中；END联络不再重连，链路已结束。    - REPLY_COMMAND_FINISH回调定义如下：     ```json     {       \"event_type\":  \"REPLY_COMMAND_FINISH\",       \"message\":\"{\\\"room_id\\\":\\\"26f065244f754b3aa853b649a21aaf66\\\",\\\"job_id\\\":\\\"e87104f76d7546ce8a46ac6b04c49c3c\\\",\\\"reply_id\\\":\\\"e87104f76d7546ce8a46ac6b04c49c3c\"}\"     }
	CallbackEventType *[]string `json:"callback_event_type,omitempty"`

	// **参数解释**： 是否继承直播平台回调配置，默认true **约束限制**： 不涉及。 **取值范围**： true：继承直播平台callback配置 false：不继承直播平台callback配置 **默认取值**： false
	IsInheritancePlatformCallback *bool `json:"is_inheritance_platform_callback,omitempty"`
}

func (o LiveRoomEventCallBackConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LiveRoomEventCallBackConfig struct{}"
	}

	return strings.Join([]string{"LiveRoomEventCallBackConfig", string(data)}, " ")
}

type LiveRoomEventCallBackConfigAuthType struct {
	value string
}

type LiveRoomEventCallBackConfigAuthTypeEnum struct {
	NONE            LiveRoomEventCallBackConfigAuthType
	MSS_A           LiveRoomEventCallBackConfigAuthType
	MSS_A_HEAD      LiveRoomEventCallBackConfigAuthType
	MEITUAN_DEFAULT LiveRoomEventCallBackConfigAuthType
}

func GetLiveRoomEventCallBackConfigAuthTypeEnum() LiveRoomEventCallBackConfigAuthTypeEnum {
	return LiveRoomEventCallBackConfigAuthTypeEnum{
		NONE: LiveRoomEventCallBackConfigAuthType{
			value: "NONE",
		},
		MSS_A: LiveRoomEventCallBackConfigAuthType{
			value: "MSS_A",
		},
		MSS_A_HEAD: LiveRoomEventCallBackConfigAuthType{
			value: "MSS_A_HEAD",
		},
		MEITUAN_DEFAULT: LiveRoomEventCallBackConfigAuthType{
			value: "MEITUAN_DEFAULT",
		},
	}
}

func (c LiveRoomEventCallBackConfigAuthType) Value() string {
	return c.value
}

func (c LiveRoomEventCallBackConfigAuthType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LiveRoomEventCallBackConfigAuthType) UnmarshalJSON(b []byte) error {
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
