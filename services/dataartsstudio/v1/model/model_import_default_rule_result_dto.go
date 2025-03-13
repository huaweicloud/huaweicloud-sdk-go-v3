package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ImportDefaultRuleResultDto struct {

	// 导入状态 * success 导入成功 * failed 导入失败
	ImportStatus *ImportDefaultRuleResultDtoImportStatus `json:"import_status,omitempty"`

	// 导入错误原因。
	ImportErrorMessage *string `json:"import_error_message,omitempty"`

	// 内置规则模板id。
	Uuid *string `json:"uuid,omitempty"`

	ImportDataClassificationRule *ImportDataClassificationRuleDto `json:"import_data_classification_rule,omitempty"`

	// 数据识别规则名称。
	RuleName *string `json:"rule_name,omitempty"`

	// 数据识别规则类型 * REGEX 正则表达式 * KEYWORD 关键字
	RuleType *ImportDefaultRuleResultDtoRuleType `json:"rule_type,omitempty"`

	// 规则描述。
	RuleDesc *string `json:"rule_desc,omitempty"`

	// 英文名称。
	RuleNameEn *string `json:"rule_name_en,omitempty"`

	// 英文描述。
	RuleDescEn *string `json:"rule_desc_en,omitempty"`

	// 规则所属地区。
	Country *string `json:"country,omitempty"`
}

func (o ImportDefaultRuleResultDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportDefaultRuleResultDto struct{}"
	}

	return strings.Join([]string{"ImportDefaultRuleResultDto", string(data)}, " ")
}

type ImportDefaultRuleResultDtoImportStatus struct {
	value string
}

type ImportDefaultRuleResultDtoImportStatusEnum struct {
	SUCCESS ImportDefaultRuleResultDtoImportStatus
	FAILED  ImportDefaultRuleResultDtoImportStatus
}

func GetImportDefaultRuleResultDtoImportStatusEnum() ImportDefaultRuleResultDtoImportStatusEnum {
	return ImportDefaultRuleResultDtoImportStatusEnum{
		SUCCESS: ImportDefaultRuleResultDtoImportStatus{
			value: "success",
		},
		FAILED: ImportDefaultRuleResultDtoImportStatus{
			value: "failed",
		},
	}
}

func (c ImportDefaultRuleResultDtoImportStatus) Value() string {
	return c.value
}

func (c ImportDefaultRuleResultDtoImportStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImportDefaultRuleResultDtoImportStatus) UnmarshalJSON(b []byte) error {
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

type ImportDefaultRuleResultDtoRuleType struct {
	value string
}

type ImportDefaultRuleResultDtoRuleTypeEnum struct {
	REGEX   ImportDefaultRuleResultDtoRuleType
	KEYWORD ImportDefaultRuleResultDtoRuleType
}

func GetImportDefaultRuleResultDtoRuleTypeEnum() ImportDefaultRuleResultDtoRuleTypeEnum {
	return ImportDefaultRuleResultDtoRuleTypeEnum{
		REGEX: ImportDefaultRuleResultDtoRuleType{
			value: "REGEX",
		},
		KEYWORD: ImportDefaultRuleResultDtoRuleType{
			value: "KEYWORD",
		},
	}
}

func (c ImportDefaultRuleResultDtoRuleType) Value() string {
	return c.value
}

func (c ImportDefaultRuleResultDtoRuleType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImportDefaultRuleResultDtoRuleType) UnmarshalJSON(b []byte) error {
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
