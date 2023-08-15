package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type DownloadDataTypeEnum struct {
	value string
}

type DownloadDataTypeEnumEnum struct {
	PRIVATE DownloadDataTypeEnum
	PUBLIC  DownloadDataTypeEnum
}

func GetDownloadDataTypeEnumEnum() DownloadDataTypeEnumEnum {
	return DownloadDataTypeEnumEnum{
		PRIVATE: DownloadDataTypeEnum{
			value: "PRIVATE",
		},
		PUBLIC: DownloadDataTypeEnum{
			value: "PUBLIC",
		},
	}
}

func (c DownloadDataTypeEnum) Value() string {
	return c.value
}

func (c DownloadDataTypeEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DownloadDataTypeEnum) UnmarshalJSON(b []byte) error {
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
