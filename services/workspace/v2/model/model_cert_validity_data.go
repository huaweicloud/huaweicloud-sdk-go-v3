package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CertValidityData 证书有效期配置。
type CertValidityData struct {

	// 时间单位, YEAR 年。
	Type CertValidityDataType `json:"type"`

	// 时间数值。
	Value int32 `json:"value"`

	// 证书过期时间。
	StartFrom *string `json:"start_from,omitempty"`
}

func (o CertValidityData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CertValidityData struct{}"
	}

	return strings.Join([]string{"CertValidityData", string(data)}, " ")
}

type CertValidityDataType struct {
	value string
}

type CertValidityDataTypeEnum struct {
	YEAR CertValidityDataType
}

func GetCertValidityDataTypeEnum() CertValidityDataTypeEnum {
	return CertValidityDataTypeEnum{
		YEAR: CertValidityDataType{
			value: "YEAR",
		},
	}
}

func (c CertValidityDataType) Value() string {
	return c.value
}

func (c CertValidityDataType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CertValidityDataType) UnmarshalJSON(b []byte) error {
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
