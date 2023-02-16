package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Answer struct {

	// 答案内容
	Content *string `json:"content,omitempty"`

	// 答案Id
	Id *string `json:"id,omitempty"`

	// 标签列表
	Tags *[]string `json:"tags,omitempty"`

	// - TEXT:  - RICH_TEXT:  - FLOW:  - QA_PAIR:
	AnswerType *AnswerAnswerType `json:"answer_type,omitempty"`

	QaFlowEntry *CbsFlowEntry `json:"qa_flow_entry,omitempty"`
}

func (o Answer) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Answer struct{}"
	}

	return strings.Join([]string{"Answer", string(data)}, " ")
}

type AnswerAnswerType struct {
	value string
}

type AnswerAnswerTypeEnum struct {
	TEXT      AnswerAnswerType
	RICH_TEXT AnswerAnswerType
	FLOW      AnswerAnswerType
	QA_PAIR   AnswerAnswerType
}

func GetAnswerAnswerTypeEnum() AnswerAnswerTypeEnum {
	return AnswerAnswerTypeEnum{
		TEXT: AnswerAnswerType{
			value: "TEXT",
		},
		RICH_TEXT: AnswerAnswerType{
			value: "RICH_TEXT",
		},
		FLOW: AnswerAnswerType{
			value: "FLOW",
		},
		QA_PAIR: AnswerAnswerType{
			value: "QA_PAIR",
		},
	}
}

func (c AnswerAnswerType) Value() string {
	return c.value
}

func (c AnswerAnswerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AnswerAnswerType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
