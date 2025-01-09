package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AppSourceType 应用来源： * `CUSTOM` - 用户上传 * `SYSTEM` - 内置应用 * `MARKET` - 市场应用
type AppSourceType struct {
	value string
}

type AppSourceTypeEnum struct {
	CUSTOM AppSourceType
	SYSTEM AppSourceType
	MARKET AppSourceType
}

func GetAppSourceTypeEnum() AppSourceTypeEnum {
	return AppSourceTypeEnum{
		CUSTOM: AppSourceType{
			value: "CUSTOM",
		},
		SYSTEM: AppSourceType{
			value: "SYSTEM",
		},
		MARKET: AppSourceType{
			value: "MARKET",
		},
	}
}

func (c AppSourceType) Value() string {
	return c.value
}

func (c AppSourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AppSourceType) UnmarshalJSON(b []byte) error {
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
