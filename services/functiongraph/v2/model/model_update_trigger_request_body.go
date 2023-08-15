package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateTriggerRequestBody struct {

	// 触发器状态
	TriggerStatus *UpdateTriggerRequestBodyTriggerStatus `json:"trigger_status,omitempty"`

	// 触发器更新事件
	EventData *[]TriggerEventData `json:"event_data,omitempty"`
}

func (o UpdateTriggerRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateTriggerRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateTriggerRequestBody", string(data)}, " ")
}

type UpdateTriggerRequestBodyTriggerStatus struct {
	value string
}

type UpdateTriggerRequestBodyTriggerStatusEnum struct {
	ACTIVE   UpdateTriggerRequestBodyTriggerStatus
	DISABLED UpdateTriggerRequestBodyTriggerStatus
}

func GetUpdateTriggerRequestBodyTriggerStatusEnum() UpdateTriggerRequestBodyTriggerStatusEnum {
	return UpdateTriggerRequestBodyTriggerStatusEnum{
		ACTIVE: UpdateTriggerRequestBodyTriggerStatus{
			value: "ACTIVE",
		},
		DISABLED: UpdateTriggerRequestBodyTriggerStatus{
			value: "DISABLED",
		},
	}
}

func (c UpdateTriggerRequestBodyTriggerStatus) Value() string {
	return c.value
}

func (c UpdateTriggerRequestBodyTriggerStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateTriggerRequestBodyTriggerStatus) UnmarshalJSON(b []byte) error {
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
