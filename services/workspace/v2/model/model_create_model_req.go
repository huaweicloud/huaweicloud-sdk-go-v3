package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateModelReq 新增模型请求项。
type CreateModelReq struct {

	// 模型名称。
	Name string `json:"name"`

	// 输入类型数组。
	Input *[]string `json:"input,omitempty"`

	// 供应商侧模型标识。
	ProviderModelId string `json:"provider_model_id"`

	// 模型描述。
	Description *string `json:"description,omitempty"`

	// 是否支持推理。
	Reasoning *bool `json:"reasoning,omitempty"`

	Cost *ModelCost `json:"cost,omitempty"`

	// 最大上下文窗口。
	ContextWindow *int32 `json:"context_window,omitempty"`

	// 最大输出Token数。
	MaxTokens *int32 `json:"max_tokens,omitempty"`

	Compat *ModelCompat `json:"compat,omitempty"`
}

func (o CreateModelReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateModelReq struct{}"
	}

	return strings.Join([]string{"CreateModelReq", string(data)}, " ")
}
