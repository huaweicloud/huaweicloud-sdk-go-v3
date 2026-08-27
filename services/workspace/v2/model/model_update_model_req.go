package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateModelReq 更新模型请求。
type UpdateModelReq struct {

	// 模型名称。
	Name *string `json:"name,omitempty"`

	// 模型描述。
	Description *string `json:"description,omitempty"`

	// 输入类型数组。
	Input *[]string `json:"input,omitempty"`

	// 最大上下文窗口。
	ContextWindow *int32 `json:"context_window,omitempty"`

	// 最大输出Token数。
	MaxTokens *int32 `json:"max_tokens,omitempty"`

	// 是否支持推理。
	Reasoning *bool `json:"reasoning,omitempty"`

	Cost *ModelCost `json:"cost,omitempty"`

	Compat *ModelCompat `json:"compat,omitempty"`
}

func (o UpdateModelReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateModelReq struct{}"
	}

	return strings.Join([]string{"UpdateModelReq", string(data)}, " ")
}
