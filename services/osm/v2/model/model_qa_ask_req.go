package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type QaAskReq struct {

	// 用户输入问题
	Question string `json:"question"`

	// 最大返回数据条数
	Top *int32 `json:"top,omitempty"`

	// 主题列表
	Themes *[]RelationTheme `json:"themes,omitempty"`

	// - PORTAL:  - INCIDENT:
	Source *QaAskReqSource `json:"source,omitempty"`

	// 会话ID
	SessionId *string `json:"session_id,omitempty"`

	// 语料ID
	SourceQaPairId *string `json:"source_qa_pair_id,omitempty"`

	// 是否需要备选答案
	AlternativeAnswerEnable *bool `json:"alternative_answer_enable,omitempty"`

	// 产品类型Id
	ProductTypeId *string `json:"product_type_id,omitempty"`

	// 指定的节点ID
	SpecifyNodeId *string `json:"specify_node_id,omitempty"`
}

func (o QaAskReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QaAskReq struct{}"
	}

	return strings.Join([]string{"QaAskReq", string(data)}, " ")
}

type QaAskReqSource struct {
	value string
}

type QaAskReqSourceEnum struct {
	PORTAL   QaAskReqSource
	INCIDENT QaAskReqSource
}

func GetQaAskReqSourceEnum() QaAskReqSourceEnum {
	return QaAskReqSourceEnum{
		PORTAL: QaAskReqSource{
			value: "PORTAL",
		},
		INCIDENT: QaAskReqSource{
			value: "INCIDENT",
		},
	}
}

func (c QaAskReqSource) Value() string {
	return c.value
}

func (c QaAskReqSource) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QaAskReqSource) UnmarshalJSON(b []byte) error {
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
