package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DataMaskInfo data mask info。
type DataMaskInfo struct {

	// data mask type：REDACT修订， HASH哈希加密，PARTIAL_MASK部分掩码，NULLIFY用NULL值替换原值，UNMASKED原样显示，DATA_ONLY_SHOW_YEAR仅显示日期字符串的年份部分，CUSTOM自定义脱敏规则。
	DataMaskType DataMaskInfoDataMaskType `json:"data_mask_type"`

	// 遮掩规则选项。
	DataMask string `json:"data_mask"`
}

func (o DataMaskInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataMaskInfo struct{}"
	}

	return strings.Join([]string{"DataMaskInfo", string(data)}, " ")
}

type DataMaskInfoDataMaskType struct {
	value string
}

type DataMaskInfoDataMaskTypeEnum struct {
	REDACT              DataMaskInfoDataMaskType
	HASH                DataMaskInfoDataMaskType
	PARTIAL_MASK        DataMaskInfoDataMaskType
	NULLIFY             DataMaskInfoDataMaskType
	UNMASKED            DataMaskInfoDataMaskType
	DATA_ONLY_SHOW_YEAR DataMaskInfoDataMaskType
	CUSTOM              DataMaskInfoDataMaskType
}

func GetDataMaskInfoDataMaskTypeEnum() DataMaskInfoDataMaskTypeEnum {
	return DataMaskInfoDataMaskTypeEnum{
		REDACT: DataMaskInfoDataMaskType{
			value: "REDACT",
		},
		HASH: DataMaskInfoDataMaskType{
			value: "HASH",
		},
		PARTIAL_MASK: DataMaskInfoDataMaskType{
			value: "PARTIAL_MASK",
		},
		NULLIFY: DataMaskInfoDataMaskType{
			value: "NULLIFY",
		},
		UNMASKED: DataMaskInfoDataMaskType{
			value: "UNMASKED",
		},
		DATA_ONLY_SHOW_YEAR: DataMaskInfoDataMaskType{
			value: "DATA_ONLY_SHOW_YEAR",
		},
		CUSTOM: DataMaskInfoDataMaskType{
			value: "CUSTOM",
		},
	}
}

func (c DataMaskInfoDataMaskType) Value() string {
	return c.value
}

func (c DataMaskInfoDataMaskType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DataMaskInfoDataMaskType) UnmarshalJSON(b []byte) error {
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
