package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 存储类型。 - SAS：高IO - SSD：超高IO
type StorageType struct {
	value string
}

type StorageTypeEnum struct {
	SAS           StorageType
	SSD           StorageType
	VS_SMALL_CAP  StorageType
	VS_MEDIUM_CAP StorageType
	VS_LARGE_CAP  StorageType
}

func GetStorageTypeEnum() StorageTypeEnum {
	return StorageTypeEnum{
		SAS: StorageType{
			value: "SAS",
		},
		SSD: StorageType{
			value: "SSD",
		},
		VS_SMALL_CAP: StorageType{
			value: "VS_SMALL_CAP",
		},
		VS_MEDIUM_CAP: StorageType{
			value: "VS_MEDIUM_CAP",
		},
		VS_LARGE_CAP: StorageType{
			value: "VS_LARGE_CAP",
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
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
