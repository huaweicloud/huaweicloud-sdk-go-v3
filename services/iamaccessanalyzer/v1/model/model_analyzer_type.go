package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AnalyzerType 分析器的类型。 - account: 账号级外部访问分析器 - organization：组织级外部访问分析器 - account_unused_access：账号级未使用访问分析器 - organization_unused_access：组织级未使用访问分析器 - account_privilege_escalation：账号级提权访问分析器 - account_iam_best_practice：账号级IAM最佳实践分析器
type AnalyzerType struct {
	value string
}

type AnalyzerTypeEnum struct {
	ACCOUNT                      AnalyzerType
	ORGANIZATION                 AnalyzerType
	ACCOUNT_UNUSED_ACCESS        AnalyzerType
	ORGANIZATION_UNUSED_ACCESS   AnalyzerType
	ACCOUNT_PRIVILEGE_ESCALATION AnalyzerType
	ACCOUNT_IAM_BEST_PRACTICE    AnalyzerType
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
		ACCOUNT_PRIVILEGE_ESCALATION: AnalyzerType{
			value: "account_privilege_escalation",
		},
		ACCOUNT_IAM_BEST_PRACTICE: AnalyzerType{
			value: "account_iam_best_practice",
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
