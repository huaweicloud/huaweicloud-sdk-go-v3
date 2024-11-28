package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExecuteTextCompletionResponse Response Object
type ExecuteTextCompletionResponse struct {

	// 响应ID
	Id *string `json:"id,omitempty"`

	// 响应时间
	Created *int32 `json:"created,omitempty"`

	// 模型回复
	Choices *[]TextChoice `json:"choices,omitempty"`

	Usage          *CompletionUsage `json:"usage,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ExecuteTextCompletionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteTextCompletionResponse struct{}"
	}

	return strings.Join([]string{"ExecuteTextCompletionResponse", string(data)}, " ")
}
