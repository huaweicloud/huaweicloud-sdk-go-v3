package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ValidateConsumedMessageRequest Request Object
type ValidateConsumedMessageRequest struct {

	// 消息引擎。
	Engine ValidateConsumedMessageRequestEngine `json:"engine"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *ResendReq `json:"body,omitempty"`
}

func (o ValidateConsumedMessageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateConsumedMessageRequest struct{}"
	}

	return strings.Join([]string{"ValidateConsumedMessageRequest", string(data)}, " ")
}

type ValidateConsumedMessageRequestEngine struct {
	value string
}

type ValidateConsumedMessageRequestEngineEnum struct {
	RELIABILITY ValidateConsumedMessageRequestEngine
}

func GetValidateConsumedMessageRequestEngineEnum() ValidateConsumedMessageRequestEngineEnum {
	return ValidateConsumedMessageRequestEngineEnum{
		RELIABILITY: ValidateConsumedMessageRequestEngine{
			value: "reliability",
		},
	}
}

func (c ValidateConsumedMessageRequestEngine) Value() string {
	return c.value
}

func (c ValidateConsumedMessageRequestEngine) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ValidateConsumedMessageRequestEngine) UnmarshalJSON(b []byte) error {
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
