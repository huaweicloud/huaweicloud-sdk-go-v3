package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// UpdateTriggerResponse Response Object
type UpdateTriggerResponse struct {

	// 触发器ID。
	TriggerId *string `json:"trigger_id,omitempty"`

	// 触发器类型。  - TIMER: \"定时触发器。\" - APIG: \"APIG触发器。\" - CTS: \"云审计服务触发器。\" - DDS: \"文档数据库服务触发器。\" - DMS: \"分布式服务触发器。\" - DIS: \"数据接入服务触发器。\" - LTS: \"云日志服务触发器。\" - OBS: \"对象存储触发器。\" - SMN: \"消息通知服务触发器。\" - KAFKA: \"专享版消息通知服务触发器。\"
	TriggerTypeCode *UpdateTriggerResponseTriggerTypeCode `json:"trigger_type_code,omitempty"`

	// \"触发器状态\"  - ACTIVE: 启用状态。 - DISABLED: 禁用状态。
	TriggerStatus *UpdateTriggerResponseTriggerStatus `json:"trigger_status,omitempty"`

	EventData *TriggerEventDataResponseBody `json:"event_data,omitempty"`

	// 最后更新时间。
	LastUpdatedTime *sdktime.SdkTime `json:"last_updated_time,omitempty"`

	// 触发器创建时间。
	CreatedTime    *sdktime.SdkTime `json:"created_time,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o UpdateTriggerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTriggerResponse struct{}"
	}

	return strings.Join([]string{"UpdateTriggerResponse", string(data)}, " ")
}

type UpdateTriggerResponseTriggerTypeCode struct {
	value string
}

type UpdateTriggerResponseTriggerTypeCodeEnum struct {
	TIMER            UpdateTriggerResponseTriggerTypeCode
	APIG             UpdateTriggerResponseTriggerTypeCode
	CTS              UpdateTriggerResponseTriggerTypeCode
	DDS              UpdateTriggerResponseTriggerTypeCode
	DMS              UpdateTriggerResponseTriggerTypeCode
	DIS              UpdateTriggerResponseTriggerTypeCode
	LTS              UpdateTriggerResponseTriggerTypeCode
	OBS              UpdateTriggerResponseTriggerTypeCode
	SMN              UpdateTriggerResponseTriggerTypeCode
	KAFKA            UpdateTriggerResponseTriggerTypeCode
	RABBITMQ         UpdateTriggerResponseTriggerTypeCode
	DEDICATEDGATEWAY UpdateTriggerResponseTriggerTypeCode
	OPENSOURCEKAFKA  UpdateTriggerResponseTriggerTypeCode
	APIC             UpdateTriggerResponseTriggerTypeCode
	GAUSSMONGO       UpdateTriggerResponseTriggerTypeCode
	EVENTGRID        UpdateTriggerResponseTriggerTypeCode
}

func GetUpdateTriggerResponseTriggerTypeCodeEnum() UpdateTriggerResponseTriggerTypeCodeEnum {
	return UpdateTriggerResponseTriggerTypeCodeEnum{
		TIMER: UpdateTriggerResponseTriggerTypeCode{
			value: "TIMER",
		},
		APIG: UpdateTriggerResponseTriggerTypeCode{
			value: "APIG",
		},
		CTS: UpdateTriggerResponseTriggerTypeCode{
			value: "CTS",
		},
		DDS: UpdateTriggerResponseTriggerTypeCode{
			value: "DDS",
		},
		DMS: UpdateTriggerResponseTriggerTypeCode{
			value: "DMS",
		},
		DIS: UpdateTriggerResponseTriggerTypeCode{
			value: "DIS",
		},
		LTS: UpdateTriggerResponseTriggerTypeCode{
			value: "LTS",
		},
		OBS: UpdateTriggerResponseTriggerTypeCode{
			value: "OBS",
		},
		SMN: UpdateTriggerResponseTriggerTypeCode{
			value: "SMN",
		},
		KAFKA: UpdateTriggerResponseTriggerTypeCode{
			value: "KAFKA",
		},
		RABBITMQ: UpdateTriggerResponseTriggerTypeCode{
			value: "RABBITMQ",
		},
		DEDICATEDGATEWAY: UpdateTriggerResponseTriggerTypeCode{
			value: "DEDICATEDGATEWAY",
		},
		OPENSOURCEKAFKA: UpdateTriggerResponseTriggerTypeCode{
			value: "OPENSOURCEKAFKA",
		},
		APIC: UpdateTriggerResponseTriggerTypeCode{
			value: "APIC",
		},
		GAUSSMONGO: UpdateTriggerResponseTriggerTypeCode{
			value: "GAUSSMONGO",
		},
		EVENTGRID: UpdateTriggerResponseTriggerTypeCode{
			value: "EVENTGRID",
		},
	}
}

func (c UpdateTriggerResponseTriggerTypeCode) Value() string {
	return c.value
}

func (c UpdateTriggerResponseTriggerTypeCode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateTriggerResponseTriggerTypeCode) UnmarshalJSON(b []byte) error {
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

type UpdateTriggerResponseTriggerStatus struct {
	value string
}

type UpdateTriggerResponseTriggerStatusEnum struct {
	ACTIVE   UpdateTriggerResponseTriggerStatus
	DISABLED UpdateTriggerResponseTriggerStatus
}

func GetUpdateTriggerResponseTriggerStatusEnum() UpdateTriggerResponseTriggerStatusEnum {
	return UpdateTriggerResponseTriggerStatusEnum{
		ACTIVE: UpdateTriggerResponseTriggerStatus{
			value: "ACTIVE",
		},
		DISABLED: UpdateTriggerResponseTriggerStatus{
			value: "DISABLED",
		},
	}
}

func (c UpdateTriggerResponseTriggerStatus) Value() string {
	return c.value
}

func (c UpdateTriggerResponseTriggerStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateTriggerResponseTriggerStatus) UnmarshalJSON(b []byte) error {
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
