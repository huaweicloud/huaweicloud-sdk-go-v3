package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChatRequestMessage chat request messqge
type ChatRequestMessage struct {

	// case
	Case *string `json:"case,omitempty"`

	// chat id
	ChatId string `json:"chat_id"`

	// message
	Message string `json:"message"`

	// infer end
	InferEnd *bool `json:"infer_end,omitempty"`

	// prompt
	MetaPrompt *[]string `json:"meta_prompt,omitempty"`

	// need or not
	NeedPreprocess *bool `json:"need_preprocess,omitempty"`

	// user id
	UserId *string `json:"user_id,omitempty"`

	// task parameters
	TaskParameters *interface{} `json:"task_parameters,omitempty"`
}

func (o ChatRequestMessage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChatRequestMessage struct{}"
	}

	return strings.Join([]string{"ChatRequestMessage", string(data)}, " ")
}
