package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AnalyzerType 分析器的类型。
type AnalyzerType struct {
	value string
}

type AnalyzerTypeEnum struct {
	ACCOUNT                    AnalyzerType
	ORGANIZATION               AnalyzerType
	ACCOUNT_UNUSED_ACCESS      AnalyzerType
	ORGANIZATION_UNUSED_ACCESS AnalyzerType
}

func GetAnalyzerTypeEnum() AnalyzerTypeEnum {
	return AnalyzerTypeEnum{
		ACCOUNT: AnalyzerType{
			value: "account",
		},
		ORGANIZATION: AnalyzerType{
			value: "organization",
		},
		ACCOUNT_UNUSED_ACCESS: AnalyzerType{
			value: "account_unused_access",
		},
		ORGANIZATION_UNUSED_ACCESS: AnalyzerType{
			value: "organization_unused_access",
		},
	}
}

func (c AnalyzerType) Value() string {
	return c.value
}

func (c AnalyzerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AnalyzerType) UnmarshalJSON(b []byte) error {
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
