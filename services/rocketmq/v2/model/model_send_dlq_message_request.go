package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SendDlqMessageRequest Request Object
type SendDlqMessageRequest struct {

	// 消息引擎。
	Engine SendDlqMessageRequestEngine `json:"engine"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *DeadletterResendReq `json:"body,omitempty"`
}

func (o SendDlqMessageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SendDlqMessageRequest struct{}"
	}

	return strings.Join([]string{"SendDlqMessageRequest", string(data)}, " ")
}

type SendDlqMessageRequestEngine struct {
	value string
}

type SendDlqMessageRequestEngineEnum struct {
	RELIABILITY SendDlqMessageRequestEngine
}

func GetSendDlqMessageRequestEngineEnum() SendDlqMessageRequestEngineEnum {
	return SendDlqMessageRequestEngineEnum{
		RELIABILITY: SendDlqMessageRequestEngine{
			value: "reliability",
		},
	}
}

func (c SendDlqMessageRequestEngine) Value() string {
	return c.value
}

func (c SendDlqMessageRequestEngine) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SendDlqMessageRequestEngine) UnmarshalJSON(b []byte) error {
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
