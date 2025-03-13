package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ImportDataClassificationRuleDto 导入成功的识别规则。
type ImportDataClassificationRuleDto struct {

	// 数据识别规则id。
	Uuid *string `json:"uuid,omitempty"`

	// 识别规则类型 * BUILTIN 内置 * CUSTOM 自定义
	ClassificationType *ImportDataClassificationRuleDtoClassificationType `json:"classification_type,omitempty"`

	// 数据密级id。
	SecrecyLevel *string `json:"secrecy_level,omitempty"`

	// 数据识别规则名称。
	Name *string `json:"name,omitempty"`

	// 是否启用。
	Enable *bool `json:"enable,omitempty"`

	// 识别规则类型 * NONE 无 * REGULAR 正则表达式 * DEFAULT 默认
	Method *ImportDataClassificationRuleDtoMethod `json:"method,omitempty"`

	// 描述。
	Description *string `json:"description,omitempty"`

	// 数据分类id。
	CategoryId *string `json:"category_id,omitempty"`

	// 预置规则id。
	BuiltinRuleId *string `json:"builtin_rule_id,omitempty"`

	// 更新人。
	UpdatedBy *string `json:"updated_by,omitempty"`

	// 更新时间。
	UpdateAt *int64 `json:"update_at,omitempty"`

	// 创建人。
	CreatedBy *string `json:"created_by,omitempty"`

	// 创建时间。
	CreateAt *int64 `json:"create_at,omitempty"`
}

func (o ImportDataClassificationRuleDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportDataClassificationRuleDto struct{}"
	}

	return strings.Join([]string{"ImportDataClassificationRuleDto", string(data)}, " ")
}

type ImportDataClassificationRuleDtoClassificationType struct {
	value string
}

type ImportDataClassificationRuleDtoClassificationTypeEnum struct {
	BUILTIN ImportDataClassificationRuleDtoClassificationType
	CUSTOM  ImportDataClassificationRuleDtoClassificationType
}

func GetImportDataClassificationRuleDtoClassificationTypeEnum() ImportDataClassificationRuleDtoClassificationTypeEnum {
	return ImportDataClassificationRuleDtoClassificationTypeEnum{
		BUILTIN: ImportDataClassificationRuleDtoClassificationType{
			value: "BUILTIN",
		},
		CUSTOM: ImportDataClassificationRuleDtoClassificationType{
			value: "CUSTOM",
		},
	}
}

func (c ImportDataClassificationRuleDtoClassificationType) Value() string {
	return c.value
}

func (c ImportDataClassificationRuleDtoClassificationType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImportDataClassificationRuleDtoClassificationType) UnmarshalJSON(b []byte) error {
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

type ImportDataClassificationRuleDtoMethod struct {
	value string
}

type ImportDataClassificationRuleDtoMethodEnum struct {
	NONE    ImportDataClassificationRuleDtoMethod
	REGULAR ImportDataClassificationRuleDtoMethod
	DEFAULT ImportDataClassificationRuleDtoMethod
}

func GetImportDataClassificationRuleDtoMethodEnum() ImportDataClassificationRuleDtoMethodEnum {
	return ImportDataClassificationRuleDtoMethodEnum{
		NONE: ImportDataClassificationRuleDtoMethod{
			value: "NONE",
		},
		REGULAR: ImportDataClassificationRuleDtoMethod{
			value: "REGULAR",
		},
		DEFAULT: ImportDataClassificationRuleDtoMethod{
			value: "DEFAULT",
		},
	}
}

func (c ImportDataClassificationRuleDtoMethod) Value() string {
	return c.value
}

func (c ImportDataClassificationRuleDtoMethod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImportDataClassificationRuleDtoMethod) UnmarshalJSON(b []byte) error {
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
