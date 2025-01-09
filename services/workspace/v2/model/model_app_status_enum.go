package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AppStatusEnum 应用状态(正常、禁用) * 'NORMAL' - 正常 * 'FORBIDDEN' - 禁用状态
type AppStatusEnum struct {
	value string
}

type AppStatusEnumEnum struct {
	NORMAL    AppStatusEnum
	FORBIDDEN AppStatusEnum
}

func GetAppStatusEnumEnum() AppStatusEnumEnum {
	return AppStatusEnumEnum{
		NORMAL: AppStatusEnum{
			value: "NORMAL",
		},
		FORBIDDEN: AppStatusEnum{
			value: "FORBIDDEN",
		},
	}
}

func (c AppStatusEnum) Value() string {
	return c.value
}

func (c AppStatusEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AppStatusEnum) UnmarshalJSON(b []byte) error {
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
