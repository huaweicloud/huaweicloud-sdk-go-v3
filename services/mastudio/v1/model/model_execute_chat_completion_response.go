package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteChatCompletionResponse Response Object
type ExecuteChatCompletionResponse struct {

	// 响应ID
	Id *string `json:"id,omitempty"`

	// 响应时间
	Created *int32 `json:"created,omitempty"`

	// 模型回复
	Choices *[]ChatChoice `json:"choices,omitempty"`

	Usage          *CompletionUsage `json:"usage,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ExecuteChatCompletionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteChatCompletionResponse struct{}"
	}

	return strings.Join([]string{"ExecuteChatCompletionResponse", string(data)}, " ")
}
