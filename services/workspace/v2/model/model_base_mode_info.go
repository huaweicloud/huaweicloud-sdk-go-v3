package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BaseModeInfo struct {

	// 模型id。
	Id *string `json:"id,omitempty"`

	// 模型名称。
	Name *string `json:"name,omitempty"`

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

	// 是否内置模型。
	IsBuiltin *bool `json:"is_builtin,omitempty"`
}

func (o BaseModeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BaseModeInfo struct{}"
	}

	return strings.Join([]string{"BaseModeInfo", string(data)}, " ")
}
