package model

import (
	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

type ListFunctionTriggerResult struct {

	// 触发器ID。
	TriggerId string `json:"trigger_id"`

	// 触发器类型。  - TIMER: \"定时触发器。\" - APIG: \"APIG触发器。\" - CTS: \"云审计服务触发器。\" - DDS: \"文档数据库服务触发器。\" - DMS: \"分布式服务触发器。\" - DIS: \"数据接入服务触发器。\" - LTS: \"云日志服务触发器。\" - OBS: \"对象存储触发器。\" - SMN: \"消息通知服务触发器。\" - KAFKA: \"专享版消息通知服务触发器。\"
	TriggerTypeCode ListFunctionTriggerResultTriggerTypeCode `json:"trigger_type_code"`

	// \"触发器状态\"  - ACTIVE: 启用状态。 - DISABLED: 禁用状态。
	TriggerStatus ListFunctionTriggerResultTriggerStatus `json:"trigger_status"`

	// 触发器源事件。
	EventData *interface{} `json:"event_data"`

	// 最后更新时间。
	LastUpdatedTime *sdktime.SdkTime `json:"last_updated_time"`

	// 触发器创建时间。
	CreatedTime *sdktime.SdkTime `json:"created_time"`
}

func (o ListFunctionTriggerResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFunctionTriggerResult struct{}"
	}

	return strings.Join([]string{"ListFunctionTriggerResult", string(data)}, " ")
}

type ListFunctionTriggerResultTriggerTypeCode struct {
	value string
}

type ListFunctionTriggerResultTriggerTypeCodeEnum struct {
	TIMER            ListFunctionTriggerResultTriggerTypeCode
	APIG             ListFunctionTriggerResultTriggerTypeCode
	CTS              ListFunctionTriggerResultTriggerTypeCode
	DDS              ListFunctionTriggerResultTriggerTypeCode
	DMS              ListFunctionTriggerResultTriggerTypeCode
	DIS              ListFunctionTriggerResultTriggerTypeCode
	LTS              ListFunctionTriggerResultTriggerTypeCode
	OBS              ListFunctionTriggerResultTriggerTypeCode
	SMN              ListFunctionTriggerResultTriggerTypeCode
	KAFKA            ListFunctionTriggerResultTriggerTypeCode
	RABBITMQ         ListFunctionTriggerResultTriggerTypeCode
	DEDICATEDGATEWAY ListFunctionTriggerResultTriggerTypeCode
	OPENSOURCEKAFKA  ListFunctionTriggerResultTriggerTypeCode
	APIC             ListFunctionTriggerResultTriggerTypeCode
	GAUSSMONGO       ListFunctionTriggerResultTriggerTypeCode
	EVENTGRID        ListFunctionTriggerResultTriggerTypeCode
}

func GetListFunctionTriggerResultTriggerTypeCodeEnum() ListFunctionTriggerResultTriggerTypeCodeEnum {
	return ListFunctionTriggerResultTriggerTypeCodeEnum{
		TIMER: ListFunctionTriggerResultTriggerTypeCode{
			value: "TIMER",
		},
		APIG: ListFunctionTriggerResultTriggerTypeCode{
			value: "APIG",
		},
		CTS: ListFunctionTriggerResultTriggerTypeCode{
			value: "CTS",
		},
		DDS: ListFunctionTriggerResultTriggerTypeCode{
			value: "DDS",
		},
		DMS: ListFunctionTriggerResultTriggerTypeCode{
			value: "DMS",
		},
		DIS: ListFunctionTriggerResultTriggerTypeCode{
			value: "DIS",
		},
		LTS: ListFunctionTriggerResultTriggerTypeCode{
			value: "LTS",
		},
		OBS: ListFunctionTriggerResultTriggerTypeCode{
			value: "OBS",
		},
		SMN: ListFunctionTriggerResultTriggerTypeCode{
			value: "SMN",
		},
		KAFKA: ListFunctionTriggerResultTriggerTypeCode{
			value: "KAFKA",
		},
		RABBITMQ: ListFunctionTriggerResultTriggerTypeCode{
			value: "RABBITMQ",
		},
		DEDICATEDGATEWAY: ListFunctionTriggerResultTriggerTypeCode{
			value: "DEDICATEDGATEWAY",
		},
		OPENSOURCEKAFKA: ListFunctionTriggerResultTriggerTypeCode{
			value: "OPENSOURCEKAFKA",
		},
		APIC: ListFunctionTriggerResultTriggerTypeCode{
			value: "APIC",
		},
		GAUSSMONGO: ListFunctionTriggerResultTriggerTypeCode{
			value: "GAUSSMONGO",
		},
		EVENTGRID: ListFunctionTriggerResultTriggerTypeCode{
			value: "EVENTGRID",
		},
	}
}

func (c ListFunctionTriggerResultTriggerTypeCode) Value() string {
	return c.value
}

func (c ListFunctionTriggerResultTriggerTypeCode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListFunctionTriggerResultTriggerTypeCode) UnmarshalJSON(b []byte) error {
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

type ListFunctionTriggerResultTriggerStatus struct {
	value string
}

type ListFunctionTriggerResultTriggerStatusEnum struct {
	ACTIVE   ListFunctionTriggerResultTriggerStatus
	DISABLED ListFunctionTriggerResultTriggerStatus
}

func GetListFunctionTriggerResultTriggerStatusEnum() ListFunctionTriggerResultTriggerStatusEnum {
	return ListFunctionTriggerResultTriggerStatusEnum{
		ACTIVE: ListFunctionTriggerResultTriggerStatus{
			value: "ACTIVE",
		},
		DISABLED: ListFunctionTriggerResultTriggerStatus{
			value: "DISABLED",
		},
	}
}

func (c ListFunctionTriggerResultTriggerStatus) Value() string {
	return c.value
}

func (c ListFunctionTriggerResultTriggerStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListFunctionTriggerResultTriggerStatus) UnmarshalJSON(b []byte) error {
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
