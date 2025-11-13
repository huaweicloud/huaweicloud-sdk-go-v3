package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// NotifyEventEnum * AUDIT_TEXT：直播在非人工退出场景（如最大时长中断，网络异常中断，欠费中断）中断。
type NotifyEventEnum struct {
	value string
}

type NotifyEventEnumEnum struct {
	AUTO_EXIT NotifyEventEnum
}

func GetNotifyEventEnumEnum() NotifyEventEnumEnum {
	return NotifyEventEnumEnum{
		AUTO_EXIT: NotifyEventEnum{
			value: "AUTO_EXIT",
		},
	}
}

func (c NotifyEventEnum) Value() string {
	return c.value
}

func (c NotifyEventEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NotifyEventEnum) UnmarshalJSON(b []byte) error {
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
