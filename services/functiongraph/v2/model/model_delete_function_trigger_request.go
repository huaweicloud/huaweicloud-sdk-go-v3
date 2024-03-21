package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteFunctionTriggerRequest Request Object
type DeleteFunctionTriggerRequest struct {

	// 函数的URN，详细解释见FunctionGraph函数模型的描述。
	FunctionUrn string `json:"function_urn"`

	// 触发器类型代码。
	TriggerTypeCode DeleteFunctionTriggerRequestTriggerTypeCode `json:"trigger_type_code"`

	// 触发器编码。
	TriggerId string `json:"trigger_id"`

	// 消息体的类型（格式）
	ContentType string `json:"Content-Type"`
}

func (o DeleteFunctionTriggerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteFunctionTriggerRequest struct{}"
	}

	return strings.Join([]string{"DeleteFunctionTriggerRequest", string(data)}, " ")
}

type DeleteFunctionTriggerRequestTriggerTypeCode struct {
	value string
}

type DeleteFunctionTriggerRequestTriggerTypeCodeEnum struct {
	TIMER            DeleteFunctionTriggerRequestTriggerTypeCode
	APIG             DeleteFunctionTriggerRequestTriggerTypeCode
	CTS              DeleteFunctionTriggerRequestTriggerTypeCode
	DDS              DeleteFunctionTriggerRequestTriggerTypeCode
	DMS              DeleteFunctionTriggerRequestTriggerTypeCode
	DIS              DeleteFunctionTriggerRequestTriggerTypeCode
	LTS              DeleteFunctionTriggerRequestTriggerTypeCode
	OBS              DeleteFunctionTriggerRequestTriggerTypeCode
	SMN              DeleteFunctionTriggerRequestTriggerTypeCode
	KAFKA            DeleteFunctionTriggerRequestTriggerTypeCode
	RABBITMQ         DeleteFunctionTriggerRequestTriggerTypeCode
	DEDICATEDGATEWAY DeleteFunctionTriggerRequestTriggerTypeCode
	OPENSOURCEKAFKA  DeleteFunctionTriggerRequestTriggerTypeCode
	APIC             DeleteFunctionTriggerRequestTriggerTypeCode
	GAUSSMONGO       DeleteFunctionTriggerRequestTriggerTypeCode
	EVENTGRID        DeleteFunctionTriggerRequestTriggerTypeCode
	IOTDA            DeleteFunctionTriggerRequestTriggerTypeCode
}

func GetDeleteFunctionTriggerRequestTriggerTypeCodeEnum() DeleteFunctionTriggerRequestTriggerTypeCodeEnum {
	return DeleteFunctionTriggerRequestTriggerTypeCodeEnum{
		TIMER: DeleteFunctionTriggerRequestTriggerTypeCode{
			value: "TIMER",
		},
		APIG: DeleteFunctionTriggerRequestTriggerTypeCode{
			value: "APIG",
		},
		CTS: DeleteFunctionTriggerRequestTriggerTypeCode{
			value: "CTS",
		},
		DDS: DeleteFunctionTriggerRequestTriggerTypeCode{
			value: "DDS",
		},
		DMS: DeleteFunctionTriggerRequestTriggerTypeCode{
			value: "DMS",
		},
		DIS: DeleteFunctionTriggerRequestTriggerTypeCode{
			value: "DIS",
		},
		LTS: DeleteFunctionTriggerRequestTriggerTypeCode{
			value: "LTS",
		},
		OBS: DeleteFunctionTriggerRequestTriggerTypeCode{
			value: "OBS",
		},
		SMN: DeleteFunctionTriggerRequestTriggerTypeCode{
			value: "SMN",
		},
		KAFKA: DeleteFunctionTriggerRequestTriggerTypeCode{
			value: "KAFKA",
		},
		RABBITMQ: DeleteFunctionTriggerRequestTriggerTypeCode{
			value: "RABBITMQ",
		},
		DEDICATEDGATEWAY: DeleteFunctionTriggerRequestTriggerTypeCode{
			value: "DEDICATEDGATEWAY",
		},
		OPENSOURCEKAFKA: DeleteFunctionTriggerRequestTriggerTypeCode{
			value: "OPENSOURCEKAFKA",
		},
		APIC: DeleteFunctionTriggerRequestTriggerTypeCode{
			value: "APIC",
		},
		GAUSSMONGO: DeleteFunctionTriggerRequestTriggerTypeCode{
			value: "GAUSSMONGO",
		},
		EVENTGRID: DeleteFunctionTriggerRequestTriggerTypeCode{
			value: "EVENTGRID",
		},
		IOTDA: DeleteFunctionTriggerRequestTriggerTypeCode{
			value: "IOTDA",
		},
	}
}

func (c DeleteFunctionTriggerRequestTriggerTypeCode) Value() string {
	return c.value
}

func (c DeleteFunctionTriggerRequestTriggerTypeCode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteFunctionTriggerRequestTriggerTypeCode) UnmarshalJSON(b []byte) error {
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
