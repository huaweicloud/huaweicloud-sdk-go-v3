package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AppTypeEnum 服务器组应用类型： * `SESSION_DESKTOP_APP` - 会话桌面app * `COMMON_APP` - 普通app
type AppTypeEnum struct {
	value string
}

type AppTypeEnumEnum struct {
	SESSION_DESKTOP_APP AppTypeEnum
	COMMON_APP          AppTypeEnum
}

func GetAppTypeEnumEnum() AppTypeEnumEnum {
	return AppTypeEnumEnum{
		SESSION_DESKTOP_APP: AppTypeEnum{
			value: "SESSION_DESKTOP_APP",
		},
		COMMON_APP: AppTypeEnum{
			value: "COMMON_APP",
		},
	}
}

func (c AppTypeEnum) Value() string {
	return c.value
}

func (c AppTypeEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AppTypeEnum) UnmarshalJSON(b []byte) error {
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
