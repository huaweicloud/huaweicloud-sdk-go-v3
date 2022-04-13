package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ExecuteSessionResponse struct {
	// 回复类型： 0   问答型机器人回复。 1   任务型机器人回复。 2   闲聊回复。

	ReplyType *int32 `json:"reply_type,omitempty"`

	QabotAnswers *QaBotAnswers `json:"qabot_answers,omitempty"`

	ChatAnswers *ChatAnswers `json:"chat_answers,omitempty"`

	TaskbotAnswers *TaskBotAnswers `json:"taskbot_answers,omitempty"`
	// 请求ID。

	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExecuteSessionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteSessionResponse struct{}"
	}

	return strings.Join([]string{"ExecuteSessionResponse", string(data)}, " ")
}
