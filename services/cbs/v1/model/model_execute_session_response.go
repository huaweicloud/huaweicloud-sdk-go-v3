package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ExecuteSessionResponse struct {

	// 回复类型： 0   问答型机器人回复。 1   任务型机器人回复。 2   闲聊回复。
	ReplyType *int32 `json:"reply_type,omitempty" xml:"reply_type"`

	QabotAnswers *QaBotAnswers `json:"qabot_answers,omitempty" xml:"qabot_answers"`

	ChatAnswers *ChatAnswers `json:"chat_answers,omitempty" xml:"chat_answers"`

	TaskbotAnswers *TaskBotAnswers `json:"taskbot_answers,omitempty" xml:"taskbot_answers"`

	// 请求ID。
	RequestId      *string `json:"request_id,omitempty" xml:"request_id"`
	HttpStatusCode int     `json:"-"`
}

func (o ExecuteSessionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteSessionResponse struct{}"
	}

	return strings.Join([]string{"ExecuteSessionResponse", string(data)}, " ")
}
