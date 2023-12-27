package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateShrinkageJobRequest Request Object
type CreateShrinkageJobRequest struct {

	// 消息引擎。
	Engine CreateShrinkageJobRequestEngine `json:"engine"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *CreateShrinkageJobRequestBody `json:"body,omitempty"`
}

func (o CreateShrinkageJobRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateShrinkageJobRequest struct{}"
	}

	return strings.Join([]string{"CreateShrinkageJobRequest", string(data)}, " ")
}

type CreateShrinkageJobRequestEngine struct {
	value string
}

type CreateShrinkageJobRequestEngineEnum struct {
	KAFKA CreateShrinkageJobRequestEngine
}

func GetCreateShrinkageJobRequestEngineEnum() CreateShrinkageJobRequestEngineEnum {
	return CreateShrinkageJobRequestEngineEnum{
		KAFKA: CreateShrinkageJobRequestEngine{
			value: "kafka",
		},
	}
}

func (c CreateShrinkageJobRequestEngine) Value() string {
	return c.value
}

func (c CreateShrinkageJobRequestEngine) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateShrinkageJobRequestEngine) UnmarshalJSON(b []byte) error {
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
