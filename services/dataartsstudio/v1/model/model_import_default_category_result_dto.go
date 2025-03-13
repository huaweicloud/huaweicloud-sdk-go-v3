package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ImportDefaultCategoryResultDto struct {

	// 导入状态 * success 导入成功 * failed 导入失败
	ImportStatus *ImportDefaultCategoryResultDtoImportStatus `json:"import_status,omitempty"`

	// 导入错误原因。
	ImportErrorMessage *string `json:"import_error_message,omitempty"`

	// 子分类导入结果。
	Children *[]ImportDefaultCategoryResultDto `json:"children,omitempty"`

	// 此分类绑定的规则导入的结果。
	RuleResult *[]ImportDefaultRuleResultDto `json:"rule_result,omitempty"`

	// 数据分类id。
	Uuid *string `json:"uuid,omitempty"`

	// 数据分类名称。
	Name *string `json:"name,omitempty"`

	// 数据分类描述。
	Description *string `json:"description,omitempty"`
}

func (o ImportDefaultCategoryResultDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportDefaultCategoryResultDto struct{}"
	}

	return strings.Join([]string{"ImportDefaultCategoryResultDto", string(data)}, " ")
}

type ImportDefaultCategoryResultDtoImportStatus struct {
	value string
}

type ImportDefaultCategoryResultDtoImportStatusEnum struct {
	SUCCESS ImportDefaultCategoryResultDtoImportStatus
	FAILED  ImportDefaultCategoryResultDtoImportStatus
}

func GetImportDefaultCategoryResultDtoImportStatusEnum() ImportDefaultCategoryResultDtoImportStatusEnum {
	return ImportDefaultCategoryResultDtoImportStatusEnum{
		SUCCESS: ImportDefaultCategoryResultDtoImportStatus{
			value: "success",
		},
		FAILED: ImportDefaultCategoryResultDtoImportStatus{
			value: "failed",
		},
	}
}

func (c ImportDefaultCategoryResultDtoImportStatus) Value() string {
	return c.value
}

func (c ImportDefaultCategoryResultDtoImportStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ImportDefaultCategoryResultDtoImportStatus) UnmarshalJSON(b []byte) error {
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
