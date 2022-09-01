package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// Response Object
type CreateFunctionTriggerResponse struct {

	// 触发器ID。
	TriggerId *string `json:"trigger_id,omitempty" xml:"trigger_id"`

	// 触发器类型。  - TIMER: \"定时触发器。\" - APIG: \"APIG触发器。\" - CTS: \"云审计服务触发器。\" - DDS: \"文档数据库服务触发器。\" - DMS: \"分布式服务触发器。\" - DIS: \"数据接入服务触发器。\" - LTS: \"云日志服务触发器。\" - OBS: \"对象存储触发器。\" - SMN: \"消息通知服务触发器。\" - KAFKA: \"专享版消息通知服务触发器。\"
	TriggerTypeCode *CreateFunctionTriggerResponseTriggerTypeCode `json:"trigger_type_code,omitempty" xml:"trigger_type_code"`

	// \"触发器状态\"  - ACTIVE: 启用状态。 - DISABLE: 禁用状态。
	TriggerStatus *CreateFunctionTriggerResponseTriggerStatus `json:"trigger_status,omitempty" xml:"trigger_status"`

	// 触发器源事件。
	EventData *interface{} `json:"event_data,omitempty" xml:"event_data"`

	// 最后更新时间。
	LastUpdatedTime *sdktime.SdkTime `json:"last_updated_time,omitempty" xml:"last_updated_time"`

	// 触发器创建时间。
	CreatedTime    *sdktime.SdkTime `json:"created_time,omitempty" xml:"created_time"`
	HttpStatusCode int              `json:"-"`
}

func (o CreateFunctionTriggerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateFunctionTriggerResponse struct{}"
	}

	return strings.Join([]string{"CreateFunctionTriggerResponse", string(data)}, " ")
}

type CreateFunctionTriggerResponseTriggerTypeCode struct {
	value string
}

type CreateFunctionTriggerResponseTriggerTypeCodeEnum struct {
	TIMER            CreateFunctionTriggerResponseTriggerTypeCode
	APIG             CreateFunctionTriggerResponseTriggerTypeCode
	CTS              CreateFunctionTriggerResponseTriggerTypeCode
	DDS              CreateFunctionTriggerResponseTriggerTypeCode
	DMS              CreateFunctionTriggerResponseTriggerTypeCode
	DIS              CreateFunctionTriggerResponseTriggerTypeCode
	LTS              CreateFunctionTriggerResponseTriggerTypeCode
	OBS              CreateFunctionTriggerResponseTriggerTypeCode
	SMN              CreateFunctionTriggerResponseTriggerTypeCode
	KAFKA            CreateFunctionTriggerResponseTriggerTypeCode
	RABBITMQ         CreateFunctionTriggerResponseTriggerTypeCode
	DEDICATEDGATEWAY CreateFunctionTriggerResponseTriggerTypeCode
	OPENSOURCEKAFKA  CreateFunctionTriggerResponseTriggerTypeCode
	APIC             CreateFunctionTriggerResponseTriggerTypeCode
	GAUSSMONGO       CreateFunctionTriggerResponseTriggerTypeCode
	EVENTGRID        CreateFunctionTriggerResponseTriggerTypeCode
}

func GetCreateFunctionTriggerResponseTriggerTypeCodeEnum() CreateFunctionTriggerResponseTriggerTypeCodeEnum {
	return CreateFunctionTriggerResponseTriggerTypeCodeEnum{
		TIMER: CreateFunctionTriggerResponseTriggerTypeCode{
			value: "TIMER",
		},
		APIG: CreateFunctionTriggerResponseTriggerTypeCode{
			value: "APIG",
		},
		CTS: CreateFunctionTriggerResponseTriggerTypeCode{
			value: "CTS",
		},
		DDS: CreateFunctionTriggerResponseTriggerTypeCode{
			value: "DDS",
		},
		DMS: CreateFunctionTriggerResponseTriggerTypeCode{
			value: "DMS",
		},
		DIS: CreateFunctionTriggerResponseTriggerTypeCode{
			value: "DIS",
		},
		LTS: CreateFunctionTriggerResponseTriggerTypeCode{
			value: "LTS",
		},
		OBS: CreateFunctionTriggerResponseTriggerTypeCode{
			value: "OBS",
		},
		SMN: CreateFunctionTriggerResponseTriggerTypeCode{
			value: "SMN",
		},
		KAFKA: CreateFunctionTriggerResponseTriggerTypeCode{
			value: "KAFKA",
		},
		RABBITMQ: CreateFunctionTriggerResponseTriggerTypeCode{
			value: "RABBITMQ",
		},
		DEDICATEDGATEWAY: CreateFunctionTriggerResponseTriggerTypeCode{
			value: "DEDICATEDGATEWAY",
		},
		OPENSOURCEKAFKA: CreateFunctionTriggerResponseTriggerTypeCode{
			value: "OPENSOURCEKAFKA",
		},
		APIC: CreateFunctionTriggerResponseTriggerTypeCode{
			value: "APIC",
		},
		GAUSSMONGO: CreateFunctionTriggerResponseTriggerTypeCode{
			value: "GAUSSMONGO",
		},
		EVENTGRID: CreateFunctionTriggerResponseTriggerTypeCode{
			value: "EVENTGRID",
		},
	}
}

func (c CreateFunctionTriggerResponseTriggerTypeCode) Value() string {
	return c.value
}

func (c CreateFunctionTriggerResponseTriggerTypeCode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateFunctionTriggerResponseTriggerTypeCode) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type CreateFunctionTriggerResponseTriggerStatus struct {
	value string
}

type CreateFunctionTriggerResponseTriggerStatusEnum struct {
	ACTIVE  CreateFunctionTriggerResponseTriggerStatus
	DISABLE CreateFunctionTriggerResponseTriggerStatus
}

func GetCreateFunctionTriggerResponseTriggerStatusEnum() CreateFunctionTriggerResponseTriggerStatusEnum {
	return CreateFunctionTriggerResponseTriggerStatusEnum{
		ACTIVE: CreateFunctionTriggerResponseTriggerStatus{
			value: "ACTIVE",
		},
		DISABLE: CreateFunctionTriggerResponseTriggerStatus{
			value: "DISABLE",
		},
	}
}

func (c CreateFunctionTriggerResponseTriggerStatus) Value() string {
	return c.value
}

func (c CreateFunctionTriggerResponseTriggerStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateFunctionTriggerResponseTriggerStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
