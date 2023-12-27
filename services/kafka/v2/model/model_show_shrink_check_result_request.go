package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowShrinkCheckResultRequest Request Object
type ShowShrinkCheckResultRequest struct {

	// 消息引擎。
	Engine ShowShrinkCheckResultRequestEngine `json:"engine"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *ShowShrinkCheckRequestBody `json:"body,omitempty"`
}

func (o ShowShrinkCheckResultRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowShrinkCheckResultRequest struct{}"
	}

	return strings.Join([]string{"ShowShrinkCheckResultRequest", string(data)}, " ")
}

type ShowShrinkCheckResultRequestEngine struct {
	value string
}

type ShowShrinkCheckResultRequestEngineEnum struct {
	KAFKA ShowShrinkCheckResultRequestEngine
}

func GetShowShrinkCheckResultRequestEngineEnum() ShowShrinkCheckResultRequestEngineEnum {
	return ShowShrinkCheckResultRequestEngineEnum{
		KAFKA: ShowShrinkCheckResultRequestEngine{
			value: "kafka",
		},
	}
}

func (c ShowShrinkCheckResultRequestEngine) Value() string {
	return c.value
}

func (c ShowShrinkCheckResultRequestEngine) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowShrinkCheckResultRequestEngine) UnmarshalJSON(b []byte) error {
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
