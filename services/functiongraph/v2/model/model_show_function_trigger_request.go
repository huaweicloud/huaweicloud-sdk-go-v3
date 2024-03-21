package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowFunctionTriggerRequest Request Object
type ShowFunctionTriggerRequest struct {

	// 函数的URN，详细解释见FunctionGraph函数模型的描述。
	FunctionUrn string `json:"function_urn"`

	// 触发器类型代码。
	TriggerTypeCode ShowFunctionTriggerRequestTriggerTypeCode `json:"trigger_type_code"`

	// 触发器编码。
	TriggerId string `json:"trigger_id"`

	// 消息体的类型（格式）
	ContentType string `json:"Content-Type"`
}

func (o ShowFunctionTriggerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowFunctionTriggerRequest struct{}"
	}

	return strings.Join([]string{"ShowFunctionTriggerRequest", string(data)}, " ")
}

type ShowFunctionTriggerRequestTriggerTypeCode struct {
	value string
}

type ShowFunctionTriggerRequestTriggerTypeCodeEnum struct {
	TIMER            ShowFunctionTriggerRequestTriggerTypeCode
	APIG             ShowFunctionTriggerRequestTriggerTypeCode
	CTS              ShowFunctionTriggerRequestTriggerTypeCode
	DDS              ShowFunctionTriggerRequestTriggerTypeCode
	DMS              ShowFunctionTriggerRequestTriggerTypeCode
	DIS              ShowFunctionTriggerRequestTriggerTypeCode
	LTS              ShowFunctionTriggerRequestTriggerTypeCode
	OBS              ShowFunctionTriggerRequestTriggerTypeCode
	SMN              ShowFunctionTriggerRequestTriggerTypeCode
	KAFKA            ShowFunctionTriggerRequestTriggerTypeCode
	RABBITMQ         ShowFunctionTriggerRequestTriggerTypeCode
	DEDICATEDGATEWAY ShowFunctionTriggerRequestTriggerTypeCode
	OPENSOURCEKAFKA  ShowFunctionTriggerRequestTriggerTypeCode
	APIC             ShowFunctionTriggerRequestTriggerTypeCode
	GAUSSMONGO       ShowFunctionTriggerRequestTriggerTypeCode
	EVENTGRID        ShowFunctionTriggerRequestTriggerTypeCode
	IOTDA            ShowFunctionTriggerRequestTriggerTypeCode
}

func GetShowFunctionTriggerRequestTriggerTypeCodeEnum() ShowFunctionTriggerRequestTriggerTypeCodeEnum {
	return ShowFunctionTriggerRequestTriggerTypeCodeEnum{
		TIMER: ShowFunctionTriggerRequestTriggerTypeCode{
			value: "TIMER",
		},
		APIG: ShowFunctionTriggerRequestTriggerTypeCode{
			value: "APIG",
		},
		CTS: ShowFunctionTriggerRequestTriggerTypeCode{
			value: "CTS",
		},
		DDS: ShowFunctionTriggerRequestTriggerTypeCode{
			value: "DDS",
		},
		DMS: ShowFunctionTriggerRequestTriggerTypeCode{
			value: "DMS",
		},
		DIS: ShowFunctionTriggerRequestTriggerTypeCode{
			value: "DIS",
		},
		LTS: ShowFunctionTriggerRequestTriggerTypeCode{
			value: "LTS",
		},
		OBS: ShowFunctionTriggerRequestTriggerTypeCode{
			value: "OBS",
		},
		SMN: ShowFunctionTriggerRequestTriggerTypeCode{
			value: "SMN",
		},
		KAFKA: ShowFunctionTriggerRequestTriggerTypeCode{
			value: "KAFKA",
		},
		RABBITMQ: ShowFunctionTriggerRequestTriggerTypeCode{
			value: "RABBITMQ",
		},
		DEDICATEDGATEWAY: ShowFunctionTriggerRequestTriggerTypeCode{
			value: "DEDICATEDGATEWAY",
		},
		OPENSOURCEKAFKA: ShowFunctionTriggerRequestTriggerTypeCode{
			value: "OPENSOURCEKAFKA",
		},
		APIC: ShowFunctionTriggerRequestTriggerTypeCode{
			value: "APIC",
		},
		GAUSSMONGO: ShowFunctionTriggerRequestTriggerTypeCode{
			value: "GAUSSMONGO",
		},
		EVENTGRID: ShowFunctionTriggerRequestTriggerTypeCode{
			value: "EVENTGRID",
		},
		IOTDA: ShowFunctionTriggerRequestTriggerTypeCode{
			value: "IOTDA",
		},
	}
}

func (c ShowFunctionTriggerRequestTriggerTypeCode) Value() string {
	return c.value
}

func (c ShowFunctionTriggerRequestTriggerTypeCode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowFunctionTriggerRequestTriggerTypeCode) UnmarshalJSON(b []byte) error {
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
