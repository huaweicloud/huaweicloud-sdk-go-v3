package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateModelResponse Response Object
type UpdateModelResponse struct {

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

	// 供应商侧模型标识。
	ProviderModelId *string `json:"provider_model_id,omitempty"`

	// 供应商id。
	ProviderId *string `json:"provider_id,omitempty"`

	// 供应商名称。
	ProviderName *string `json:"provider_name,omitempty"`

	// 模型描述。
	Description *string `json:"description,omitempty"`

	// 组内排序优先级。
	Priority *int32 `json:"priority,omitempty"`

	// 创建时间。
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间。
	UpdateTime *string `json:"update_time,omitempty"`

	// 关联为默认模型的模型分组
	Groups         *[]AttachModelGroupInfo `json:"groups,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o UpdateModelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateModelResponse struct{}"
	}

	return strings.Join([]string{"UpdateModelResponse", string(data)}, " ")
}
