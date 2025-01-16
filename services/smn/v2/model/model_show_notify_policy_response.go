package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowNotifyPolicyResponse Response Object
type ShowNotifyPolicyResponse struct {

	// 请求ID。
	RequestId *string `json:"request_id,omitempty"`

	// 通知策略ID。
	Id *string `json:"id,omitempty"`

	// 通知策略类型，当前仅支持语音。
	Protocol *ShowNotifyPolicyResponseProtocol `json:"protocol,omitempty"`

	// 轮询策略订阅终端。
	Polling        *[]PollingPolicyResponse `json:"polling,omitempty"`
	HttpStatusCode int                      `json:"-"`
}

func (o ShowNotifyPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowNotifyPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowNotifyPolicyResponse", string(data)}, " ")
}

type ShowNotifyPolicyResponseProtocol struct {
	value string
}

type ShowNotifyPolicyResponseProtocolEnum struct {
	CALLNOTIFY ShowNotifyPolicyResponseProtocol
}

func GetShowNotifyPolicyResponseProtocolEnum() ShowNotifyPolicyResponseProtocolEnum {
	return ShowNotifyPolicyResponseProtocolEnum{
		CALLNOTIFY: ShowNotifyPolicyResponseProtocol{
			value: "callnotify",
		},
	}
}

func (c ShowNotifyPolicyResponseProtocol) Value() string {
	return c.value
}

func (c ShowNotifyPolicyResponseProtocol) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowNotifyPolicyResponseProtocol) UnmarshalJSON(b []byte) error {
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
