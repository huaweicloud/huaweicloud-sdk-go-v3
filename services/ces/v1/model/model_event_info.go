package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// EventInfo 一条事件监控信息
type EventInfo struct {

	// 事件名称。
	EventName *string `json:"event_name,omitempty"`

	// 枚举类型 EVENT.SYS或EVENT.CUSTOM，EVENT.SYS为系统事件，EVENT.CUSTOM为自定义事件
	EventType *EventInfoEventType `json:"event_type,omitempty"`

	// 事件子类。 枚举类型：SUB_EVENT.OPS为运维事件，SUB_EVENT.PLAN为计划事件，SUB_EVENT.CUSTOM为自定义事件。
	SubEventType *EventInfoSubEventType `json:"sub_event_type,omitempty"`

	// 选择查询的时间范围内，此事件发生的数量。
	EventCount *int32 `json:"event_count,omitempty"`

	// 此事件最近一次发生的时间。
	LatestOccurTime *int64 `json:"latest_occur_time,omitempty"`

	// 事件来源，如果是系统事件则值为各服务的命名空间，各服务命名空间请参阅[[支持监控的服务列表](https://support.huaweicloud.com/api-ces/ces_03_0059.html)](tag:hc)[[支持监控的服务列表](https://support.huaweicloud.com/intl/en-us/api-ces/ces_03_0059.html)](tag:hk)[[支持监控的服务列表](https://support.huaweicloud.com/eu/en-us/api-ces/ces_03_0059.html)](tag:hws_eu)[[支持监控的服务列表](ces_03_0059.xml)](tag:ax,cmcc,ctc,dt,dt_test,hcso_dt,fcs,fcs_vm,mix,g42,hk_g42,hk_sbc,hk_tm,hk_vdf,hws_ocb,ocb,sbc,srg)。；如果是自定义事件，则为用户自定义上报定义。
	LatestEventSource *string `json:"latest_event_source,omitempty"`
}

func (o EventInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EventInfo struct{}"
	}

	return strings.Join([]string{"EventInfo", string(data)}, " ")
}

type EventInfoEventType struct {
	value string
}

type EventInfoEventTypeEnum struct {
	EVENT_SYS    EventInfoEventType
	EVENT_CUSTOM EventInfoEventType
}

func GetEventInfoEventTypeEnum() EventInfoEventTypeEnum {
	return EventInfoEventTypeEnum{
		EVENT_SYS: EventInfoEventType{
			value: "EVENT.SYS",
		},
		EVENT_CUSTOM: EventInfoEventType{
			value: "EVENT.CUSTOM",
		},
	}
}

func (c EventInfoEventType) Value() string {
	return c.value
}

func (c EventInfoEventType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EventInfoEventType) UnmarshalJSON(b []byte) error {
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

type EventInfoSubEventType struct {
	value string
}

type EventInfoSubEventTypeEnum struct {
	SUB_EVENT_OPS    EventInfoSubEventType
	SUB_EVENT_PLAN   EventInfoSubEventType
	SUB_EVENT_CUSTOM EventInfoSubEventType
}

func GetEventInfoSubEventTypeEnum() EventInfoSubEventTypeEnum {
	return EventInfoSubEventTypeEnum{
		SUB_EVENT_OPS: EventInfoSubEventType{
			value: "SUB_EVENT.OPS",
		},
		SUB_EVENT_PLAN: EventInfoSubEventType{
			value: "SUB_EVENT.PLAN",
		},
		SUB_EVENT_CUSTOM: EventInfoSubEventType{
			value: "SUB_EVENT.CUSTOM",
		},
	}
}

func (c EventInfoSubEventType) Value() string {
	return c.value
}

func (c EventInfoSubEventType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EventInfoSubEventType) UnmarshalJSON(b []byte) error {
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
