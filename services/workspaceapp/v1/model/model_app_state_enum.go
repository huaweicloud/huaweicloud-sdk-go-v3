package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AppStateEnum 应用状态： * `NORMAL` - 正常状态。 * `FORBIDDEN` - 禁用状态。
type AppStateEnum struct {
	value string
}

type AppStateEnumEnum struct {
	NORMAL    AppStateEnum
	FORBIDDEN AppStateEnum
}

func GetAppStateEnumEnum() AppStateEnumEnum {
	return AppStateEnumEnum{
		NORMAL: AppStateEnum{
			value: "NORMAL",
		},
		FORBIDDEN: AppStateEnum{
			value: "FORBIDDEN",
		},
	}
}

func (c AppStateEnum) Value() string {
	return c.value
}

func (c AppStateEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AppStateEnum) UnmarshalJSON(b []byte) error {
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
