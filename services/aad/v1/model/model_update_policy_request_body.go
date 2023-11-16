package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdatePolicyRequestBody 更新策略的请求体
type UpdatePolicyRequestBody struct {

	// 策略名
	Name *string `json:"name,omitempty"`

	// 清洗阈值
	Threshold *int32 `json:"threshold,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// udp协议封禁。block：封禁，unblock：不封禁
	Udp *UpdatePolicyRequestBodyUdp `json:"udp,omitempty"`
}

func (o UpdatePolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePolicyRequestBody struct{}"
	}

	return strings.Join([]string{"UpdatePolicyRequestBody", string(data)}, " ")
}

type UpdatePolicyRequestBodyUdp struct {
	value string
}

type UpdatePolicyRequestBodyUdpEnum struct {
	BLOCK   UpdatePolicyRequestBodyUdp
	UNBLOCK UpdatePolicyRequestBodyUdp
}

func GetUpdatePolicyRequestBodyUdpEnum() UpdatePolicyRequestBodyUdpEnum {
	return UpdatePolicyRequestBodyUdpEnum{
		BLOCK: UpdatePolicyRequestBodyUdp{
			value: "block",
		},
		UNBLOCK: UpdatePolicyRequestBodyUdp{
			value: "unblock",
		},
	}
}

func (c UpdatePolicyRequestBodyUdp) Value() string {
	return c.value
}

func (c UpdatePolicyRequestBodyUdp) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdatePolicyRequestBodyUdp) UnmarshalJSON(b []byte) error {
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
