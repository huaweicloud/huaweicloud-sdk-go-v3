package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// StorageType 存储类型。 - SAS：高IO - SSD：超高IO - SAS_SD：高IO(软件定义型专用) - SSD_SD：超高IO(软件定义型专用) - SAS_ARM：高IO（鲲鹏） - SSD_ARM：超高IO（鲲鹏） [- VS_SMALL_CAP：视图存储-小容量型](tag:cmcc) [- VS_MEDIUM_CAP：视图存储-中容量型](tag:cmcc) [- VS_LARGE_CAP：视图存储-大容量型](tag:cmcc) [- VS_SMALL_CAP：边缘对象存储-小容量型](tag:hws) [- VS_MEDIUM_CAP：边缘对象存储-中容量型](tag:hws) [- VS_LARGE_CAP：边缘对象存储-大容量型](tag:hws) [- CBR_STANDARD_SMALL_CAP:云备份(基础版)-小容量型](tag:hws) [- CBR_STANDARD_MEDIUM_CAP:云备份(基础版)-中容量型](tag:hws) [- CBR_STANDARD_LARGE_CAP:云备份(基础版)-大容量型](tag:hws) [- CBR_PROFESSIONAL_SMALL_CAP:云备份(专业版)-小容量型](tag:hws) [- CBR_PROFESSIONAL_MEDIUM_CAP:云备份(专业版)-中容量型](tag:hws) [- CBR_PROFESSIONAL_LARGE_CAP:云备份(专业版)-大容量型](tag:hws)
type StorageType struct {
	value string
}

type StorageTypeEnum struct {
	SAS                         StorageType
	SSD                         StorageType
	SAS_SD                      StorageType
	SSD_SD                      StorageType
	SAS_ARM                     StorageType
	SSD_ARM                     StorageType
	VS_SMALL_CAP                StorageType
	VS_MEDIUM_CAP               StorageType
	VS_LARGE_CAP                StorageType
	CBR_STANDARD_SMALL_CAP      StorageType
	CBR_STANDARD_MEDIUM_CAP     StorageType
	CBR_STANDARD_LARGE_CAP      StorageType
	CBR_PROFESSIONAL_SMALL_CAP  StorageType
	CBR_PROFESSIONAL_MEDIUM_CAP StorageType
	CBR_PROFESSIONAL_LARGE_CAP  StorageType
}

func GetStorageTypeEnum() StorageTypeEnum {
	return StorageTypeEnum{
		SAS: StorageType{
			value: "SAS",
		},
		SSD: StorageType{
			value: "SSD",
		},
		SAS_SD: StorageType{
			value: "SAS_SD",
		},
		SSD_SD: StorageType{
			value: "SSD_SD",
		},
		SAS_ARM: StorageType{
			value: "SAS_ARM",
		},
		SSD_ARM: StorageType{
			value: "SSD_ARM",
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
		CBR_STANDARD_SMALL_CAP: StorageType{
			value: "CBR_STANDARD_SMALL_CAP",
		},
		CBR_STANDARD_MEDIUM_CAP: StorageType{
			value: "CBR_STANDARD_MEDIUM_CAP",
		},
		CBR_STANDARD_LARGE_CAP: StorageType{
			value: "CBR_STANDARD_LARGE_CAP",
		},
		CBR_PROFESSIONAL_SMALL_CAP: StorageType{
			value: "CBR_PROFESSIONAL_SMALL_CAP",
		},
		CBR_PROFESSIONAL_MEDIUM_CAP: StorageType{
			value: "CBR_PROFESSIONAL_MEDIUM_CAP",
		},
		CBR_PROFESSIONAL_LARGE_CAP: StorageType{
			value: "CBR_PROFESSIONAL_LARGE_CAP",
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
