package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateTriggerRequest Request Object
type UpdateTriggerRequest struct {

	// 函数的URN，详细解释见FunctionGraph函数模型的描述。
	FunctionUrn string `json:"function_urn"`

	// 触发器类型代码。
	TriggerTypeCode UpdateTriggerRequestTriggerTypeCode `json:"trigger_type_code"`

	// 触发器编码。
	TriggerId string `json:"trigger_id"`

	// 消息体的类型（格式）
	ContentType string `json:"Content-Type"`

	Body *UpdateTriggerRequestBody `json:"body,omitempty"`
}

func (o UpdateTriggerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTriggerRequest struct{}"
	}

	return strings.Join([]string{"UpdateTriggerRequest", string(data)}, " ")
}

type UpdateTriggerRequestTriggerTypeCode struct {
	value string
}

type UpdateTriggerRequestTriggerTypeCodeEnum struct {
	TIMER            UpdateTriggerRequestTriggerTypeCode
	APIG             UpdateTriggerRequestTriggerTypeCode
	CTS              UpdateTriggerRequestTriggerTypeCode
	DDS              UpdateTriggerRequestTriggerTypeCode
	DMS              UpdateTriggerRequestTriggerTypeCode
	DIS              UpdateTriggerRequestTriggerTypeCode
	LTS              UpdateTriggerRequestTriggerTypeCode
	OBS              UpdateTriggerRequestTriggerTypeCode
	SMN              UpdateTriggerRequestTriggerTypeCode
	KAFKA            UpdateTriggerRequestTriggerTypeCode
	RABBITMQ         UpdateTriggerRequestTriggerTypeCode
	DEDICATEDGATEWAY UpdateTriggerRequestTriggerTypeCode
	OPENSOURCEKAFKA  UpdateTriggerRequestTriggerTypeCode
	APIC             UpdateTriggerRequestTriggerTypeCode
	GAUSSMONGO       UpdateTriggerRequestTriggerTypeCode
	EVENTGRID        UpdateTriggerRequestTriggerTypeCode
	IOTDA            UpdateTriggerRequestTriggerTypeCode
}

func GetUpdateTriggerRequestTriggerTypeCodeEnum() UpdateTriggerRequestTriggerTypeCodeEnum {
	return UpdateTriggerRequestTriggerTypeCodeEnum{
		TIMER: UpdateTriggerRequestTriggerTypeCode{
			value: "TIMER",
		},
		APIG: UpdateTriggerRequestTriggerTypeCode{
			value: "APIG",
		},
		CTS: UpdateTriggerRequestTriggerTypeCode{
			value: "CTS",
		},
		DDS: UpdateTriggerRequestTriggerTypeCode{
			value: "DDS",
		},
		DMS: UpdateTriggerRequestTriggerTypeCode{
			value: "DMS",
		},
		DIS: UpdateTriggerRequestTriggerTypeCode{
			value: "DIS",
		},
		LTS: UpdateTriggerRequestTriggerTypeCode{
			value: "LTS",
		},
		OBS: UpdateTriggerRequestTriggerTypeCode{
			value: "OBS",
		},
		SMN: UpdateTriggerRequestTriggerTypeCode{
			value: "SMN",
		},
		KAFKA: UpdateTriggerRequestTriggerTypeCode{
			value: "KAFKA",
		},
		RABBITMQ: UpdateTriggerRequestTriggerTypeCode{
			value: "RABBITMQ",
		},
		DEDICATEDGATEWAY: UpdateTriggerRequestTriggerTypeCode{
			value: "DEDICATEDGATEWAY",
		},
		OPENSOURCEKAFKA: UpdateTriggerRequestTriggerTypeCode{
			value: "OPENSOURCEKAFKA",
		},
		APIC: UpdateTriggerRequestTriggerTypeCode{
			value: "APIC",
		},
		GAUSSMONGO: UpdateTriggerRequestTriggerTypeCode{
			value: "GAUSSMONGO",
		},
		EVENTGRID: UpdateTriggerRequestTriggerTypeCode{
			value: "EVENTGRID",
		},
		IOTDA: UpdateTriggerRequestTriggerTypeCode{
			value: "IOTDA",
		},
	}
}

func (c UpdateTriggerRequestTriggerTypeCode) Value() string {
	return c.value
}

func (c UpdateTriggerRequestTriggerTypeCode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateTriggerRequestTriggerTypeCode) UnmarshalJSON(b []byte) error {
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
