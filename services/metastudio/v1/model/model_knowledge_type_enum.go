package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// KnowledgeTypeEnum 知识库类型 * QUESTION_ANSWER：问答 * DOCUMENT：文档
type KnowledgeTypeEnum struct {
	value string
}

type KnowledgeTypeEnumEnum struct {
	QUESTION_ANSWER KnowledgeTypeEnum
	DOCUMENT        KnowledgeTypeEnum
}

func GetKnowledgeTypeEnumEnum() KnowledgeTypeEnumEnum {
	return KnowledgeTypeEnumEnum{
		QUESTION_ANSWER: KnowledgeTypeEnum{
			value: "QUESTION_ANSWER",
		},
		DOCUMENT: KnowledgeTypeEnum{
			value: "DOCUMENT",
		},
	}
}

func (c KnowledgeTypeEnum) Value() string {
	return c.value
}

func (c KnowledgeTypeEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *KnowledgeTypeEnum) UnmarshalJSON(b []byte) error {
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
