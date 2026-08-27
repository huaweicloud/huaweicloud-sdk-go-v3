package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModelInfo 模型信息，下发和查询共用。
type ModelInfo struct {

	// 模型 ID（业务主键）。
	Id *string `json:"id,omitempty"`

	// 模型名称。
	Name *string `json:"name,omitempty"`

	// 供应商侧模型 ID。
	ProviderModelId *string `json:"provider_model_id,omitempty"`

	// 输入类型数组。
	Input *[]string `json:"input,omitempty"`

	// 是否支持推理。
	Reasoning *bool `json:"reasoning,omitempty"`

	// 模型更新时间。
	UpdateTime *string `json:"update_time,omitempty"`

	// 上下文窗口。
	ContextWindow *int32 `json:"context_window,omitempty"`

	// 最大输出 token 数。
	MaxTokens *int32 `json:"max_tokens,omitempty"`

	// 纳管类型（BACKEND_MANAGE后台管理/CUSTOM自定义）,业务下发的都是BACKEND_MANAGE。
	Type *string `json:"type,omitempty"`
}

func (o ModelInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModelInfo struct{}"
	}

	return strings.Join([]string{"ModelInfo", string(data)}, " ")
}
