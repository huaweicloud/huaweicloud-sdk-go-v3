package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type StorageType struct {
	value string
}

type StorageTypeEnum struct {
	STANDARD StorageType
	COLD     StorageType
}

func GetStorageTypeEnum() StorageTypeEnum {
	return StorageTypeEnum{
		STANDARD: StorageType{
			value: "STANDARD",
		},
		COLD: StorageType{
			value: "COLD",
		},
	}
}

func (c StorageType) Value() string {
	return c.value
}

func (c StorageType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StorageType) UnmarshalJSON(b []byte) error {
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
