package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowSlaOrderResponse Response Object
type ShowSlaOrderResponse struct {

	// 主键
	Id *string `json:"id,omitempty"`

	// 工单标题
	OrderTitle *string `json:"order_title,omitempty"`

	// 工单ID
	OrderId *string `json:"order_id,omitempty"`

	// 触发类型(EVENT_TICKET,ALARM_TICKET,CHANGE_NOTE,TO_DO_TASKS,ISSUE_TICKET)
	TriggerType *ShowSlaOrderResponseTriggerType `json:"trigger_type,omitempty"`

	// 工单等级
	OrderLevel *string `json:"order_level,omitempty"`

	// 工单开始时间
	CreateTime *string `json:"create_time,omitempty"`

	// SLA规则记录
	SlaRecord      *[]SlaRecord `json:"sla_record,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowSlaOrderResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSlaOrderResponse struct{}"
	}

	return strings.Join([]string{"ShowSlaOrderResponse", string(data)}, " ")
}

type ShowSlaOrderResponseTriggerType struct {
	value string
}

type ShowSlaOrderResponseTriggerTypeEnum struct {
	EVENT_TICKET ShowSlaOrderResponseTriggerType
	ALARM_TICKET ShowSlaOrderResponseTriggerType
	CHANGE_NOTE  ShowSlaOrderResponseTriggerType
	TO_DO_TASKS  ShowSlaOrderResponseTriggerType
	ISSUE_TICKET ShowSlaOrderResponseTriggerType
}

func GetShowSlaOrderResponseTriggerTypeEnum() ShowSlaOrderResponseTriggerTypeEnum {
	return ShowSlaOrderResponseTriggerTypeEnum{
		EVENT_TICKET: ShowSlaOrderResponseTriggerType{
			value: "EVENT_TICKET",
		},
		ALARM_TICKET: ShowSlaOrderResponseTriggerType{
			value: "ALARM_TICKET",
		},
		CHANGE_NOTE: ShowSlaOrderResponseTriggerType{
			value: "CHANGE_NOTE",
		},
		TO_DO_TASKS: ShowSlaOrderResponseTriggerType{
			value: "TO_DO_TASKS",
		},
		ISSUE_TICKET: ShowSlaOrderResponseTriggerType{
			value: "ISSUE_TICKET",
		},
	}
}

func (c ShowSlaOrderResponseTriggerType) Value() string {
	return c.value
}

func (c ShowSlaOrderResponseTriggerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowSlaOrderResponseTriggerType) UnmarshalJSON(b []byte) error {
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
