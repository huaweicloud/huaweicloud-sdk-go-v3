package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CompletionUsage tokens统计对象
type CompletionUsage struct {

	// 表示模型生成的答案中包含的tokens的数量。
	CompletionTokens float32 `json:"completion_tokens"`

	// 表示生成结果时使用的提示文本的tokens的数量。
	PromptTokens float32 `json:"prompt_tokens"`

	// 对话过程中使用的tokens总数。
	TotalTokens float32 `json:"total_tokens"`
}

func (o CompletionUsage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CompletionUsage struct{}"
	}

	return strings.Join([]string{"CompletionUsage", string(data)}, " ")
}
