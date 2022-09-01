package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ExecuteQaChatResponse struct {

	// 回复类型： 0   问答型机器人回复。 1   任务型机器人回复。 2   闲聊回复。 3   图谱问答回复。 4   文档问答回复。 5   表格问答回复。
	ReplyType *int32 `json:"reply_type,omitempty" xml:"reply_type"`

	QabotAnswers *QaBotAnswersNew `json:"qabot_answers,omitempty" xml:"qabot_answers"`

	ChatAnswers *ChatAnswers `json:"chat_answers,omitempty" xml:"chat_answers"`

	TaskbotAnswers *TaskBotAnswers `json:"taskbot_answers,omitempty" xml:"taskbot_answers"`

	DocqaAnswers *DocBotAnswers `json:"docqa_answers,omitempty" xml:"docqa_answers"`

	TableqaAnswers *TableQaAnswers `json:"tableqa_answers,omitempty" xml:"tableqa_answers"`

	// 会话ID，在下一次请求中传入改id表示继续会话。
	SessionId *string `json:"session_id,omitempty" xml:"session_id"`

	KbqaAnswers *KbqaAnswers `json:"kbqa_answers,omitempty" xml:"kbqa_answers"`

	// 请求ID。用来标记调用失败时，用来标记本次问答。
	RequestId      *string `json:"request_id,omitempty" xml:"request_id"`
	HttpStatusCode int     `json:"-"`
}

func (o ExecuteQaChatResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteQaChatResponse struct{}"
	}

	return strings.Join([]string{"ExecuteQaChatResponse", string(data)}, " ")
}
