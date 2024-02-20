package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ValidatePolicyFinding 可用于改进策略的可操作建议。
type ValidatePolicyFinding struct {

	// 一条本地化消息提供了如何解决该问题的指导。
	FindingDetails string `json:"finding_details"`

	// 影响级别。  安全警告：策略允许的范围过大。  错误：策略不符合策略语法规则。  警告：非安全问题，策略不符合策略编写最佳实践。  建议：改进策略，不影响访问范围。
	FindingType ValidatePolicyFindingFindingType `json:"finding_type"`

	// 问题码提供了与此查找结果关联的问题的标识符。
	IssueCode string `json:"issue_code"`

	// 指向与此查找结果关联的相关文档的链接。
	LearnMoreLink string `json:"learn_more_link"`

	// 策略文档中与查找结果相关的位置列表。
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
