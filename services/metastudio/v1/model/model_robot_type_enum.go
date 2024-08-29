package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// RobotTypeEnum 交互对接类型 * LIVE:直播交互 * CHAT:智能交互
type RobotTypeEnum struct {
	value string
}

type RobotTypeEnumEnum struct {
	LIVE RobotTypeEnum
	CHAT RobotTypeEnum
}

func GetRobotTypeEnumEnum() RobotTypeEnumEnum {
	return RobotTypeEnumEnum{
		LIVE: RobotTypeEnum{
			value: "LIVE",
		},
		CHAT: RobotTypeEnum{
			value: "CHAT",
		},
	}
}

func (c RobotTypeEnum) Value() string {
	return c.value
}

func (c RobotTypeEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RobotTypeEnum) UnmarshalJSON(b []byte) error {
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
