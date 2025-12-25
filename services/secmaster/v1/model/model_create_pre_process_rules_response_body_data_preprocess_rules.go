package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreatePreProcessRulesResponseBodyDataPreprocessRules struct {

	// 映射id
	Id *string `json:"id,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 映射id
	MapperId *string `json:"mapper_id,omitempty"`

	// 映射id
	MapperTypeId *string `json:"mapper_type_id,omitempty"`

	// 预处理选项: drop-丢弃
	Action *CreatePreProcessRulesResponseBodyDataPreprocessRulesAction `json:"action,omitempty"`

	// 表达式
	Expression *string `json:"expression,omitempty"`
}

func (o CreatePreProcessRulesResponseBodyDataPreprocessRules) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePreProcessRulesResponseBodyDataPreprocessRules struct{}"
	}

	return strings.Join([]string{"CreatePreProcessRulesResponseBodyDataPreprocessRules", string(data)}, " ")
}

type CreatePreProcessRulesResponseBodyDataPreprocessRulesAction struct {
	value string
}

type CreatePreProcessRulesResponseBodyDataPreprocessRulesActionEnum struct {
	DROP CreatePreProcessRulesResponseBodyDataPreprocessRulesAction
}

func GetCreatePreProcessRulesResponseBodyDataPreprocessRulesActionEnum() CreatePreProcessRulesResponseBodyDataPreprocessRulesActionEnum {
	return CreatePreProcessRulesResponseBodyDataPreprocessRulesActionEnum{
		DROP: CreatePreProcessRulesResponseBodyDataPreprocessRulesAction{
			value: "drop",
		},
	}
}

func (c CreatePreProcessRulesResponseBodyDataPreprocessRulesAction) Value() string {
	return c.value
}

func (c CreatePreProcessRulesResponseBodyDataPreprocessRulesAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreatePreProcessRulesResponseBodyDataPreprocessRulesAction) UnmarshalJSON(b []byte) error {
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
