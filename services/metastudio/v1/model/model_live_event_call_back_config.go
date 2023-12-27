package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// LiveEventCallBackConfig 直播事件回调通知配置
type LiveEventCallBackConfig struct {

	// 直播事件回调地址。https地址，需自带鉴权串。
	LiveEventTypeCallbackUrl *string `json:"live_event_type_callback_url,omitempty"`

	// 认证类型。 * NONE。URL中自带认证。 * MSS_A。HMACSHA256签名模式，在URL中追加参数:hwSecret,hwTime。取值方式：hwSecret=hmac_sha256(Key, URI（live_event_callback_url）+ hwTime)&hwTime=hex(timestamp)
	AuthType *LiveEventCallBackConfigAuthType `json:"auth_type,omitempty"`

	// 密钥Key
	Key *string `json:"key,omitempty"`

	// 回调的直播事件类型列表
	CallbackEventType *[]string `json:"callback_event_type,omitempty"`
}

func (o LiveEventCallBackConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LiveEventCallBackConfig struct{}"
	}

	return strings.Join([]string{"LiveEventCallBackConfig", string(data)}, " ")
}

type LiveEventCallBackConfigAuthType struct {
	value string
}

type LiveEventCallBackConfigAuthTypeEnum struct {
	NONE  LiveEventCallBackConfigAuthType
	MSS_A LiveEventCallBackConfigAuthType
}

func GetLiveEventCallBackConfigAuthTypeEnum() LiveEventCallBackConfigAuthTypeEnum {
	return LiveEventCallBackConfigAuthTypeEnum{
		NONE: LiveEventCallBackConfigAuthType{
			value: "NONE",
		},
		MSS_A: LiveEventCallBackConfigAuthType{
			value: "MSS_A",
		},
	}
}

func (c LiveEventCallBackConfigAuthType) Value() string {
	return c.value
}

func (c LiveEventCallBackConfigAuthType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LiveEventCallBackConfigAuthType) UnmarshalJSON(b []byte) error {
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
