package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type QaFlowHitNodeVo struct {

	// ID
	Id *string `json:"id,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 内容
	Content *string `json:"content,omitempty"`

	// - TEXT:  - RICH_TEXT:  - FLOW:  - QA_PAIR:
	AnswerType *QaFlowHitNodeVoAnswerType `json:"answer_type,omitempty"`
}

func (o QaFlowHitNodeVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QaFlowHitNodeVo struct{}"
	}

	return strings.Join([]string{"QaFlowHitNodeVo", string(data)}, " ")
}

type QaFlowHitNodeVoAnswerType struct {
	value string
}

type QaFlowHitNodeVoAnswerTypeEnum struct {
	TEXT      QaFlowHitNodeVoAnswerType
	RICH_TEXT QaFlowHitNodeVoAnswerType
	FLOW      QaFlowHitNodeVoAnswerType
	QA_PAIR   QaFlowHitNodeVoAnswerType
}

func GetQaFlowHitNodeVoAnswerTypeEnum() QaFlowHitNodeVoAnswerTypeEnum {
	return QaFlowHitNodeVoAnswerTypeEnum{
		TEXT: QaFlowHitNodeVoAnswerType{
			value: "TEXT",
		},
		RICH_TEXT: QaFlowHitNodeVoAnswerType{
			value: "RICH_TEXT",
		},
		FLOW: QaFlowHitNodeVoAnswerType{
			value: "FLOW",
		},
		QA_PAIR: QaFlowHitNodeVoAnswerType{
			value: "QA_PAIR",
		},
	}
}

func (c QaFlowHitNodeVoAnswerType) Value() string {
	return c.value
}

func (c QaFlowHitNodeVoAnswerType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QaFlowHitNodeVoAnswerType) UnmarshalJSON(b []byte) error {
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
