package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// NotifyPolicyRequestBody 通知策略请求体
type NotifyPolicyRequestBody struct {

	// 通知策略类型，当前仅支持语音。
	Protocol NotifyPolicyRequestBodyProtocol `json:"protocol"`

	// 轮询策略订阅终端。
	Polling []PollingPolicyRequest `json:"polling"`
}

func (o NotifyPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NotifyPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"NotifyPolicyRequestBody", string(data)}, " ")
}

type NotifyPolicyRequestBodyProtocol struct {
	value string
}

type NotifyPolicyRequestBodyProtocolEnum struct {
	CALLNOTIFY NotifyPolicyRequestBodyProtocol
}

func GetNotifyPolicyRequestBodyProtocolEnum() NotifyPolicyRequestBodyProtocolEnum {
	return NotifyPolicyRequestBodyProtocolEnum{
		CALLNOTIFY: NotifyPolicyRequestBodyProtocol{
			value: "callnotify",
		},
	}
}

func (c NotifyPolicyRequestBodyProtocol) Value() string {
	return c.value
}

func (c NotifyPolicyRequestBodyProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NotifyPolicyRequestBodyProtocol) UnmarshalJSON(b []byte) error {
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
