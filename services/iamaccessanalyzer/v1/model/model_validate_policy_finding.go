package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ValidatePolicyFinding struct {

	// 一条本地化消息提供了如何解决该问题的指导。
	FindingDetails string `json:"finding_details"`

	// 影响级别。  安全：策略存在安全风险，可能是允许访问的权限过于宽松等导致。  错误：存在策略无法运行的错误，如语法错误、参数错误等。存在错误的情况下策略无法创建。  警告：存在策略无法运行的警告，如参数取值类型不匹配等。存在警告的情况下策略可以创建。  建议：不影响策略运行，但策略可能不能达到预期的效果。如存在空数组、空对象条件等。
	FindingType ValidatePolicyFindingFindingType `json:"finding_type"`

	// 问题码提供了与此校验结果关联的问题的标识符。
	IssueCode string `json:"issue_code"`

	// 指向与此校验结果关联的相关文档的链接。
	LearnMoreLink string `json:"learn_more_link"`

	// 策略文档中与校验结果相关的位置列表。
	Locations []Location `json:"locations"`
}

func (o ValidatePolicyFinding) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidatePolicyFinding struct{}"
	}

	return strings.Join([]string{"ValidatePolicyFinding", string(data)}, " ")
}

type ValidatePolicyFindingFindingType struct {
	value string
}

type ValidatePolicyFindingFindingTypeEnum struct {
	SECURITY_WARNING ValidatePolicyFindingFindingType
	ERROR            ValidatePolicyFindingFindingType
	WARNING          ValidatePolicyFindingFindingType
	SUGGESTION       ValidatePolicyFindingFindingType
}

func GetValidatePolicyFindingFindingTypeEnum() ValidatePolicyFindingFindingTypeEnum {
	return ValidatePolicyFindingFindingTypeEnum{
		SECURITY_WARNING: ValidatePolicyFindingFindingType{
			value: "security_warning",
		},
		ERROR: ValidatePolicyFindingFindingType{
			value: "error",
		},
		WARNING: ValidatePolicyFindingFindingType{
			value: "warning",
		},
		SUGGESTION: ValidatePolicyFindingFindingType{
			value: "suggestion",
		},
	}
}

func (c ValidatePolicyFindingFindingType) Value() string {
	return c.value
}

func (c ValidatePolicyFindingFindingType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ValidatePolicyFindingFindingType) UnmarshalJSON(b []byte) error {
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
