package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ScheduleTaskTypeEnum 定时任务类型： * `RESTART_SERVER` - 定时重启服务器 * `START_SERVER` - 定时开机 * `STOP_SERVER` - 定时关机 * `REINSTALL_OS` - 定时重装操作系统
type ScheduleTaskTypeEnum struct {
	value string
}

type ScheduleTaskTypeEnumEnum struct {
	RESTART_SERVER ScheduleTaskTypeEnum
	START_SERVER   ScheduleTaskTypeEnum
	STOP_SERVER    ScheduleTaskTypeEnum
	REINSTALL_OS   ScheduleTaskTypeEnum
}

func GetScheduleTaskTypeEnumEnum() ScheduleTaskTypeEnumEnum {
	return ScheduleTaskTypeEnumEnum{
		RESTART_SERVER: ScheduleTaskTypeEnum{
			value: "RESTART_SERVER",
		},
		START_SERVER: ScheduleTaskTypeEnum{
			value: "START_SERVER",
		},
		STOP_SERVER: ScheduleTaskTypeEnum{
			value: "STOP_SERVER",
		},
		REINSTALL_OS: ScheduleTaskTypeEnum{
			value: "REINSTALL_OS",
		},
	}
}

func (c ScheduleTaskTypeEnum) Value() string {
	return c.value
}

func (c ScheduleTaskTypeEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ScheduleTaskTypeEnum) UnmarshalJSON(b []byte) error {
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
