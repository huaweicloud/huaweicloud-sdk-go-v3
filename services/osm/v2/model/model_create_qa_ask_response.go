package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateQaAskResponse Response Object
type CreateQaAskResponse struct {

	// 错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 错误描述
	ErrorMsg *string `json:"error_msg,omitempty"`

	LangResult *LangResult `json:"lang_result,omitempty"`

	// - QA_BOT:  - TASK_BOT:  - CHAT_BOT:  - GRAPH_BOT:  - HW_CLOUD:
	ReplyType *CreateQaAskResponseReplyType `json:"reply_type,omitempty"`

	// 会话ID
	SessionId *string `json:"session_id,omitempty"`

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	AnswersDetail  *AnswerDetail `json:"answers_detail,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o CreateQaAskResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateQaAskResponse struct{}"
	}

	return strings.Join([]string{"CreateQaAskResponse", string(data)}, " ")
}

type CreateQaAskResponseReplyType struct {
	value string
}

type CreateQaAskResponseReplyTypeEnum struct {
	QA_BOT    CreateQaAskResponseReplyType
	TASK_BOT  CreateQaAskResponseReplyType
	CHAT_BOT  CreateQaAskResponseReplyType
	GRAPH_BOT CreateQaAskResponseReplyType
	HW_CLOUD  CreateQaAskResponseReplyType
}

func GetCreateQaAskResponseReplyTypeEnum() CreateQaAskResponseReplyTypeEnum {
	return CreateQaAskResponseReplyTypeEnum{
		QA_BOT: CreateQaAskResponseReplyType{
			value: "QA_BOT",
		},
		TASK_BOT: CreateQaAskResponseReplyType{
			value: "TASK_BOT",
		},
		CHAT_BOT: CreateQaAskResponseReplyType{
			value: "CHAT_BOT",
		},
		GRAPH_BOT: CreateQaAskResponseReplyType{
			value: "GRAPH_BOT",
		},
		HW_CLOUD: CreateQaAskResponseReplyType{
			value: "HW_CLOUD",
		},
	}
}

func (c CreateQaAskResponseReplyType) Value() string {
	return c.value
}

func (c CreateQaAskResponseReplyType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateQaAskResponseReplyType) UnmarshalJSON(b []byte) error {
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
