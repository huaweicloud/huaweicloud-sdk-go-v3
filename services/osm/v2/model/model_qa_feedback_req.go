package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type QaFeedbackReq struct {

	// 用户问题
	Question string `json:"question"`

	// 反馈记录Id
	FeedbackId *string `json:"feedback_id,omitempty"`

	// - IROBOT_QA:  - RECOMMEND_WORD_QA:
	QaPairSource *QaFeedbackReqQaPairSource `json:"qa_pair_source,omitempty"`

	// 反馈选项id
	FeedbackOptionId *string `json:"feedback_option_id,omitempty"`

	// 反馈描述
	FeedbackDescription *string `json:"feedback_description,omitempty"`

	// 语料id
	QaPairId *string `json:"qa_pair_id,omitempty"`

	// 请求id
	RequestId *string `json:"request_id,omitempty"`

	// - FAQ:  - FLOW:
	FeedbackSource QaFeedbackReqFeedbackSource `json:"feedback_source"`

	// 流程节点Id
	FlowNodeId *string `json:"flow_node_id,omitempty"`
}

func (o QaFeedbackReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QaFeedbackReq struct{}"
	}

	return strings.Join([]string{"QaFeedbackReq", string(data)}, " ")
}

type QaFeedbackReqQaPairSource struct {
	value string
}

type QaFeedbackReqQaPairSourceEnum struct {
	IROBOT_QA         QaFeedbackReqQaPairSource
	RECOMMEND_WORD_QA QaFeedbackReqQaPairSource
}

func GetQaFeedbackReqQaPairSourceEnum() QaFeedbackReqQaPairSourceEnum {
	return QaFeedbackReqQaPairSourceEnum{
		IROBOT_QA: QaFeedbackReqQaPairSource{
			value: "IROBOT_QA",
		},
		RECOMMEND_WORD_QA: QaFeedbackReqQaPairSource{
			value: "RECOMMEND_WORD_QA",
		},
	}
}

func (c QaFeedbackReqQaPairSource) Value() string {
	return c.value
}

func (c QaFeedbackReqQaPairSource) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QaFeedbackReqQaPairSource) UnmarshalJSON(b []byte) error {
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

type QaFeedbackReqFeedbackSource struct {
	value string
}

type QaFeedbackReqFeedbackSourceEnum struct {
	FAQ  QaFeedbackReqFeedbackSource
	FLOW QaFeedbackReqFeedbackSource
}

func GetQaFeedbackReqFeedbackSourceEnum() QaFeedbackReqFeedbackSourceEnum {
	return QaFeedbackReqFeedbackSourceEnum{
		FAQ: QaFeedbackReqFeedbackSource{
			value: "FAQ",
		},
		FLOW: QaFeedbackReqFeedbackSource{
			value: "FLOW",
		},
	}
}

func (c QaFeedbackReqFeedbackSource) Value() string {
	return c.value
}

func (c QaFeedbackReqFeedbackSource) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QaFeedbackReqFeedbackSource) UnmarshalJSON(b []byte) error {
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
