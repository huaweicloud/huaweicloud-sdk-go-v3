package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WeakConstraintDto 分子约束
type WeakConstraintDto struct {

	// 自定义模型id，仅弱约束为模型时填写
	Id *string `json:"id,omitempty"`

	// 属性名称
	Name *string `json:"name,omitempty"`

	Type *WeakConstraintType `json:"type"`

	// 属性约束类型bool的参数
	Bool *bool `json:"bool,omitempty"`

	// 属性约束类型range的参数
	Range *[]float32 `json:"range,omitempty"`

	Struct *StructureConstraintParamsDto `json:"struct,omitempty"`

	// 属性约束类型minimize和maximize的参数
	Quantiles *[]float32 `json:"quantiles,omitempty"`

	Interaction *InteractionConstraintDto `json:"interaction,omitempty"`
}

func (o WeakConstraintDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WeakConstraintDto struct{}"
	}

	return strings.Join([]string{"WeakConstraintDto", string(data)}, " ")
}
