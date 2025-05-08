package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type UpdateRunbookRequestBody struct {

	// 作业名称
	Name string `json:"name"`

	// 作业内容，DSL语句，最大长度64KB
	Content string `json:"content"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 风险等级
	RiskLevel *UpdateRunbookRequestBodyRiskLevel `json:"risk_level,omitempty"`
}

func (o UpdateRunbookRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRunbookRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateRunbookRequestBody", string(data)}, " ")
}

type UpdateRunbookRequestBodyRiskLevel struct {
	value string
}

type UpdateRunbookRequestBodyRiskLevelEnum struct {
	LOW    UpdateRunbookRequestBodyRiskLevel
	MEDIUM UpdateRunbookRequestBodyRiskLevel
	HIGH   UpdateRunbookRequestBodyRiskLevel
}

func GetUpdateRunbookRequestBodyRiskLevelEnum() UpdateRunbookRequestBodyRiskLevelEnum {
	return UpdateRunbookRequestBodyRiskLevelEnum{
		LOW: UpdateRunbookRequestBodyRiskLevel{
			value: "LOW",
		},
		MEDIUM: UpdateRunbookRequestBodyRiskLevel{
			value: "MEDIUM",
		},
		HIGH: UpdateRunbookRequestBodyRiskLevel{
			value: "HIGH",
		},
	}
}

func (c UpdateRunbookRequestBodyRiskLevel) Value() string {
	return c.value
}

func (c UpdateRunbookRequestBodyRiskLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateRunbookRequestBodyRiskLevel) UnmarshalJSON(b []byte) error {
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
