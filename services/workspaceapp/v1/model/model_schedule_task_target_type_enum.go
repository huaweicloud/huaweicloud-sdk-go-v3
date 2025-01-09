package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ScheduleTaskTargetTypeEnum 定时任务对象类型： * `SERVER` - 服务器 * `SERVER_GROUP` - 服务器组
type ScheduleTaskTargetTypeEnum struct {
	value string
}

type ScheduleTaskTargetTypeEnumEnum struct {
	SERVER       ScheduleTaskTargetTypeEnum
	SERVER_GROUP ScheduleTaskTargetTypeEnum
}

func GetScheduleTaskTargetTypeEnumEnum() ScheduleTaskTargetTypeEnumEnum {
	return ScheduleTaskTargetTypeEnumEnum{
		SERVER: ScheduleTaskTargetTypeEnum{
			value: "SERVER",
		},
		SERVER_GROUP: ScheduleTaskTargetTypeEnum{
			value: "SERVER_GROUP",
		},
	}
}

func (c ScheduleTaskTargetTypeEnum) Value() string {
	return c.value
}

func (c ScheduleTaskTargetTypeEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ScheduleTaskTargetTypeEnum) UnmarshalJSON(b []byte) error {
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
