package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateDocumentRequestBody struct {

	// 作业名称
	Name string `json:"name"`

	// 作业内容，DSL语句，最大长度64KB
	Content string `json:"content"`

	// 企业项目id
	EnterpriseProjectId string `json:"enterprise_project_id"`

	// 风险等级。 - LOW：低。 - MEDIUM：中。 - HIGH：高。
	RiskLevel CreateDocumentRequestBodyRiskLevel `json:"risk_level"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 标签列表
	Tags *[]CreateDocumentRequestBodyTags `json:"tags,omitempty"`
}

func (o CreateDocumentRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDocumentRequestBody struct{}"
	}

	return strings.Join([]string{"CreateDocumentRequestBody", string(data)}, " ")
}

type CreateDocumentRequestBodyRiskLevel struct {
	value string
}

type CreateDocumentRequestBodyRiskLevelEnum struct {
	LOW    CreateDocumentRequestBodyRiskLevel
	MEDIUM CreateDocumentRequestBodyRiskLevel
	HIGH   CreateDocumentRequestBodyRiskLevel
}

func GetCreateDocumentRequestBodyRiskLevelEnum() CreateDocumentRequestBodyRiskLevelEnum {
	return CreateDocumentRequestBodyRiskLevelEnum{
		LOW: CreateDocumentRequestBodyRiskLevel{
			value: "LOW",
		},
		MEDIUM: CreateDocumentRequestBodyRiskLevel{
			value: "MEDIUM",
		},
		HIGH: CreateDocumentRequestBodyRiskLevel{
			value: "HIGH",
		},
	}
}

func (c CreateDocumentRequestBodyRiskLevel) Value() string {
	return c.value
}

func (c CreateDocumentRequestBodyRiskLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateDocumentRequestBodyRiskLevel) UnmarshalJSON(b []byte) error {
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
