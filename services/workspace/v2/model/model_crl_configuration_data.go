package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CrlConfigurationData 证书crl配置。
type CrlConfigurationData struct {

	// 是否开启crl配置。
	Enable bool `json:"enable"`

	// 系统生成SYSTEM, 用户自定义CUSTOMIZE。
	Type CrlConfigurationDataType `json:"type"`

	// 当用户自定义时手动输入。
	CrlUrl *string `json:"crl_url,omitempty"`

	// 更新周期。
	ValidDay int32 `json:"valid_day"`
}

func (o CrlConfigurationData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CrlConfigurationData struct{}"
	}

	return strings.Join([]string{"CrlConfigurationData", string(data)}, " ")
}

type CrlConfigurationDataType struct {
	value string
}

type CrlConfigurationDataTypeEnum struct {
	SYSTEM    CrlConfigurationDataType
	CUSTOMIZE CrlConfigurationDataType
}

func GetCrlConfigurationDataTypeEnum() CrlConfigurationDataTypeEnum {
	return CrlConfigurationDataTypeEnum{
		SYSTEM: CrlConfigurationDataType{
			value: "SYSTEM",
		},
		CUSTOMIZE: CrlConfigurationDataType{
			value: "CUSTOMIZE",
		},
	}
}

func (c CrlConfigurationDataType) Value() string {
	return c.value
}

func (c CrlConfigurationDataType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CrlConfigurationDataType) UnmarshalJSON(b []byte) error {
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
