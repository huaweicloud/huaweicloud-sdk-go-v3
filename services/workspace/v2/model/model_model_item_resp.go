package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModelItemResp 模型信息项响应（简化版）。
type ModelItemResp struct {

	// 模型id。
	Id *string `json:"id,omitempty"`

	// 模型名称。
	Name *string `json:"name,omitempty"`

	// 供应商侧模型标识。
	ProviderModelId *string `json:"provider_model_id,omitempty"`

	// 输入类型数组。
	Input *[]string `json:"input,omitempty"`

	// 模型描述。
	Description *string `json:"description,omitempty"`

	// 是否为内置模型。
	IsBuiltin *bool `json:"is_builtin,omitempty"`
}

func (o ModelItemResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModelItemResp struct{}"
	}

	return strings.Join([]string{"ModelItemResp", string(data)}, " ")
}
